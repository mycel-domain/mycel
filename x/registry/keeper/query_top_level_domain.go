package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/mycel-domain/mycel/x/registry/types"
)

func (k Keeper) TopLevelDomainAll(ctx context.Context, req *types.QueryAllTopLevelDomainRequest) (*types.QueryAllTopLevelDomainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var topLevelDomains []types.TopLevelDomain

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	topLevelDomainStore := prefix.NewStore(store, types.KeyPrefix(types.TopLevelDomainKeyPrefix))

	pageRes, err := query.Paginate(topLevelDomainStore, req.Pagination, func(key []byte, value []byte) error {
		var topLevelDomain types.TopLevelDomain
		if err := k.cdc.Unmarshal(value, &topLevelDomain); err != nil {
			return err
		}

		topLevelDomains = append(topLevelDomains, topLevelDomain)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTopLevelDomainResponse{TopLevelDomain: topLevelDomains, Pagination: pageRes}, nil
}

func (k Keeper) TopLevelDomain(ctx context.Context, req *types.QueryGetTopLevelDomainRequest) (*types.QueryGetTopLevelDomainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetTopLevelDomain(
		ctx,
		req.Name,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTopLevelDomainResponse{TopLevelDomain: val}, nil
}
