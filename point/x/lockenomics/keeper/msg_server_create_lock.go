package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"point/x/lockenomics/types"
)

func (k msgServer) CreateLock(goCtx context.Context, msg *types.MsgCreateLock) (*types.MsgCreateLockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateLockResponse{}, nil
}
