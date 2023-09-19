package keeper

import (
	"errors"
	"fmt"

	"github.com/mycel-domain/mycel/x/registry/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Get is domain already taken
func (k Keeper) GetIsDomainAlreadyTaken(ctx sdk.Context, domain types.SecondLevelDomain) (isDomainAlreadyTaken bool) {
	_, isDomainAlreadyTaken = k.GetSecondLevelDomain(ctx, domain.Name, domain.Parent)
	return isDomainAlreadyTaken
}

// Get is parent domain exist
func (k Keeper) GetIsParentDomainExist(ctx sdk.Context, domain types.SecondLevelDomain) (isParentDomainExist bool) {
	parent := domain.ParseParent()
	_, isParentDomainExist = k.GetTopLevelDomain(ctx, parent)
	return isParentDomainExist
}

// Get is domain already taken
func (k Keeper) GetIsTopLevelDomainAlreadyTaken(ctx sdk.Context, domain types.TopLevelDomain) (isDomainAlreadyTaken bool) {
	_, isDomainAlreadyTaken = k.GetTopLevelDomain(ctx, domain.Name)
	return isDomainAlreadyTaken
}

// Validate TLD registration
func (k Keeper) ValidateRegisterTLD(ctx sdk.Context, domain types.TopLevelDomain) (err error) {
	// TODO: Validate TLD
	// TODO: Is Staked enough to register TLD
	return err
}

// Validate SLD registration
func (k Keeper) ValidateRegisterSLD(ctx sdk.Context, domain types.SecondLevelDomain) (err error) {
	isParentDomainExist := k.GetIsParentDomainExist(ctx, domain)
	if !isParentDomainExist {
		err = errorsmod.Wrapf(errors.New(domain.Parent),
			types.ErrParentDomainDoesNotExist.Error())
	}

	return err
}

// Validate subdomain GetRegistrationFee
func (k Keeper) ValidateRegsiterSubdomain(ctx sdk.Context, domain types.SecondLevelDomain) (err error) {
	isParentDomainExist := k.GetIsParentDomainExist(ctx, domain)
	if !isParentDomainExist {
		err = errorsmod.Wrapf(errors.New(domain.Parent),
			types.ErrParentDomainDoesNotExist.Error())
	}
	return err
}

// Validate second-level-domain
func (k Keeper) ValidateSecondLevelDomain(ctx sdk.Context, domain types.SecondLevelDomain) (err error) {
	// Type check
	err = domain.Validate()
	if err != nil {
		return err
	}
	// Check if domain is already taken
	isDomainAlreadyTaken := k.GetIsDomainAlreadyTaken(ctx, domain)
	if isDomainAlreadyTaken {
		err = errorsmod.Wrapf(errors.New(fmt.Sprintf("%s.%s", domain.Name, domain.Parent)),
			types.ErrDomainIsAlreadyTaken.Error())
		return err
	}

	// Validate SLD
	err = k.ValidateRegisterSLD(ctx, domain)
	if err != nil {
		return err
	}
	// TODO: check is there any subdomain registration case?
	// default: // Subdomain
	// 	// Validate Subdomain
	// 	err = k.ValidateRegsiterSubdomain(ctx, domain)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return err
}

// Validate top-level-domain
func (k Keeper) ValidateTopLevelDomain(ctx sdk.Context, domain types.TopLevelDomain) (err error) {
	// Type check
	err = domain.Validate()
	if err != nil {
		return err
	}
	// Check if domain is already taken
	isDomainAlreadyTaken := k.GetIsTopLevelDomainAlreadyTaken(ctx, domain)
	if isDomainAlreadyTaken {
		err = errorsmod.Wrapf(errors.New(fmt.Sprintf("%s", domain.Name)),
			types.ErrDomainIsAlreadyTaken.Error())
		return err
	}

	// Validate TLD
	err = k.ValidateRegisterTLD(ctx, domain)
	if err != nil {
		return err
	}

	return err
}
