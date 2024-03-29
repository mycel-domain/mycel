package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

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

	value := secondLevelDomain.GetDnsRecord(req.DnsRecordType)
	recordType := registrytypes.DnsRecordType(registrytypes.DnsRecordType_value[req.DnsRecordType])

	return &types.QueryDnsRecordResponse{
		Value: &registrytypes.DnsRecord{DnsRecordType: recordType, Value: value},
	}, nil
}
