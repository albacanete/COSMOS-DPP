package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/albacg5/COSMOS-DPP/testutil/keeper"
	"github.com/albacg5/COSMOS-DPP/x/cosmosdpp/keeper"
	"github.com/albacg5/COSMOS-DPP/x/cosmosdpp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmosdppKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
