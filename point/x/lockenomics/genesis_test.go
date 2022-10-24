package lockenomics_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "point/testutil/keeper"
	"point/testutil/nullify"
	"point/x/lockenomics"
	"point/x/lockenomics/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DelegationLockList: []types.DelegationLock{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LockenomicsKeeper(t)
	lockenomics.InitGenesis(ctx, *k, genesisState)
	got := lockenomics.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DelegationLockList, got.DelegationLockList)
	// this line is used by starport scaffolding # genesis/test/assert
}
