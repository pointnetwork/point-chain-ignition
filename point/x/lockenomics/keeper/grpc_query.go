package keeper

import (
	"point/x/lockenomics/types"
)

var _ types.QueryServer = Keeper{}
