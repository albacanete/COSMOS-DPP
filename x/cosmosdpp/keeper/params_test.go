package keeper_test

import (
	"testing"

	testkeeper "github.com/albacg5/COSMOS-DPP/testutil/keeper"
	"github.com/albacg5/COSMOS-DPP/x/cosmosdpp/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CosmosdppKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
