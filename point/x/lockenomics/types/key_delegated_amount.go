package types

import (
	"encoding/binary"
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

var _ binary.ByteOrder

const (
	// DelegatedAmountKeyPrefix is the prefix to retrieve all DelegatedAmount
	DelegatedAmountKeyPrefix = "DelegatedAmount/value/"
)

// DelegatedAmountKey returns the store key to retrieve a DelegatedAmount from the index fields
func DelegatedAmountKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

// GetDelegationKey creates the key for delegator bond with validator
func GetDelegatedAmountKeyStringFromString(delegator string, validator string) (key string, err error) {
	keyBytes, err := GetDelegatedAmountKeyFromString(delegator, validator)
	return hex.EncodeToString(keyBytes), err
}

// GetDelegationKey creates the key for delegator bond with validator
func GetDelegatedAmountKeyFromString(delegator string, validator string) (key []byte, err error) {
	delAddr, err := sdk.AccAddressFromBech32(delegator)
	if err != nil {
		return nil, err
	}
	valAddr, err := sdk.ValAddressFromBech32(validator)
	if err != nil {
		return nil, err
	}

	return GetDelegatedAmountKey(delAddr, valAddr), nil
}

// GetDelegationKey creates the key for delegator bond with validator
func GetDelegatedAmountKey(delAddr sdk.AccAddress, valAddr sdk.ValAddress) []byte {
	return append(GetDelegatedAmountDelAddrKey(delAddr), address.MustLengthPrefix(valAddr)...)
}

// GetDelegationsKey creates the prefix for a delegator for all validators
func GetDelegatedAmountDelAddrKey(delAddr sdk.AccAddress) []byte {
	return append(KeyPrefix(DelegationLockKeyPrefix), address.MustLengthPrefix(delAddr)...)
}
