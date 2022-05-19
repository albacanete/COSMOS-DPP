package cosmosdpp_test

import (
	"testing"

	keepertest "github.com/albacg5/COSMOS-DPP/testutil/keeper"
	"github.com/albacg5/COSMOS-DPP/testutil/nullify"
	"github.com/albacg5/COSMOS-DPP/x/cosmosdpp"
	"github.com/albacg5/COSMOS-DPP/x/cosmosdpp/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmosdppKeeper(t)
	cosmosdpp.InitGenesis(ctx, *k, genesisState)
	got := cosmosdpp.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
