package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "point/testutil/keeper"
	"point/x/point/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PointKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
