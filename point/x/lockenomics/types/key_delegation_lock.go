package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DelegationLockKeyPrefix is the prefix to retrieve all DelegationLock
	DelegationLockKeyPrefix = "DelegationLock/value/"
)

// DelegationLockKey returns the store key to retrieve a DelegationLock from the index fields
func DelegationLockKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
