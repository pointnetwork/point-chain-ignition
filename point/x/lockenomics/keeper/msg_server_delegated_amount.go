package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"point/x/lockenomics/types"
)

func (k msgServer) CreateDelegatedAmount(goCtx context.Context, msg *types.MsgCreateDelegatedAmount) (*types.MsgCreateDelegatedAmountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetDelegatedAmount(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var delegatedAmount = types.DelegatedAmount{
		Creator:   msg.Creator,
		Index:     msg.Index,
		Delegator: msg.Delegator,
		Validator: msg.Validator,
		Amount:    msg.Amount,
	}

	k.SetDelegatedAmount(
		ctx,
		delegatedAmount,
	)
	return &types.MsgCreateDelegatedAmountResponse{}, nil
}

func (k msgServer) UpdateDelegatedAmount(goCtx context.Context, msg *types.MsgUpdateDelegatedAmount) (*types.MsgUpdateDelegatedAmountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDelegatedAmount(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var delegatedAmount = types.DelegatedAmount{
		Creator:   msg.Creator,
		Index:     msg.Index,
		Delegator: msg.Delegator,
		Validator: msg.Validator,
		Amount:    msg.Amount,
	}

	k.SetDelegatedAmount(ctx, delegatedAmount)

	return &types.MsgUpdateDelegatedAmountResponse{}, nil
}

func (k msgServer) DeleteDelegatedAmount(goCtx context.Context, msg *types.MsgDeleteDelegatedAmount) (*types.MsgDeleteDelegatedAmountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDelegatedAmount(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveDelegatedAmount(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteDelegatedAmountResponse{}, nil
}
