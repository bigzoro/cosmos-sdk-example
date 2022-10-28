package cosmossdkexample_test

import (
	"testing"

	keepertest "github.com/bigzoro/cosmos-sdk-example/testutil/keeper"
	"github.com/bigzoro/cosmos-sdk-example/testutil/nullify"
	"github.com/bigzoro/cosmos-sdk-example/x/cosmossdkexample"
	"github.com/bigzoro/cosmos-sdk-example/x/cosmossdkexample/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmossdkexampleKeeper(t)
	cosmossdkexample.InitGenesis(ctx, *k, genesisState)
	got := cosmossdkexample.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
