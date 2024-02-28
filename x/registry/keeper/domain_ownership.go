package keeper

import (
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/x/registry/types"
)

// SetDomainOwnership set a specific domainOwnership in the store from its index
func (k Keeper) SetDomainOwnership(ctx sdk.Context, domainOwnership types.DomainOwnership) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DomainOwnershipKeyPrefix))

	b := k.cdc.MustMarshal(&domainOwnership)
	store.Set(types.DomainOwnershipKey(
		domainOwnership.Owner,
	), b)
}

// GetDomainOwnership returns a domainOwnership from its index
func (k Keeper) GetDomainOwnership(
	ctx sdk.Context,
	owner string,
) (val types.DomainOwnership, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DomainOwnershipKeyPrefix))

	b := store.Get(types.DomainOwnershipKey(
		owner,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDomainOwnership removes a domainOwnership from the store
func (k Keeper) RemoveDomainOwnership(
	ctx sdk.Context,
	owner string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DomainOwnershipKeyPrefix))

	store.Delete(types.DomainOwnershipKey(
		owner,
	))
}

// GetAllDomainOwnership returns all domainOwnership
func (k Keeper) GetAllDomainOwnership(ctx sdk.Context) (list []types.DomainOwnership) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DomainOwnershipKeyPrefix))

	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DomainOwnership
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Append to owned domain
func (k Keeper) AppendToOwnedDomain(ctx sdk.Context, owner string, name string, parent string) {
	domainOwnership, found := k.GetDomainOwnership(ctx, owner)
	if found {
		domainOwnership.Domains = append(domainOwnership.Domains, &types.OwnedDomain{Name: name, Parent: parent})
		k.SetDomainOwnership(ctx, domainOwnership)
	} else {
		k.SetDomainOwnership(ctx, types.DomainOwnership{Owner: owner, Domains: []*types.OwnedDomain{{Name: name, Parent: parent}}})
	}
}
