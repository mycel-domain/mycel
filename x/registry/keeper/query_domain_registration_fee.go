package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/x/registry/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DomainRegistrationFee(goCtx context.Context, req *types.QueryDomainRegistrationFeeRequest) (*types.QueryDomainRegistrationFeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx
	domain := types.Domain{Name: req.Name, Parent: req.Parent}
	fee := domain.GetRegistrationFee()

	return &types.QueryDomainRegistrationFeeResponse{Fee: fee}, nil
}
