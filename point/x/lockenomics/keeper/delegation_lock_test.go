package keeper_test

import (
	"strconv"
	"testing"

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

		keeper.SetDelegationLock(ctx, items[i])
	}
	return items
}

func TestDelegationLockGet(t *testing.T) {
	keeper, ctx := keepertest.LockenomicsKeeper(t)
	items := createNDelegationLock(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDelegationLock(ctx,
			item.Index,
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
