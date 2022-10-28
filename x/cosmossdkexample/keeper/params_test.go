package keeper_test

import (
	"testing"

	testkeeper "github.com/bigzoro/cosmos-sdk-example/testutil/keeper"
	"github.com/bigzoro/cosmos-sdk-example/x/cosmossdkexample/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CosmossdkexampleKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
