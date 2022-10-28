package keeper

import (
	"github.com/bigzoro/cosmos-sdk-example/x/cosmossdkexample/types"
)

var _ types.QueryServer = Keeper{}
