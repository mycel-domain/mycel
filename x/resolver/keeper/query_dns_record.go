package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	registrytypes "github.com/mycel-domain/mycel/x/registry/types"
	"github.com/mycel-domain/mycel/x/resolver/types"
)

func (k Keeper) DnsRecord(goCtx context.Context, req *types.QueryDnsRecordRequest) (*types.QueryDnsRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate request parameters
	err := registrytypes.ValidateDnsRecordType(req.DnsRecordType)
	if err != nil {
		return nil, err
	}

	// Query domain record
	_, err = k.registryKeeper.GetValidTopLevelDomain(ctx, req.DomainParent)
	if err != nil {
		return nil, err
	}
	secondLevelDomain, err := k.registryKeeper.GetValidSecondLevelDomain(ctx, req.DomainName, req.DomainParent)
	if err != nil {
		return nil, err
	}

	value := secondLevelDomain.Records[req.DnsRecordType].GetDnsRecord()

	return &types.QueryDnsRecordResponse{Value: value}, nil
}
