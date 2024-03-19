package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "github.com/mycel-domain/mycel/testutil/keeper"
	"github.com/mycel-domain/mycel/testutil/nullify"
	"github.com/mycel-domain/mycel/x/registry/keeper"
	"github.com/mycel-domain/mycel/x/registry/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDomainOwnership(keeper keeper.Keeper, ctx context.Context, n int) []types.DomainOwnership {
	items := make([]types.DomainOwnership, n)
	for i := range items {
		items[i].Owner = strconv.Itoa(i)

		keeper.SetDomainOwnership(ctx, items[i])
	}
	return items
}

func TestDomainOwnershipGet(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNDomainOwnership(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDomainOwnership(ctx,
			item.Owner,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestDomainOwnershipRemove(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNDomainOwnership(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDomainOwnership(ctx,
			item.Owner,
		)
		_, found := keeper.GetDomainOwnership(ctx,
			item.Owner,
		)
		require.False(t, found)
	}
}

func TestDomainOwnershipGetAll(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNDomainOwnership(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDomainOwnership(ctx)),
	)
}
