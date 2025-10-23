package userauth

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/NorskHelsenett/ror/pkg/auth/authtools"
	"github.com/NorskHelsenett/ror/pkg/auth/userauth/activedirectory"
	"github.com/NorskHelsenett/ror/pkg/auth/userauth/ldaps"
	"github.com/NorskHelsenett/ror/pkg/auth/userauth/msgraph"
	"github.com/NorskHelsenett/ror/pkg/clients"
	"github.com/NorskHelsenett/ror/pkg/helpers/rorhealth"
	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"
)

type DomainResolverConfigs struct {
	DomainResolvers []DomainResolverConfig `json:"domainResolvers"`
}

// TODO: add enum for resolverType
type DomainResolverConfig struct {
	ResolverType  string                    `json:"resolverType"`
	AdConfig      *activedirectory.AdConfig `json:"adConfig,omitempty"`
	LdapConfig    *ldaps.LdapConfig         `json:"ldapConfig,omitempty"`
	MsGraphConfig *msgraph.MsGraphConfig    `json:"msGraphConfig,omitempty"`
}

type DomainResolverInterface interface {
	GetUser(ctx context.Context, userId string) (*identitymodels.User, error)
	clients.CommonHealthChecker
}

type DomainResolvers struct {
	resolvers map[string]DomainResolverInterface
}

func (d DomainResolvers) GetUser(ctx context.Context, userId string) (*identitymodels.User, error) {
	_, domain, err := authtools.SplitUserId(userId)
	if err != nil {
		return nil, err
	}

	if len(d.resolvers) == 0 {
		return nil, fmt.Errorf("no domainresolvers configured")
	}

	if domainResolver, ok := d.resolvers[domain]; ok {
		return domainResolver.GetUser(ctx, userId)
	}
	return nil, fmt.Errorf("no domain resolver found for domain: %s", domain)
}
func (d *DomainResolvers) SetDomain(domain string, resolver DomainResolverInterface) {
	if d.resolvers == nil {
		d.resolvers = map[string]DomainResolverInterface{}
	}
	d.resolvers[domain] = resolver
}

func (d *DomainResolvers) RemoveDomain(domain string) {
	delete(d.resolvers, domain)
}

func (d DomainResolvers) RegisterHealthChecks() {
	if len(d.resolvers) != 0 && d.resolvers != nil {
		for key, resolver := range d.resolvers {
			checkname := fmt.Sprintf("domainresolvers-%s", key)
			rorhealth.Register(context.TODO(), checkname, resolver)
		}
	}
}

func NewDomainResolversFromJson(jsonBytes []byte) (*DomainResolvers, error) {
	var domainResolverConfigs DomainResolverConfigs
	var domainResolvers *DomainResolvers = &DomainResolvers{}
	err := json.Unmarshal(jsonBytes, &domainResolverConfigs)
	if err != nil {
		return nil, err
	}

	var errs []error

	if len(domainResolverConfigs.DomainResolvers) == 0 {
		return nil, fmt.Errorf("no domain resolvers found")
	}

	for _, domainResolverConfig := range domainResolverConfigs.DomainResolvers {
		if domainResolverConfig.ResolverType == "ldaps" {
			domainResolverConfig.ResolverType = "ldap"
		}

		switch domainResolverConfig.ResolverType {
		case "ldap":
			ldapsClient, err := ldaps.NewLdapsClient(*domainResolverConfig.LdapConfig)
			if err != nil {
				errs = append(errs, err)
			}
			domainResolvers.SetDomain(domainResolverConfig.LdapConfig.Domain, ldapsClient)

		case "ad":
			adClient, err := activedirectory.NewAdClient(*domainResolverConfig.AdConfig)
			if err != nil {
				errs = append(errs, err)
			}
			domainResolvers.SetDomain(domainResolverConfig.AdConfig.Domain, adClient)
		case "msgraph":
			msGraphClient, err := msgraph.NewMsGraphClient(*domainResolverConfig.MsGraphConfig, nil)
			if err != nil {
				errs = append(errs, err)
			}
			domainResolvers.SetDomain(domainResolverConfig.MsGraphConfig.Domain, msGraphClient)
		default:
			err = fmt.Errorf("unknown resolver type: %s", domainResolverConfig.ResolverType)
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		reterr := "the following error were encountered trying to load the resolverconfigs:"
		for i, err := range errs {
			reterr = reterr + " " + strconv.Itoa(int(i+1)) + ": " + err.Error()
		}
		err = fmt.Errorf("%s", reterr)
		return domainResolvers, err
	}

	return domainResolvers, nil
}
