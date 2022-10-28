package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"point/x/lockenomics/types"
)

// SetDelegatedAmount set a specific delegatedAmount in the store from its index
func (k Keeper) SetDelegatedAmount(ctx sdk.Context, delegatedAmount types.DelegatedAmount) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegatedAmountKeyPrefix))
	index, err := types.GetDelegatedAmountKeyStringFromString(delegatedAmount.Delegator, delegatedAmount.Validator)
	if err != nil {
		return err
	}
	delegatedAmount.Index = index
	b := k.cdc.MustMarshal(&delegatedAmount)
	store.Set(types.DelegatedAmountKey(
		delegatedAmount.Index,
	), b)
	return nil
}

// GetDelegatedAmount returns a delegatedAmount from its index
func (k Keeper) GetDelegatedAmount(
	ctx sdk.Context,
	index string,
) (val types.DelegatedAmount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegatedAmountKeyPrefix))
	keyBytes := types.DelegatedAmountKey(index)
	if keyBytes == nil {
		return val, false
	}
	b := store.Get(keyBytes)
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDelegatedAmount removes a delegatedAmount from the store
func (k Keeper) RemoveDelegatedAmount(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegatedAmountKeyPrefix))
	store.Delete(types.DelegatedAmountKey(
		index,
	))
}

// GetAllDelegatedAmount returns all delegatedAmount
func (k Keeper) GetAllDelegatedAmount(ctx sdk.Context) (list []types.DelegatedAmount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegatedAmountKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DelegatedAmount
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
