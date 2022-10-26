package lockenomics_test

import (
	"testing"
	"time"

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
				Index:     "0",
				Start:     uint64(time.Now().Unix()),
				Length:    123534234,
				Delegator: "cosmos1sjllsnramtg3ewxqwwrwjxfgc4n4ef9u0tvx7u",
				Validator: "cosmosvaloper1sjllsnramtg3ewxqwwrwjxfgc4n4ef9u2lcnj0",
			},
			{
				Index:     "1",
				Start:     uint64(time.Now().Unix()),
				Length:    756356987,
				Delegator: "cosmos14lultfckehtszvzw4ehu0apvsr77afvyhgqhwh",
				Validator: "cosmosvaloper14lultfckehtszvzw4ehu0apvsr77afvyju5zzy",
			},
		},
		DelegatedAmountList: []types.DelegatedAmount{
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
	require.ElementsMatch(t, genesisState.DelegatedAmountList, got.DelegatedAmountList)
	// this line is used by starport scaffolding # genesis/test/assert
}
