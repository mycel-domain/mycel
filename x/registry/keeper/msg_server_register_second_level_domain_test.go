package keeper_test

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/mycel-domain/mycel/app/params"
	"github.com/mycel-domain/mycel/testutil"
	"github.com/mycel-domain/mycel/x/registry/types"
)

func (suite *KeeperTestSuite) TestRegisterSecondLevelDomain() {
	testCases := []struct {
		creator                  string
		name                     string
		parent                   string
		registrationPeriodInYear uint64
		domainOwnership          types.DomainOwnership
		expErr                   error
		fn                       func()
	}{
		{
			creator:                  testutil.Alice,
			name:                     "foo",
			parent:                   "cel",
			registrationPeriodInYear: 1,
			domainOwnership: types.DomainOwnership{
				Owner:   testutil.Alice,
				Domains: []*types.OwnedDomain{{Name: "foo", Parent: "cel"}},
			},
			expErr: nil,
			fn:     func() {},
		},
		{
			creator:                  testutil.Alice,
			name:                     "foo",
			parent:                   "cel",
			registrationPeriodInYear: 4,
			domainOwnership: types.DomainOwnership{
				Owner:   testutil.Alice,
				Domains: []*types.OwnedDomain{{Name: "foo", Parent: "cel"}},
			},
			expErr: nil,
			fn:     func() {},
		},
		{
			creator:                  testutil.Alice,
			name:                     "foo",
			parent:                   "cel",
			registrationPeriodInYear: 1,
			expErr:                   errorsmod.Wrapf(types.ErrSecondLevelDomainAlreadyTaken, "foo.cel"),
			fn: func() {
				// Register domain once
				domain := &types.MsgRegisterSecondLevelDomain{
					Creator:                  testutil.Alice,
					Name:                     "foo",
					Parent:                   "cel",
					RegistrationPeriodInYear: 1,
				}
				_, err := suite.msgServer.RegisterSecondLevelDomain(suite.ctx, domain)
				suite.Require().Nil(err)
			},
		},
		{
			creator:                  testutil.Alice,
			name:                     "foo",
			parent:                   "xxx",
			registrationPeriodInYear: 1,
			expErr:                   errorsmod.Wrapf(types.ErrSecondLevelDomainParentDoesNotExist, "xxx"),
			fn: func() {
			},
		},
		{
			creator:                  testutil.Bob,
			name:                     "foo",
			parent:                   "alice",
			registrationPeriodInYear: 1,
			domainOwnership: types.DomainOwnership{
				Owner:   testutil.Alice,
				Domains: []*types.OwnedDomain{{Name: "foo", Parent: "alice"}},
			},
			expErr: errorsmod.Wrapf(types.ErrNotAllowedRegisterDomain, "alice"),
			fn: func() {
				owner := testutil.Alice
				// Create .bob TLD by Alice
				_, _, err := suite.app.RegistryKeeper.RegisterTopLevelDomain(suite.ctx, owner, "alice", 1)
				suite.Require().Nil(err)
			},
		},
	}

	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			suite.SetupTest()

			registerMsg := &types.MsgRegisterSecondLevelDomain{
				Creator:                  tc.creator,
				Name:                     tc.name,
				Parent:                   tc.parent,
				RegistrationPeriodInYear: tc.registrationPeriodInYear,
			}

			// Run test case function
			tc.fn()

			if tc.expErr == nil {
				beforeParent, found := suite.app.RegistryKeeper.GetTopLevelDomain(suite.ctx, tc.parent)
				suite.Require().True(found)

				moduleAddress := authtypes.NewModuleAddress(types.ModuleName)
				beforeModuleBalance := suite.app.BankKeeper.GetBalance(suite.ctx, moduleAddress, params.DefaultBondDenom)

				// Register second level domain
				_, err := suite.msgServer.RegisterSecondLevelDomain(suite.ctx, registerMsg)
				suite.Require().Nil(err)

				// Evaluate domain ownership
				domainOwnership, found := suite.app.RegistryKeeper.GetDomainOwnership(suite.ctx, tc.creator)
				suite.Require().True(found)
				suite.Require().Equal(tc.domainOwnership, domainOwnership)

				// Evaluate if domain is registered
				domain, found := suite.app.RegistryKeeper.GetSecondLevelDomain(suite.ctx, tc.name, tc.parent)
				suite.Require().True(found)
				suite.Require().Equal(domain.GetRole(tc.creator), types.DomainRole_OWNER)

				// Evaluate if domain is appended to owned domain list
				ownedDomains, found := suite.app.RegistryKeeper.GetDomainOwnership(suite.ctx, tc.creator)
				suite.Require().True(found)
				found = false
				for _, ownedDomain := range ownedDomains.Domains {
					if ownedDomain.Name == tc.name && ownedDomain.Parent == tc.parent {
						found = true
						break
					}
				}
				suite.Require().True(found)

				// Evaluate if parent's subdomainCount is increased
				afterParent, found := suite.app.RegistryKeeper.GetTopLevelDomain(suite.ctx, tc.parent)
				suite.Require().True(found)
				suite.Require().Equal(beforeParent.SubdomainCount+1, afterParent.SubdomainCount)

				// Evaluate if module account balance is increased
				// Get registration fee
				config := suite.app.RegistryKeeper.GetSecondLevelDomainParentsSubdomainConfig(suite.ctx, domain)
				fee, err := config.GetRegistrationFee(tc.name, tc.registrationPeriodInYear)
				suite.Require().Nil(err)

				afterModuleBalance := suite.app.BankKeeper.GetBalance(suite.ctx, moduleAddress, params.DefaultBondDenom)
				suite.Require().Equal(beforeModuleBalance.Add(fee), afterModuleBalance)
				suite.Require().Equal(beforeParent.TotalWithdrawalAmount.Add(fee), afterParent.TotalWithdrawalAmount)

				// Evalute events
				events, found := testutil.FindEventsByType(suite.ctx.EventManager().Events(), types.EventTypeRegisterSecondLevelDomain)
				suite.Require().True(found)
				for _, event := range events {
					suite.Require().Equal(tc.name, event.Attributes[0].Value)
					suite.Require().Equal(tc.parent, event.Attributes[1].Value)
					suite.Require().Equal(fee.String(), event.Attributes[3].Value)
				}
			} else {
				// Register second level domain
				_, err := suite.msgServer.RegisterSecondLevelDomain(suite.ctx, registerMsg)
				suite.Require().EqualError(err, tc.expErr.Error())
			}
		})
	}
}
