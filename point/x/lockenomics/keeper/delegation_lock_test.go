package keeper_test

import (
	"encoding/hex"
	"strconv"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "point/testutil/keeper"
	"point/testutil/nullify"
	"point/x/lockenomics/keeper"
	"point/x/lockenomics/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDelegationLock(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DelegationLock {
	items := make([]types.DelegationLock, n)
	delegators := [2]string{"cosmos156gqf9837u7d4c4678yt3rl4ls9c5vuuxyhkw6",
		"cosmos14lultfckehtszvzw4ehu0apvsr77afvyhgqhwh"}
	validators := [2]string{"cosmosvaloper156gqf9837u7d4c4678yt3rl4ls9c5vuursrrzf",
		"cosmosvaloper14lultfckehtszvzw4ehu0apvsr77afvyju5zzy"}
	for i := range items {
		delegAddr, _ := sdk.AccAddressFromBech32(delegators[i])
		valAddr, _ := sdk.ValAddressFromBech32(validators[i])
		key := types.GetDelegationLockKey(delegAddr, valAddr)
		items[i].Index = hex.EncodeToString(key)
		items[i].Delegator = delegators[i]
		items[i].Validator = validators[i]
		items[i].Start = uint64(time.Now().Unix())
		items[i].Length = 50000000
		keeper.SetDelegationLock(ctx, items[i])
	}
	return items
}

func TestDelegationLockGet(t *testing.T) {
	keeper, ctx := keepertest.LockenomicsKeeper(t)
	items := createNDelegationLock(keeper, ctx, 2)
	for _, item := range items {
		delegatorAddress := sdk.MustAccAddressFromBech32(item.Delegator)
		validatorAddress, _ := sdk.ValAddressFromBech32(item.Validator)

		rst, found := keeper.GetDelegationLockByDelegatorAndValidator(ctx,
			delegatorAddress,
			validatorAddress,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDelegationLockRemove(t *testing.T) {
	keeper, ctx := keepertest.LockenomicsKeeper(t)
	items := createNDelegationLock(keeper, ctx, 2)
	for _, item := range items {
		keeper.RemoveDelegationLock(ctx,
			item.Index,
		)
		_, found := keeper.GetDelegationLock(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestDelegationLockGetAll(t *testing.T) {
	keeper, ctx := keepertest.LockenomicsKeeper(t)
	items := createNDelegationLock(keeper, ctx, 2)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDelegationLock(ctx)),
	)
}
