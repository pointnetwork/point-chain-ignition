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

func createNDelegatedAmount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DelegatedAmount {
	items := make([]types.DelegatedAmount, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetDelegatedAmount(ctx, items[i])
	}
	return items
}

func TestDelegatedAmountGet(t *testing.T) {
	keeper, ctx := keepertest.LockenomicsKeeper(t)
	items := createNDelegatedAmount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDelegatedAmount(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDelegatedAmountRemove(t *testing.T) {
	keeper, ctx := keepertest.LockenomicsKeeper(t)
	items := createNDelegatedAmount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDelegatedAmount(ctx,
			item.Index,
		)
		_, found := keeper.GetDelegatedAmount(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestDelegatedAmountGetAll(t *testing.T) {
	keeper, ctx := keepertest.LockenomicsKeeper(t)
	items := createNDelegatedAmount(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDelegatedAmount(ctx)),
	)
}
