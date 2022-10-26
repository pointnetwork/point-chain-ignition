package types_test

import (
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"point/x/lockenomics/types"
)

func TestGenesisState_Validate(t *testing.T) {
	delegatorAddr, _ := sdk.AccAddressFromBech32("cosmos1sjllsnramtg3ewxqwwrwjxfgc4n4ef9u0tvx7u")
	validatorAddr, _ := sdk.ValAddressFromBech32("cosmosvaloper1sjllsnramtg3ewxqwwrwjxfgc4n4ef9u2lcnj0")
	index1 := hex.EncodeToString(types.GetDelegationLockKey(delegatorAddr, validatorAddr))

	delegatorAddr2, _ := sdk.AccAddressFromBech32("cosmos14lultfckehtszvzw4ehu0apvsr77afvyhgqhwh")
	validatorAddr2, _ := sdk.ValAddressFromBech32("cosmosvaloper1qaa9zej9a0ge3ugpx3pxyx602lxh3ztqgfnp42")
	index2 := hex.EncodeToString(types.GetDelegationLockKey(delegatorAddr2, validatorAddr2))

	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				DelegationLockList: []types.DelegationLock{
					{
						Index:     index1,
						Start:     uint64(time.Now().Unix()),
						Length:    123534234,
						Delegator: "cosmos1sjllsnramtg3ewxqwwrwjxfgc4n4ef9u0tvx7u",
						Validator: "cosmosvaloper1sjllsnramtg3ewxqwwrwjxfgc4n4ef9u2lcnj0",
					},
					{
						Index:     index2,
						Start:     uint64(time.Now().Unix()),
						Length:    756356987,
						Delegator: "cosmos14lultfckehtszvzw4ehu0apvsr77afvyhgqhwh",
						Validator: "cosmosvaloper1qaa9zej9a0ge3ugpx3pxyx602lxh3ztqgfnp42",
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated delegationLock",
			genState: &types.GenesisState{
				DelegationLockList: []types.DelegationLock{
					{
						Index:     index1,
						Start:     uint64(time.Now().Unix()),
						Length:    123534234,
						Delegator: "cosmos1sjllsnramtg3ewxqwwrwjxfgc4n4ef9u0tvx7u",
						Validator: "cosmosvaloper1sjllsnramtg3ewxqwwrwjxfgc4n4ef9u2lcnj0",
					},
					{
						Index:     index1,
						Start:     uint64(time.Now().Unix()),
						Length:    123534234,
						Delegator: "cosmos1sjllsnramtg3ewxqwwrwjxfgc4n4ef9u0tvx7u",
						Validator: "cosmosvaloper1sjllsnramtg3ewxqwwrwjxfgc4n4ef9u2lcnj0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated delegatedAmount",
			genState: &types.GenesisState{
				DelegatedAmountList: []types.DelegatedAmount{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
