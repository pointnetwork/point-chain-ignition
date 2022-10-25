package types

import (
	"encoding/binary"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

var _ binary.ByteOrder

const (
	// DelegationLockKeyPrefix is the prefix to retrieve all DelegationLock
	DelegationLockKeyPrefix = "DelegationLock/value/"
)

// DelegationLockKey returns the store key to retrieve a DelegationLock from the index fields
func DelegationLockIndexKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

// GetDelegationKey creates the key for delegator bond with validator
// VALUE: staking/Delegation
func GetDelegationLockKey(delAddr sdk.AccAddress, valAddr sdk.ValAddress) []byte {
	return append(GetDelegationsLockKey(delAddr), address.MustLengthPrefix(valAddr)...)
}

// GetDelegationsKey creates the prefix for a delegator for all validators
func GetDelegationsLockKey(delAddr sdk.AccAddress) []byte {
	return append(KeyPrefix(DelegationLockKeyPrefix), address.MustLengthPrefix(delAddr)...)
}
