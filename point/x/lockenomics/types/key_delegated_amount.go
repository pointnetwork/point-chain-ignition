package types

import "encoding/binary"

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
