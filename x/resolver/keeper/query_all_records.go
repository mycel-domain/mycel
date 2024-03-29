package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	registrytypes "github.com/mycel-domain/mycel/x/registry/types"
	"github.com/mycel-domain/mycel/x/resolver/types"
)

func (k Keeper) AllRecords(goCtx context.Context, req *types.QueryAllRecordsRequest) (*types.QueryAllRecordsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Query domain record
	_, err := k.registryKeeper.GetValidTopLevelDomain(ctx, req.DomainParent)
	if err != nil {
		return nil, err
	}
	secondLevelDomain, err := k.registryKeeper.GetValidSecondLevelDomain(ctx, req.DomainName, req.DomainParent)
	if err != nil {
		return nil, err
	}

	// Convert repeated Record to map
	values := make(map[string]*registrytypes.Record)
	for _, record := range secondLevelDomain.Records {
		key := generateRecordKey(record)
		if key != "" {
			values[key] = record
		}
	}
	return &types.QueryAllRecordsResponse{Values: values}, nil
}

func generateRecordKey(record *registrytypes.Record) string {
	switch {
	case record.GetDnsRecord() != nil:
		return record.GetDnsRecord().DnsRecordType.String()
	case record.GetWalletRecord() != nil:
		return record.GetWalletRecord().WalletRecordType.String()
	case record.GetTextRecord() != nil:
		return record.GetTextRecord().Key
	default:
		return ""
	}
}
