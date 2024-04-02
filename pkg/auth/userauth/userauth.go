package userauth

import (
	"fmt"
	"strings"

	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"
)

type DomainResolverInterface interface {
	GetUser(userId string) (*identitymodels.User, error)
}

type Domains map[string]DomainResolverInterface

func (d Domains) GetUser(userId string) (*identitymodels.User, error) {
	domain, _ := splitUserId(userId)
	if domainResolver, ok := d[domain]; ok {
		return domainResolver.GetUser(userId)
	}
	return nil, fmt.Errorf("no domain resolver found for domain: %s", domain)
}

func splitUserId(userId string) (string, string) {
	parts := strings.Split(userId, "@")
	if len(parts) != 2 {
		return "", ""
	}
	return parts[1], parts[0]
}
