package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"point/x/lockenomics/types"
)

func (k msgServer) CreateLock(goCtx context.Context, msg *types.MsgCreateLock) (*types.MsgCreateLockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_, valErr := sdk.ValAddressFromBech32(msg.Validator)
	if valErr != nil {
		return nil, valErr
	}

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	delegationLock := types.DelegationLock{
		Start:     12345,
		Length:    msg.Lenght,
		Delegator: msg.Creator,
		Validator: msg.Validator,
	}

	k.Keeper.SetDelegationLock(ctx, delegationLock)
	if err != nil {
		return nil, err
	}

	// TODO: Handling the message end
	_ = ctx

	return &types.MsgCreateLockResponse{}, nil
}
