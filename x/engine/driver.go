package engine

import (
	"encoding/binary"
	"fmt"

	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Driver implements the ABCI methods needed to driver the external engine
type Driver struct {
	keeper     *Keeper
	commitNext func(context sdk.Context, state *State)
}

func NewEngineDriver(keeper *Keeper) *Driver {
	return &Driver{keeper: keeper}
}

func (d *Driver) mustGetState(context sdk.Context) *State {
	state, err := d.keeper.GetState(context)
	if err != nil {
		panic(err)
	}
	return state
}

func (d *Driver) PrepareProposal(context sdk.Context, proposal *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
	fmt.Println(">>> PREPARE <<<")
	return &abci.ResponsePrepareProposal{Txs: proposal.Txs}, nil
}

func (d *Driver) ProcessProposal(context sdk.Context, proposal *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
	fmt.Println(">>> PROCESS <<<")
	return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
}

func (d *Driver) FinalizeBlock(context sdk.Context, block *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
	state := d.mustGetState(context)
	fmt.Println(">>> (PRE) FINALIZE BLOCK <<<")
	// if there is an engine execution tx, do it

	//d.commitNext = func(c sdk.Context, s *State) {
	//	d.keeper.SetState(c, s)
	//}

	// this really needs to happen in the commit function
	d.keeper.SetState(context, &State{
		Height: state.Height + 1,
		NumTxs: state.Height + 1,
	})

	return &sdk.ResponsePreBlock{}, nil
}

func appHash(s *State) []byte {
	bytes := make([]byte, 8)
	binary.PutVarint(bytes, s.Height+s.NumTxs)
	return bytes
}

func (d *Driver) Commit(context sdk.Context) {
	fmt.Println(">>> (PRE) COMMIT <<<")

	/*
		err := d.keeper.SetState(context, d.nextState)
		if err != nil {
			panic(err)
		}

	*/
}
