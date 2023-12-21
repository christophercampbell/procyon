package engine

import (
	"cosmossdk.io/core/appmodule"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

const (
	engineVersion uint64 = 0
)

var (
	_ module.AppModuleBasic = AppModuleBasic{}
	_ appmodule.AppModule   = AppModule{}
)

type AppModuleBasic struct {
}

func (a AppModuleBasic) Name() string {
	return ModuleName
}

func (a AppModuleBasic) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {
}

func (a AppModuleBasic) RegisterInterfaces(registry types.InterfaceRegistry) {
}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(context client.Context, mux *runtime.ServeMux) {
}

// NewAppModule creates a new AppModule object
func NewAppModule(keeper *Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         keeper,
	}
}

type AppModule struct {
	AppModuleBasic
	keeper *Keeper
}

func (a AppModule) IsOnePerModuleType() {
}

func (a AppModule) IsAppModule() {
}

//
// App Wiring Setup
//

/*
func init() {
	appmodule.Register(&modulev1.Module{},
		appmodule.Provide(
			ProvideModule,
		))
}

type ModuleInputs struct {
	depinject.In

	KvStoreService store.KVStoreService
	Cdc            codec.Codec
	//LegacyAmino       *codec.LegacyAmino
}

type ModuleOutputs struct {
	depinject.Out

	Keeper *Keeper
	Module appmodule.AppModule
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	k := NewEngineKeeper(in.Cdc, in.KvStoreService)
	m := NewAppModule(k)
	return ModuleOutputs{Keeper: k, Module: m}
}
*/
