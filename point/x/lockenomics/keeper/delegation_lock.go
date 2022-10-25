package keeper

import (
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"point/x/lockenomics/types"
)

// SetDelegationLock set a specific delegationLock in the store from its index
func (k Keeper) SetDelegationLock(ctx sdk.Context, delegationLock types.DelegationLock) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegationLockKeyPrefix))
	b := k.cdc.MustMarshal(&delegationLock)
	delegatorAddress := sdk.MustAccAddressFromBech32(delegationLock.Delegator)
	validatorAddress, err := sdk.ValAddressFromBech32(delegationLock.Validator)
	if err != nil {
		panic(err)
	}
	key := types.GetDelegationLockKey(delegatorAddress, validatorAddress)
	store.Set(key, b)
}

// GetDelegationLock returns a delegationLock by Delegator And Validator
func (k Keeper) GetDelegationLockByDelegatorAndValidator(
	ctx sdk.Context,
	delegator sdk.AccAddress,
	validator sdk.ValAddress,
) (val types.DelegationLock, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegationLockKeyPrefix))
	key := types.GetDelegationLockKey(delegator, validator)
	b := store.Get(key)
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetDelegationLock returns a delegationLock from its index
func (k Keeper) GetDelegationLock(
	ctx sdk.Context,
	index string,

) (val types.DelegationLock, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegationLockKeyPrefix))
	byteIndex, _ := hex.DecodeString(index)
	b := store.Get(byteIndex)
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDelegationLock removes a delegationLock from the store
func (k Keeper) RemoveDelegationLock(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegationLockKeyPrefix))
	byteIndex, _ := hex.DecodeString(index)
	store.Delete(byteIndex)
}

// GetAllDelegationLock returns all delegationLock
func (k Keeper) GetAllDelegationLock(ctx sdk.Context) (list []types.DelegationLock) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegationLockKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DelegationLock
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
