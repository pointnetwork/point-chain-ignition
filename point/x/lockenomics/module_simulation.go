package lockenomics

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"point/testutil/sample"
	lockenomicssimulation "point/x/lockenomics/simulation"
	"point/x/lockenomics/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = lockenomicssimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateLock = "op_weight_msg_create_lock"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateLock int = 100

	opWeightMsgCreateDelegatedAmount = "op_weight_msg_delegated_amount"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDelegatedAmount int = 100

	opWeightMsgUpdateDelegatedAmount = "op_weight_msg_delegated_amount"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDelegatedAmount int = 100

	opWeightMsgDeleteDelegatedAmount = "op_weight_msg_delegated_amount"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteDelegatedAmount int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	lockenomicsGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		DelegatedAmountList: []types.DelegatedAmount{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&lockenomicsGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateLock int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateLock, &weightMsgCreateLock, nil,
		func(_ *rand.Rand) {
			weightMsgCreateLock = defaultWeightMsgCreateLock
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateLock,
		lockenomicssimulation.SimulateMsgCreateLock(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
