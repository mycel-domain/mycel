package furnace_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/mycel-domain/mycel/testutil/keeper"
	"github.com/mycel-domain/mycel/testutil/nullify"
	"github.com/mycel-domain/mycel/x/furnace/module"
	"github.com/mycel-domain/mycel/x/furnace/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		EpochBurnConfig: types.EpochBurnConfig{
			EpochIdentifier: "11",
		},
		BurnAmounts: []types.BurnAmount{
			{
				Index: 0,
			},
			{
				Index: 1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FurnaceKeeper(t)
	furnace.InitGenesis(ctx, *k, genesisState)
	got := furnace.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.EpochBurnConfig, got.EpochBurnConfig)
	require.ElementsMatch(t, genesisState.BurnAmounts, got.BurnAmounts)
	// this line is used by starport scaffolding # genesis/test/assert
}
