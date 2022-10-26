package keeper

import (
	"context"
	"encoding/hex"
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

	delAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()

	if now < 0 {
		return nil, types.ErrInvalidStartParams
	}
	nowUint := uint64(now)

	if msg.Lenght < 0 {
		return nil, types.ErrInvalidLengthParams
	}
	lengthUint := uint64(msg.Lenght)

	key := types.GetDelegationLockKey(delAddress, valAddr)

	delegationLock := types.DelegationLock{
		Index:     hex.EncodeToString(key),
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
