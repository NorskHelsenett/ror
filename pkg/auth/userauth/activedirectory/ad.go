package activedirectory

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"regexp"
	"strconv"
	"time"

	"github.com/NorskHelsenett/ror/pkg/auth/authtools"
	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	newhealth "github.com/dotse/go-health"
	"github.com/go-ldap/ldap"
)

type AdConfig struct {
	Domain       string     `json:"domain"`
	BindUser     string     `json:"bindUser"`
	BindPassword string     `json:"bindPassword"`
	BaseDN       string     `json:"basedn"`
	Servers      []AdServer `json:"servers,omitempty"`
	Certificate  []byte     `json:"certificate,omitempty"` // This is the CA certificate
}

type AdServer struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type AdClient struct {
	connection *ldap.Conn
	config     *AdConfig
}

func NewAdClient(config AdConfig) (*AdClient, error) {
	ldapsClient := &AdClient{config: &config}

	err := ldapsClient.Connect()
	if err != nil {
		return nil, err
	}
	return ldapsClient, nil
}
func (l *AdClient) Connect() error {
	var client *ldap.Conn
	ldapSecurePort, err := strconv.Atoi(ldap.DefaultLdapsPort)
	if err != nil {
		return fmt.Errorf("failed to parse default ldaps port")
	}

	ldapPort, err := strconv.Atoi(ldap.DefaultLdapPort)
	if err != nil {
		return fmt.Errorf("failed to parse default ldaps port")
	}

	if len(l.config.Servers) == 0 {
		rlog.Infof("Using dns-discovery to find ldap servers for domain %s", l.config.Domain)
		_, srvs, err := net.LookupSRV("ldap", "tcp", l.config.Domain)

		for _, srv := range srvs {
			var port int
			if l.config.Certificate != nil {
				port = ldapSecurePort
			} else {
				port = ldapPort
			}
			l.config.Servers = append(l.config.Servers, AdServer{
				Host: srv.Target,
				Port: port,
			})
		}

		if err != nil {
			rlog.Error("error getting servers from dns", err, rlog.Any("Domain", l.config.Domain))
			return err
		}
	}

	//shuffle servers to spread the love
	rand.Shuffle(len(l.config.Servers), func(i, j int) { l.config.Servers[i], l.config.Servers[j] = l.config.Servers[j], l.config.Servers[i] })

	for _, ldapserver := range l.config.Servers {
		rlog.Infof("Trying server %s for domain %s.", ldapserver.Host, l.config.Domain)
		if l.config.Certificate != nil {
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
			continue
		}

		err = client.Bind(l.config.BindUser, l.config.BindPassword)
		if err != nil {
			rlog.Error("an error occurred authenticating to LDAP-host.", err, rlog.Any("Host", ldapserver.Host), rlog.Any("Port", ldapserver.Port), rlog.Any("BindUser", l.config.BindUser))
		} else {
			rlog.Infof("Connected to server server %s for domain %s.", ldapserver.Host, l.config.Domain)
			break
		}
	}

	if client == nil {
		return fmt.Errorf("could not connect to any LDAP server")
	}
	l.connection = client
	return nil
}

func (l *AdClient) search(basedn, filter string, attributes []string) (*ldap.SearchResult, error) {
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

func (l *AdClient) GetUser(userId string) (*identitymodels.User, error) {

	userpart, domainpart, err := authtools.SplitUserId(userId)
	if err != nil {
		return nil, err
	}
	filter := fmt.Sprintf("(&(objectClass=user)(sAMAccountName=%s))", userpart)
	attributes := []string{"cn", "memberOf", "userAccountControl", "accountExpires"}

	if l.connection.IsClosing() {
		rlog.Debug("Reconnecting to Active Directory")
		err := l.Connect()
		if err != nil {
			return nil, err
		}
	}

	result, err := l.search(l.config.BaseDN, filter, attributes)

	if err != nil {
		return nil, err
	}
	var userEntry *ldap.Entry
	if result != nil && len(result.Entries) == 1 {
		for _, entry := range result.Entries {
			userEntry = entry
		}
	} else {
		return nil, fmt.Errorf("could not find user")
	}

	userAccountControl := userEntry.GetAttributeValue("userAccountControl")
	err = checkUserAccountControl(userAccountControl)
	if err != nil {
		return nil, err
	}

	accountExpires := userEntry.GetAttributeValue("accountExpires")
	err = checkUserExpiration(accountExpires)
	if err != nil {
		return nil, err
	}

	userGroups := make([]string, 0)
	memberOfString := result.Entries[0].GetAttributeValues("memberOf")
	if len(userEntry.GetAttributeValue("memberOf")) > 0 {
		re := regexp.MustCompile("CN=([^,]+)")
		for _, entry := range memberOfString {
			match := re.FindStringSubmatch(entry)
			userGroups = append(userGroups, fmt.Sprintf("%s@%s", match[1], domainpart))
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

func checkUserAccountControl(userAccountControl string) error {
	intuac, err := strconv.ParseUint(userAccountControl, 10, 64)
	if err != nil {
		return errors.Join(errors.New("error parsing userAccountControl"), err)
	}
	if !checkBitwiseUnion(512, intuac) {
		return errors.New("user dont have the flag 'NORMAL_ACCOUNT' set")
	}
	if checkBitwiseUnion(2, intuac) {
		return errors.New("user has th flag 'ACCOUNTDISABLE' set")
	}
	if checkBitwiseUnion(16, intuac) {
		return errors.New("user has the flag 'LOCKOUT' set")
	}
	if checkBitwiseUnion(32, intuac) {
		return errors.New("user has the flag 'PASSWD_NOTREQD' set")
	}
	if checkBitwiseUnion(8388608, intuac) {
		return errors.New("user has the flag 'PASSWORD_EXPIRED' set")
	}
	return nil
}

func checkUserExpiration(accountExpiresString string) error {
	accountExpires, err := strconv.ParseInt(accountExpiresString, 10, 64)
	if err != nil {
		return errors.Join(errors.New("error parsing account expired"), err)
	}
	if isAccountExpired(accountExpires) {
		return errors.New("account is expired")
	}
	return nil
}

func checkBitwiseUnion(a uint64, b uint64) bool {
	return (a & b) != 0
}

// Helper functions for readability
func isAccountExpired(accountExpires int64) bool {
	// These values are implicit never-expires
	if accountExpires == 9223372036854775807 || accountExpires == 0 {
		return false
	}

	expiresTime := convertFileTime(accountExpires)
	currentTime := time.Now()
	return !expiresTime.After(currentTime)
}

func convertFileTime(fileTime int64) time.Time {
	unixTime := (fileTime - 116444736000000000) / 10000000
	return time.Unix(unixTime, 0)
}

// TODO: Implement
func (l *AdClient) CheckHealth() []newhealth.Check {
	return []newhealth.Check{
		{
			ComponentID:   l.config.Domain,
			ComponentType: "activeDirectoryDomainResolver",
			Status:        newhealth.StatusPass,
		},
	}
}
