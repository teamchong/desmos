package simulation

// DONTCOVER

import (
	"math/rand"

	"github.com/desmos-labs/desmos/testutil/simtesting"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/keeper"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

// SimulateMsgBlockUser tests and runs a single msg block user
// nolint: funlen
func SimulateMsgBlockUser(k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		blocker, blocked, subspaceID, skip := randomUserBlocksFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, ""), nil, nil
		}

		msg := types.NewMsgBlockUser(
			blocker.Address.String(),
			blocked.String(),
			"reason",
			subspaceID,
		)
		err = simtesting.SendMsg(r, app, ak, bk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{blocker.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, ""), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// randomUserBlocksFields returns random block user fields
func randomUserBlocksFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (simtypes.Account, sdk.AccAddress, string, bool) {
	if len(accs) == 0 {
		return simtypes.Account{}, nil, "", true
	}

	// Get random accounts
	blocker, _ := simtypes.RandomAcc(r, accs)
	blocked, _ := simtypes.RandomAcc(r, accs)

	subspace := RandomSubspace(r)

	if err := k.CheckUserPermissionsInSubspace(ctx, subspace, blocker.Address.String()); err != nil {
		return simtypes.Account{}, nil, "", true
	}

	if err := k.CheckUserPermissionsInSubspace(ctx, subspace, blocked.Address.String()); err != nil {
		return simtypes.Account{}, nil, "", true
	}

	// Skip if the blocker and blocked user are equals
	if blocker.Equals(blocked) {
		return simtypes.Account{}, nil, "", true
	}

	// Skip if the blocker does not have a profile
	if !k.HasProfile(ctx, blocker.Address.String()) {
		return simtypes.Account{}, nil, "", true
	}

	// skip if user block already exists
	userBlocks := k.GetUserBlocks(ctx, blocker.Address.String())
	for _, userBlock := range userBlocks {
		if userBlock.Blocked == blocked.Address.String() {
			return simtypes.Account{}, nil, "", true
		}
	}

	return blocker, blocked.Address, subspace, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgUnblockUser tests and runs a single msg unblock user
func SimulateMsgUnblockUser(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		acc, userBlock, skip := randomUnblockUserFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, ""), nil, nil
		}

		msg := types.NewMsgUnblockUser(
			acc.Address.String(),
			userBlock.Blocked,
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		)
		err = simtesting.SendMsg(r, app, ak, bk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{acc.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, ""), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// randomUnblockUserFields returns random unblock user fields
func randomUnblockUserFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (simtypes.Account, types.UserBlock, bool) {
	if len(accs) == 0 {
		return simtypes.Account{}, types.UserBlock{}, true
	}

	// Get random accounts
	user, _ := simtypes.RandomAcc(r, accs)
	userBlocks := k.GetUserBlocks(ctx, user.Address.String())

	// skip the test if the user has no userBlocks
	if len(userBlocks) == 0 {
		return simtypes.Account{}, types.UserBlock{}, true
	}

	return user, RandomUserBlock(r, userBlocks), false
}
