package keeper_test

import (
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
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].Delegator = "cosmos156gqf9837u7d4c4678yt3rl4ls9c5vuuxyhkw6"
		items[i].Validator = "cosmosvaloper156gqf9837u7d4c4678yt3rl4ls9c5vuursrrzf"
		items[i].Start = uint64(time.Now().Unix())
		items[i].Length = 50000000
		keeper.SetDelegationLock(ctx, items[i])
	}
	return items
}

func TestDelegationLockGet(t *testing.T) {
	keeper, ctx := keepertest.LockenomicsKeeper(t)
	items := createNDelegationLock(keeper, ctx, 1)
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
	items := createNDelegationLock(keeper, ctx, 10)
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
	items := createNDelegationLock(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDelegationLock(ctx)),
	)
}
