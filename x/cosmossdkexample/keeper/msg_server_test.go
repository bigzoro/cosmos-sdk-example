package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/bigzoro/cosmos-sdk-example/testutil/keeper"
	"github.com/bigzoro/cosmos-sdk-example/x/cosmossdkexample/keeper"
	"github.com/bigzoro/cosmos-sdk-example/x/cosmossdkexample/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmossdkexampleKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
