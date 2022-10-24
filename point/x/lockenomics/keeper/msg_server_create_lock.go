package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"point/x/lockenomics/types"
)

func (k msgServer) CreateLock(goCtx context.Context, msg *types.MsgCreateLock) (*types.MsgCreateLockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	/*valAddr, valErr := sdk.ValAddressFromBech32(msg.Validator)
	if valErr != nil {
		return nil, valErr
	}

	validator, found := k.stakingKeeper.GetValidator(ctx, valAddr)

	if !found {
		return nil, stakingTypes.ErrNoValidatorFound
	}

	delegatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	err := k.Keeper.SetDelegationLock(ctx, )
	if err != nil {
		return nil, err
	}*/

	// TODO: Handling the message end
	_ = ctx

	return &types.MsgCreateLockResponse{}, nil
}
