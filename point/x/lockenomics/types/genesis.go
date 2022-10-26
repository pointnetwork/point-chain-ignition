package types

import (
	"encoding/hex"
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DelegationLockList:  []DelegationLock{},
		DelegatedAmountList: []DelegatedAmount{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in delegationLock
	delegationLockIndexMap := make(map[string]struct{})

	for _, elem := range gs.DelegationLockList {
		byteIndex, _ := hex.DecodeString(elem.Index)
		index := string(byteIndex)
		if _, ok := delegationLockIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for delegationLock")
		}
		delegationLockIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in delegatedAmount
	delegatedAmountIndexMap := make(map[string]struct{})

	for _, elem := range gs.DelegatedAmountList {
		index := string(DelegatedAmountKey(elem.Index))
		if _, ok := delegatedAmountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for delegatedAmount")
		}
		delegatedAmountIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
