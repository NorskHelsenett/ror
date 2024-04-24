package ldaps

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/NorskHelsenett/ror/pkg/auth/authtools"
	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	newhealth "github.com/dotse/go-health"
	"github.com/go-ldap/ldap"
)

var DefaultTimeout = 10 * time.Second

type LdapConfig struct {
	Domain       string       `json:"domain"`
	BindUser     string       `json:"bindUser"`
	BindPassword string       `json:"bindPassword"`
	BaseDN       string       `json:"basedn"`
	Servers      []LdapServer `json:"servers"`
	Certificate  []byte       `json:"certificate,omitempty"` // This is the CA certificate
}

type LdapServer struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type LdapsClient struct {
	connection *ldap.Conn
	config     *LdapConfig
}

func NewLdapsClient(config LdapConfig) (*LdapsClient, error) {
	ldapsClient := &LdapsClient{config: &config}

	err := ldapsClient.Connect()
	if err != nil {
		return nil, err
	}
	return ldapsClient, nil
}
func (l *LdapsClient) Connect() error {
	var client *ldap.Conn

	ldap.DefaultTimeout = DefaultTimeout

	for _, ldapserver := range l.config.Servers {
		ldapsport, err := strconv.Atoi(ldap.DefaultLdapsPort)
		rlog.Infof("Trying server %s for domain %s.", ldapserver.Host, l.config.Domain)
		if err != nil {
			return fmt.Errorf("failed to parse default ldaps port")
		}
		connectionStart := time.Now()
		if ldapserver.Port == ldapsport {
			caCert := l.config.Certificate
			caCertPool := x509.NewCertPool()
			ok := caCertPool.AppendCertsFromPEM(caCert)
			if !ok {
				return fmt.Errorf("failed to parse root certificate")
			}
			tlsConf := &tls.Config{
				RootCAs: caCertPool,
			}

			client, err = ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", ldapserver.Host, ldapserver.Port), tlsConf)

		} else {
			client, err = ldap.DialURL(fmt.Sprintf("ldap://%s:%d", ldapserver.Host, ldapserver.Port))
		}

		if err != nil {
			rlog.Error("an error occurred connecting to LDAP-host.", err, rlog.Any("Host", ldapserver.Host), rlog.Any("Port", ldapserver.Port))
			authtools.ServerConnectionHistogram.WithLabelValues("openldap", l.config.Domain, ldapserver.Host, strconv.Itoa(ldapserver.Port), "500").Observe(time.Since(connectionStart).Seconds())
		}
		err = client.Bind(l.config.BindUser, l.config.BindPassword)
		if err != nil {
			rlog.Error("an error occurred authenticating to LDAP-host.", err, rlog.Any("Host", ldapserver.Host), rlog.Any("Port", ldapserver.Port), rlog.Any("BindUser", l.config.BindUser))
			authtools.ServerConnectionHistogram.WithLabelValues("openldap", l.config.Domain, ldapserver.Host, strconv.Itoa(ldapserver.Port), "401").Observe(time.Since(connectionStart).Seconds())
		} else {
			rlog.Infof("Connected to server server %s for domain %s.", ldapserver.Host, l.config.Domain)
			authtools.ServerConnectionHistogram.WithLabelValues("openldap", l.config.Domain, ldapserver.Host, strconv.Itoa(ldapserver.Port), "200").Observe(time.Since(connectionStart).Seconds())
			break
		}

	}

	if client == nil {
		return fmt.Errorf("could not connect to any LDAP server")
	}
	l.connection = client
	return nil
}

func (l *LdapsClient) search(basedn, filter string, attributes []string) (*ldap.SearchResult, error) {
	request := ldap.NewSearchRequest(
		basedn,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		filter,
		attributes,
		nil,
	)
	result, err := l.connection.Search(request)
	if err != nil {
		return nil, fmt.Errorf("search error: %s", err)
	}

	if len(result.Entries) > 0 {
		return result, nil
	}

	return nil, fmt.Errorf("could not fetch search entries")
}

func (l *LdapsClient) GetUser(ctx context.Context, userId string) (*identitymodels.User, error) {

	_, domainpart, err := authtools.SplitUserId(userId)
	if err != nil {
		return nil, err
	}

	filter := fmt.Sprintf("(&(objectClass=organizationalPerson)(mail=%s))", userId)

	//result, err := l.aearch(client, ldapConfig.BaseDN, filter, []string{"DN", "cn", "mail"})
	attributes := []string{"DN", "cn", "mail"}

	if l.connection.IsClosing() {
		rlog.Debug("Reconnecting to LDAP")
		authtools.ServerReconnectCounter.WithLabelValues("openldap", l.config.Domain).Inc()
		err := l.Connect()
		if err != nil {
			return nil, err
		}
	}

	queryStart := time.Now()
	result, err := l.search(l.config.BaseDN, filter, attributes)

	if err != nil {
		authtools.UserLookupHistogram.WithLabelValues("openldap", l.config.Domain, "500").Observe(time.Since(queryStart).Seconds())

		return nil, err
	}
	authtools.UserLookupHistogram.WithLabelValues("openldap", l.config.Domain, "200").Observe(time.Since(queryStart).Seconds())

	var userEntry *ldap.Entry
	if result != nil && len(result.Entries) == 1 {
		for _, entry := range result.Entries {
			userEntry = entry
		}
	} else {
		return nil, fmt.Errorf("could not find user")
	}

	groupFilter := fmt.Sprintf("(&(objectClass=groupOfNames)(member=%s))", userEntry.DN)
	groups, err := l.search(l.config.BaseDN, groupFilter, []string{"cn", "member"})
	if err != nil {
		return nil, err
	}
	userGroups := make([]string, 0)
	if groups != nil && len(groups.Entries) > 0 {
		for _, entry := range groups.Entries {
			userGroups = append(userGroups, fmt.Sprintf("%s@%s", entry.GetAttributeValue("cn"), domainpart))
		}
	} else {
		return nil, errors.New("account has no groups")
	}

	user := identitymodels.User{
		Email:           userId,
		Name:            userEntry.GetAttributeValue("cn"),
		IsEmailVerified: true,
		Groups:          userGroups,
	}
	rlog.Debug(fmt.Sprintf("Got user %s with %d groups", userId, len(user.Groups)))
	return &user, nil
}

// TODO: Implement
func (l *LdapsClient) CheckHealth() []newhealth.Check {
	return []newhealth.Check{
		{
			ComponentID:   l.config.Domain,
			ComponentType: "openLdapDomainResolver",
			Status:        newhealth.StatusPass,
		},
	}
}
