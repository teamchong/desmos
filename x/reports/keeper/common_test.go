package keeper_test

import (
	"testing"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	profileskeeper "github.com/desmos-labs/desmos/v4/x/profiles/keeper"
	profilestypes "github.com/desmos-labs/desmos/v4/x/profiles/types"

	postskeeper "github.com/desmos-labs/desmos/v4/x/posts/keeper"
	poststypes "github.com/desmos-labs/desmos/v4/x/posts/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	db "github.com/tendermint/tm-db"

	"github.com/desmos-labs/desmos/v4/app"
	relationshipskeeper "github.com/desmos-labs/desmos/v4/x/relationships/keeper"
	relationshipstypes "github.com/desmos-labs/desmos/v4/x/relationships/types"
	"github.com/desmos-labs/desmos/v4/x/reports/keeper"
	"github.com/desmos-labs/desmos/v4/x/reports/types"
	subspaceskeeper "github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

type KeeperTestsuite struct {
	suite.Suite

	cdc            codec.Codec
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context

	storeKey sdk.StoreKey
	k        keeper.Keeper

	ak profileskeeper.Keeper
	sk subspaceskeeper.Keeper
	rk relationshipskeeper.Keeper
	pk postskeeper.Keeper
}

func (suite *KeeperTestsuite) SetupTest() {
	// Define store keys
	keys := sdk.NewMemoryStoreKeys(
		paramstypes.StoreKey, authtypes.StoreKey,
		profilestypes.StoreKey, relationshipstypes.StoreKey,
		subspacestypes.StoreKey, poststypes.StoreKey,
		types.StoreKey,
	)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	suite.storeKey = keys[types.StoreKey]

	// Create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	for _, key := range keys {
		ms.MountStoreWithDB(key, sdk.StoreTypeIAVL, memDB)
	}
	for _, tKey := range tKeys {
		ms.MountStoreWithDB(tKey, sdk.StoreTypeTransient, memDB)
	}

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	paramsKeeper := paramskeeper.NewKeeper(suite.cdc, suite.legacyAminoCdc, keys[paramstypes.StoreKey], tKeys[paramstypes.TStoreKey])

	// Define keeper
	suite.sk = subspaceskeeper.NewKeeper(suite.cdc, keys[subspacestypes.StoreKey])
	suite.rk = relationshipskeeper.NewKeeper(suite.cdc, keys[relationshipstypes.StoreKey], suite.sk)
	authKeeper := authkeeper.NewAccountKeeper(suite.cdc, keys[authtypes.StoreKey], paramsKeeper.Subspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, app.GetMaccPerms())
	suite.ak = profileskeeper.NewKeeper(suite.cdc, suite.legacyAminoCdc, keys[profilestypes.StoreKey], paramsKeeper.Subspace(profilestypes.DefaultParamsSpace), authKeeper, suite.rk, nil, nil, nil)
	suite.pk = postskeeper.NewKeeper(suite.cdc, keys[poststypes.StoreKey], paramsKeeper.Subspace(poststypes.DefaultParamsSpace), suite.ak, suite.sk, suite.rk)
	suite.k = keeper.NewKeeper(
		suite.cdc,
		suite.storeKey,
		paramsKeeper.Subspace(types.DefaultParamsSpace),
		suite.ak,
		suite.sk,
		suite.rk,
		suite.pk,
	)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestsuite))
}
