package userauth

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	"github.com/NorskHelsenett/ror/pkg/auth/authtools"
	"github.com/NorskHelsenett/ror/pkg/auth/userauth/activedirectory"
	"github.com/NorskHelsenett/ror/pkg/auth/userauth/ldaps"
	"github.com/NorskHelsenett/ror/pkg/auth/userauth/msgraph"
	"github.com/NorskHelsenett/ror/pkg/clients"
	"github.com/NorskHelsenett/ror/pkg/helpers/rorhealth"
	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"
	"github.com/NorskHelsenett/ror/pkg/telemetry/rortracer"
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
	mu        sync.RWMutex
	resolvers map[string]DomainResolverInterface
}

// NewDomainResolvers returns an empty, ready-to-use resolver registry. Resolvers
// can be added later with SetDomain or AddResolverFromConfig, which makes it safe
// to construct the registry up front and populate it asynchronously while the rest
// of the application starts.
func NewDomainResolvers() *DomainResolvers {
	return &DomainResolvers{resolvers: map[string]DomainResolverInterface{}}
}

func (d *DomainResolvers) GetUser(ctx context.Context, userId string) (*identitymodels.User, error) {
	ctx, span := rortracer.StartSpan(ctx, "userauth.DomainResolvers.GetUser")
	defer span.End()
	_, domain, err := authtools.SplitUserId(userId)
	if err != nil {
		return nil, err
	}

	d.mu.RLock()
	count := len(d.resolvers)
	domainResolver, ok := d.resolvers[domain]
	d.mu.RUnlock()

	if count == 0 {
		return nil, fmt.Errorf("no domainresolvers configured")
	}

	if ok {
		return domainResolver.GetUser(ctx, userId)
	}
	return nil, fmt.Errorf("no domain resolver found for domain: %s", domain)
}
func (d *DomainResolvers) SetDomain(domain string, resolver DomainResolverInterface) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.resolvers == nil {
		d.resolvers = map[string]DomainResolverInterface{}
	}
	d.resolvers[domain] = resolver
}

func (d *DomainResolvers) RemoveDomain(domain string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.resolvers, domain)
}

func (d *DomainResolvers) RegisterHealthChecks() {
	d.mu.RLock()
	defer d.mu.RUnlock()
	for key, resolver := range d.resolvers {
		checkname := fmt.Sprintf("domainresolvers-%s", key)
		rorhealth.Register(context.TODO(), checkname, resolver)
	}
}

// ParseDomainResolverConfigs unmarshals the resolver configuration json into a
// slice of configs without building any clients. This lets callers decide how to
// instantiate the resolvers (e.g. synchronously or asynchronously).
func ParseDomainResolverConfigs(jsonBytes []byte) ([]DomainResolverConfig, error) {
	var domainResolverConfigs DomainResolverConfigs
	if err := json.Unmarshal(jsonBytes, &domainResolverConfigs); err != nil {
		return nil, err
	}
	if len(domainResolverConfigs.DomainResolvers) == 0 {
		return nil, fmt.Errorf("no domain resolvers found")
	}
	return domainResolverConfigs.DomainResolvers, nil
}

// AddResolverFromConfig builds a single domain resolver from its config, adds it
// to the registry and registers its health check. The network connection is
// established here, so this call can block; it is safe to call concurrently and at
// runtime, allowing additional resolvers to be added after startup.
func (d *DomainResolvers) AddResolverFromConfig(cfg DomainResolverConfig) error {
	if cfg.ResolverType == "ldaps" {
		cfg.ResolverType = "ldap"
	}

	var domain string
	var resolver DomainResolverInterface

	switch cfg.ResolverType {
	case "ldap":
		if cfg.LdapConfig == nil {
			return fmt.Errorf("ldap resolver config is missing")
		}
		ldapsClient, err := ldaps.NewLdapsClient(*cfg.LdapConfig)
		if err != nil {
			return err
		}
		domain, resolver = cfg.LdapConfig.Domain, ldapsClient
	case "ad":
		if cfg.AdConfig == nil {
			return fmt.Errorf("ad resolver config is missing")
		}
		adClient, err := activedirectory.NewAdClient(*cfg.AdConfig)
		if err != nil {
			return err
		}
		domain, resolver = cfg.AdConfig.Domain, adClient
	case "msgraph":
		if cfg.MsGraphConfig == nil {
			return fmt.Errorf("msgraph resolver config is missing")
		}
		msGraphClient, err := msgraph.NewMsGraphClient(*cfg.MsGraphConfig, nil)
		if err != nil {
			return err
		}
		domain, resolver = cfg.MsGraphConfig.Domain, msGraphClient
	default:
		return fmt.Errorf("unknown resolver type: %s", cfg.ResolverType)
	}

	d.SetDomain(domain, resolver)
	rorhealth.Register(context.TODO(), fmt.Sprintf("domainresolvers-%s", domain), resolver)
	return nil
}

func NewDomainResolversFromJson(jsonBytes []byte) (*DomainResolvers, error) {
	configs, err := ParseDomainResolverConfigs(jsonBytes)
	if err != nil {
		return nil, err
	}

	domainResolvers := NewDomainResolvers()

	var errs []error
	for _, cfg := range configs {
		if err := domainResolvers.AddResolverFromConfig(cfg); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		reterr := "the following error were encountered trying to load the resolverconfigs:"
		for i, err := range errs {
			reterr = reterr + " " + strconv.Itoa(int(i+1)) + ": " + err.Error()
		}
		return domainResolvers, fmt.Errorf("%s", reterr)
	}

	return domainResolvers, nil
}
