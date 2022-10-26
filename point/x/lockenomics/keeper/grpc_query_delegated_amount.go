package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"point/x/lockenomics/types"
)

func (k Keeper) DelegatedAmountAll(c context.Context, req *types.QueryAllDelegatedAmountRequest) (*types.QueryAllDelegatedAmountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var delegatedAmounts []types.DelegatedAmount
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	delegatedAmountStore := prefix.NewStore(store, types.KeyPrefix(types.DelegatedAmountKeyPrefix))

	pageRes, err := query.Paginate(delegatedAmountStore, req.Pagination, func(key []byte, value []byte) error {
		var delegatedAmount types.DelegatedAmount
		if err := k.cdc.Unmarshal(value, &delegatedAmount); err != nil {
			return err
		}

		delegatedAmounts = append(delegatedAmounts, delegatedAmount)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDelegatedAmountResponse{DelegatedAmount: delegatedAmounts, Pagination: pageRes}, nil
}

func (k Keeper) DelegatedAmount(c context.Context, req *types.QueryGetDelegatedAmountRequest) (*types.QueryGetDelegatedAmountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetDelegatedAmount(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDelegatedAmountResponse{DelegatedAmount: val}, nil
}
