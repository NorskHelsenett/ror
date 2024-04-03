package userauth

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/auth/userauth/ldaps"
	"github.com/NorskHelsenett/ror/pkg/auth/userauth/msgraph"
	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"
)

type DomainResolverConfigs struct {
	DomainResolvers []DomainResolverConfig `json:"domainResolvers"`
}
type DomainResolverConfig struct {
	ResolverType  string                 `json:"resolverType"`
	LdapsConfig   *ldaps.LdapConfig      `json:"ldapsConfig,omitempty"`
	MsGraphConfig *msgraph.MsGraphConfig `json:"msGraphConfig,omitempty"`
}

type DomainResolverInterface interface {
	GetUser(userId string) (*identitymodels.User, error)
}

type DomainResolvers map[string]DomainResolverInterface

func (d DomainResolvers) GetUser(userId string) (*identitymodels.User, error) {
	domain, _, err := splitUserId(userId)
	if err != nil {
		return nil, err
	}
	if domainResolver, ok := d[domain]; ok {
		return domainResolver.GetUser(userId)
	}
	return nil, fmt.Errorf("no domain resolver found for domain: %s", domain)
}
func (d DomainResolvers) SetDomain(domain string, resolver DomainResolverInterface) {
	d[domain] = resolver
}

func (d DomainResolvers) RemoveDomain(domain string) {
	delete(d, domain)
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
	var domainResolvers *DomainResolvers
	err := json.Unmarshal(jsonBytes, &domainResolverConfigs)
	if err != nil {
		return nil, err
	}

	if len(domainResolverConfigs.DomainResolvers) == 0 {
		return nil, fmt.Errorf("no domain resolvers found")
	}

	for _, domainResolverConfig := range domainResolverConfigs.DomainResolvers {
		if domainResolverConfig.ResolverType == "ldaps" {
			domainResolverConfig.ResolverType = "ldap"
		}

		switch domainResolverConfig.ResolverType {
		case "ldap":
			ldapsClient, err := ldaps.NewLdapsClient(*domainResolverConfig.LdapsConfig)
			if err != nil {
				return nil, err
			}
			domainResolvers.SetDomain(domainResolverConfig.LdapsConfig.Domain, ldapsClient)
		case "msgraph":
			msGraphClient, err := msgraph.NewMsGraphClient(*domainResolverConfig.MsGraphConfig, nil)
			if err != nil {
				return nil, err
			}
			domainResolvers.SetDomain(domainResolverConfig.MsGraphConfig.Domain, msGraphClient)
		default:
			return nil, fmt.Errorf("unknown resolver type: %s", domainResolverConfig.ResolverType)
		}
	}
	return domainResolvers, nil
}
