package engine

import (
	"testing"

	"cosmossdk.io/core/header"
	storetypes "cosmossdk.io/store/types"
	"github.com/cometbft/cometbft/libs/rand"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	module_testutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite
	ctx    sdk.Context
	encCfg module_testutil.TestEncodingConfig
	keeper *Keeper
}

func (suite *KeeperTestSuite) SetupTest() {
	key := storetypes.NewKVStoreKey(StoreKey)
	testCtx := testutil.DefaultContextWithDB(suite.T(), key, storetypes.NewTransientStoreKey("transient_test"))

	suite.ctx = testCtx.Ctx.WithHeaderInfo(header.Info{})
	suite.encCfg = module_testutil.MakeTestEncodingConfig(AppModuleBasic{})
	suite.keeper = NewEngineKeeper(suite.encCfg.Codec, runtime.NewKVStoreService(key))
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestInitialState() {
	initial, err := suite.keeper.GetState(suite.ctx)
	suite.NoError(err)
	suite.Equal(int64(0), initial.Height)
	suite.Equal(int64(0), initial.NumTxs)
}

func (suite *KeeperTestSuite) TestGetAndSetState() {
	height := rand.Int64()
	numTxs := rand.Int64()
	err := suite.keeper.SetState(suite.ctx,
		&State{
			Height: height,
			NumTxs: numTxs,
		})
	suite.NoError(err)

	read, err := suite.keeper.GetState(suite.ctx)
	suite.NoError(err)

	suite.Equal(height, read.Height)
	suite.Equal(numTxs, read.NumTxs)
}
