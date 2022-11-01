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
	delegators := [7]string{"cosmos156gqf9837u7d4c4678yt3rl4ls9c5vuuxyhkw6",
		"cosmos14lultfckehtszvzw4ehu0apvsr77afvyhgqhwh",
		"cosmos19lss6zgdh5vvcpjhfftdghrpsw7a4434ut4md0",
		"cosmos196ax4vc0lwpxndu9dyhvca7jhxp70rmcfhxsrt",
		"cosmos1z8zjv3lntpwxua0rtpvgrcwl0nm0tltgyuy0nd",
		"cosmos1qaa9zej9a0ge3ugpx3pxyx602lxh3ztqda85ee",
		"cosmos1tflk30mq5vgqjdly92kkhhq3raev2hnzldd74z",
	}
	validators := [7]string{"cosmosvaloper156gqf9837u7d4c4678yt3rl4ls9c5vuursrrzf",
		"cosmosvaloper14lultfckehtszvzw4ehu0apvsr77afvyju5zzy",
		"cosmosvaloper19lss6zgdh5vvcpjhfftdghrpsw7a4434elpwpu",
		"cosmosvaloper196ax4vc0lwpxndu9dyhvca7jhxp70rmcvrj90c",
		"cosmosvaloper1z8zjv3lntpwxua0rtpvgrcwl0nm0tltgpgs6l7",
		"cosmosvaloper1qaa9zej9a0ge3ugpx3pxyx602lxh3ztqgfnp42",
		"cosmosvaloper1tflk30mq5vgqjdly92kkhhq3raev2hnz6eete3",
	}
	for i := range items {
		delegAddr, _ := sdk.AccAddressFromBech32(delegators[i])
		valAddr, _ := sdk.ValAddressFromBech32(validators[i])
		items[i].Index = types.GetDelegationLockKeyString(delegAddr, valAddr)
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
