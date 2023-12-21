package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/cenkalti/backoff/v4"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type Driver struct {
	ID       string
	addr     common.Address
	state    *State
	eth      *ethclient.Client
	engine   *EngineClient
	selfAddr common.Address
}

func NewDriver(
	ethUrl, engineUrl, jwtFile, statePath string, selfAddr common.Address) *Driver {

	state := NewState(statePath)
	//defer state.Close()

	eth, err := ethclient.DialContext(context.TODO(), ethUrl)
	if err != nil {
		panic(err)
	}

	eng, err := NewEngineClient(engineUrl, jwtFile)
	if err != nil {
		panic(err)
	}

	return &Driver{
		state:    state,
		eth:      eth,
		engine:   eng,
		selfAddr: selfAddr,
	}
}

func (d *Driver) Close() {
	d.state.Close()
}

func (d *Driver) PrepareProposal(context sdk.Context, proposal *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
	latestBlock, err := d.eth.BlockNumber(context)
	if err != nil {
		return nil, err
	}

	latestHash, err := d.eth.BlockByNumber(context, big.NewInt(int64(latestBlock)))
	if err != nil {
		return nil, err
	}

	state := ForkChoiceState{
		HeadHash:           latestHash.Hash(),
		SafeBlockHash:      latestHash.Hash(),
		FinalizedBlockHash: common.Hash{},
	}

	// The engine complains when the withdrawals are empty
	withdrawals := []*Withdrawal{
		{
			Index:     "0x0",
			Validator: "0x0",
			Address:   common.Address{}.Hex(),
			Amount:    "0x0",
		},
	}

	attrs := PayloadAttributes{
		Timestamp:             hexutil.Uint64(time.Now().UnixMilli()),
		PrevRandao:            common.Hash{}, // do we need to generate a randao for the EVM?
		SuggestedFeeRecipient: d.selfAddr,
		Withdrawals:           withdrawals,
	}

	choice, err := d.engine.ForkchoiceUpdatedV2(&state, &attrs)
	if err != nil {
		return nil, err
	}

	payloadId := choice.PayloadId
	status := choice.PayloadStatus

	if status.Status != "VALID" {
		fmt.Printf("validation err: %v, critical err: %v\n", status.ValidationError, status.CriticalError)
		return nil, errors.New(status.ValidationError)
	}

	payload, err := d.engine.GetPayloadV2(payloadId)
	if err != nil {
		return nil, err
	}

	// this is where we could filter/reorder transactions, or mark them for filtering so consensus could be checked

	bytes, err := json.Marshal(payload.ExecutionPayload)
	if err != nil {
		return nil, err
	}
	txs := make([][]byte, 1)
	txs[0] = bytes

	return &abci.ResponsePrepareProposal{
		Txs: txs,
	}, nil
}

func (d *Driver) ProcessProposal(context sdk.Context, proposal *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
	return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
}

// PreBlocker executes before FinalizeBlock - this may be problematic if that fails?
func (d *Driver) PreBlocker(context sdk.Context, block *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
	// the first and only "tx" is the serialized ExecutionPayload (with the actual txs)

	// this actually needs to loop through and find the EngineBlock message

	var tx0 ExecutionPayload
	err := json.Unmarshal(block.Txs[0], &tx0)
	if err != nil {
		return nil, err
	}

	block.GetTime()

	payload, err := d.retryUntilNewPayload(tx0)
	if err != nil {
		return nil, err
	}

	state := ForkChoiceState{
		HeadHash:           common.HexToHash(payload.LatestValidHash),
		SafeBlockHash:      common.HexToHash(payload.LatestValidHash),
		FinalizedBlockHash: common.HexToHash(tx0.ParentHash), // latestHash from the Proposal stage
	}

	// The engine complains when the withdrawals are empty
	withdrawals := []*Withdrawal{
		{
			Index:     "0x0",
			Validator: "0x0",
			Address:   common.Address{}.Hex(),
			Amount:    "0x0",
		},
	}

	attrs := PayloadAttributes{
		Timestamp:             hexutil.Uint64(time.Now().UnixMilli()),
		PrevRandao:            common.Hash{},
		SuggestedFeeRecipient: d.selfAddr,
		Withdrawals:           withdrawals,
	}

	// var choice *ForkchoiceUpdatedResponse
	_, err = d.retryUntilForkchoiceUpdated(&state, &attrs)
	if err != nil {
		return nil, err
	}

	d.state.Height = block.Height // save state later in Commit
	d.state.Size = block.Height   // only one transaction is executed, the ExecutionPayload, so the size is the height

	// in this situation, should state be stored here or wait to commit, like ABCI???

	return &sdk.ResponsePreBlock{}, nil
}

func (d *Driver) Commit(context sdk.Context) {
	if err := d.state.Save(); err != nil {
		panic(err) //???
	}
}

func (d *Driver) retryUntilNewPayload(payload ExecutionPayload) (response *NewPayloadResponse, err error) {
	forever := backoff.NewExponentialBackOff()
	err = backoff.Retry(func() error {
		response, err = d.engine.NewPayloadV2(payload)
		if forever.NextBackOff() > RESET_EXPONENTIAL_BACKOFF {
			forever.Reset()
		}
		if err != nil {
			return err
		}
		return nil
	}, forever)
	if err != nil {
		return nil, err // should not happen, retries forever
	}
	return response, nil
}

func (d *Driver) retryUntilForkchoiceUpdated(state *ForkChoiceState, attrs *PayloadAttributes) (response *ForkchoiceUpdatedResponse, err error) {
	forever := backoff.NewExponentialBackOff()
	err = backoff.Retry(func() error {
		response, err = d.engine.ForkchoiceUpdatedV2(state, attrs)
		if forever.NextBackOff() > RESET_EXPONENTIAL_BACKOFF {
			forever.Reset()
		}
		if err != nil {
			return err
		}
		return nil
	}, forever)
	if err != nil {
		return nil, err // should not happen, retries forever
	}
	return response, nil
}
