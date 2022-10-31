package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/lockenomics module sentinel errors
var (
	ErrSample            = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrTooLongDelegation = sdkerrors.Register(ModuleName, 1201, "Delegation length should be less than max value")
)
