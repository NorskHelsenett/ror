package userauth

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/auth/userauth/activedirectory"
	"github.com/NorskHelsenett/ror/pkg/auth/userauth/ldaps"
	"github.com/NorskHelsenett/ror/pkg/auth/userauth/msgraph"
	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"
	newhealth "github.com/dotse/go-health"
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
	GetUser(userId string) (*identitymodels.User, error)
	CheckHealth() []newhealth.Check
}

type DomainResolvers struct {
	resolvers map[string]DomainResolverInterface
}

func (d DomainResolvers) GetUser(userId string) (*identitymodels.User, error) {
	domain, _, err := splitUserId(userId)
	if err != nil {
		return nil, err
	}

	if len(d.resolvers) == 0 {
		return nil, fmt.Errorf("no domainresolvers configured")
	}

	if domainResolver, ok := d.resolvers[domain]; ok {
		return domainResolver.GetUser(userId)
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
			newhealth.Register(checkname, resolver)
		}
	}
}

func splitUserId(userId string) (string, string, error) {
	parts := strings.Split(userId, "@")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid userId: %s", userId)
	}
	return parts[1], parts[0], nil
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
