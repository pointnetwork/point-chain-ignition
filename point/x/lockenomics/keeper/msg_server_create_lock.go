package keeper

import (
	"context"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"point/x/lockenomics/types"
)

func (k msgServer) CreateLock(goCtx context.Context, msg *types.MsgCreateLock) (*types.MsgCreateLockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	valAddr, valErr := sdk.ValAddressFromBech32(msg.Validator)
	if valErr != nil {
		return nil, valErr
	}
	_, found := k.stakingKeeper.GetValidator(ctx, valAddr)
	if !found {
		return nil, stakingTypes.ErrNoValidatorFound
	}
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	var UINT64_MAX int64 = 18446744073709551615
	var UINT64_MIN int64 = 0

	now := time.Now().Unix()

	if now > UINT64_MAX || now < UINT64_MIN {
		return nil, types.ErrIntOverflowDelegationLock
	}
	nowUint := uint64(now)
	lenghtInt64 := int64(msg.Lenght)
	if lenghtInt64 > UINT64_MAX || lenghtInt64 < UINT64_MIN {
		return nil, types.ErrIntOverflowParams
	}
	lengthUint := uint64(msg.Lenght)

	delegationLock := types.DelegationLock{
		Start:     nowUint,
		Length:    lengthUint,
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
