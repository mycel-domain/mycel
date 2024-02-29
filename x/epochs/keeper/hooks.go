package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AfterEpochEnd epoch hook
func (k Keeper) AfterEpochEnd(ctx sdk.Context, identifier string, epochNumber int64) {
	if k.hooks != nil {
		k.hooks.AfterEpochEnd(ctx, identifier, epochNumber)
	}
}

// BeforeEpochStart epoch hook
func (k Keeper) BeforeEpochStart(ctx sdk.Context, identifier string, epochNumber int64) {
	if k.hooks != nil {
		k.hooks.BeforeEpochStart(ctx, identifier, epochNumber)
	}
}
