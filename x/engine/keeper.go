package engine

import (
	"context"

	"cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
)

type Keeper struct {
	cdc          codec.BinaryCodec
	storeService store.KVStoreService
}

func NewEngineKeeper(cdc codec.BinaryCodec,
	storeService store.KVStoreService) *Keeper {
	return &Keeper{
		cdc:          cdc,
		storeService: storeService,
	}
}

func (k *Keeper) GetState(ctx context.Context) (*State, error) {
	kv := k.storeService.OpenKVStore(ctx)
	bytes, err := kv.Get(stateStorageKey)
	if err != nil {
		return nil, err
	}
	var state State
	if len(bytes) == 0 {
		return &state, nil
	}
	err = k.cdc.Unmarshal(bytes, &state)
	if err != nil {
		return nil, err
	}
	return &state, nil
}

func (k *Keeper) SetState(ctx context.Context, state *State) error {
	bytes, err := k.cdc.Marshal(state)
	if err != nil {
		return err
	}
	kv := k.storeService.OpenKVStore(ctx)
	return kv.Set(stateStorageKey, bytes)
}
