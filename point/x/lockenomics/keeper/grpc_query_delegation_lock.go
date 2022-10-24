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

func (k Keeper) DelegationLockAll(c context.Context, req *types.QueryAllDelegationLockRequest) (*types.QueryAllDelegationLockResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var delegationLocks []types.DelegationLock
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	delegationLockStore := prefix.NewStore(store, types.KeyPrefix(types.DelegationLockKeyPrefix))

	pageRes, err := query.Paginate(delegationLockStore, req.Pagination, func(key []byte, value []byte) error {
		var delegationLock types.DelegationLock
		if err := k.cdc.Unmarshal(value, &delegationLock); err != nil {
			return err
		}

		delegationLocks = append(delegationLocks, delegationLock)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDelegationLockResponse{DelegationLock: delegationLocks, Pagination: pageRes}, nil
}

func (k Keeper) DelegationLock(c context.Context, req *types.QueryGetDelegationLockRequest) (*types.QueryGetDelegationLockResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetDelegationLock(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDelegationLockResponse{DelegationLock: val}, nil
}
