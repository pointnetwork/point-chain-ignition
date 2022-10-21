package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"point/x/lockenomics/keeper"
	"point/x/lockenomics/types"
)

func SimulateMsgCreateLock(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateLock{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateLock simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateLock simulation not implemented"), nil, nil
	}
}
