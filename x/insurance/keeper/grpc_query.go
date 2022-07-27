package keeper

import (
	"github.com/DenisCherkasskikh/insurance/x/insurance/types"
)

var _ types.QueryServer = Keeper{}
