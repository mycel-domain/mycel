package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "mycel/testutil/keeper"
	"mycel/x/registry/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RegistryKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
