package keeper

import (
	"github.com/airchains-network/airsettle/x/airsettle/types"
)

var _ types.QueryServer = Keeper{}
