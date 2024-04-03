// This package contains the msgraph package which is used to interact with the Microsoft Graph API
// it implements the UserWithGroup struct and the GraphClient struct
package msgraph

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/NorskHelsenett/ror/pkg/helpers/kvcachehelper"
	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	graphusers "github.com/microsoftgraph/msgraph-sdk-go/users"
)

type MsGraphConfig struct {
	Domain       string `json:"domain"`
	TenantID     string `json:"tenantId"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

// The cache interface defines the methods that a cache should implement.
type CacheInterface interface {
	Add(key string, value string)
	Get(key string) (string, bool)
	Remove(key string)
}

type MsGraphClient struct {
	Client     *msgraphsdk.GraphServiceClient
	GroupCache kvcachehelper.CacheInterface
	config     MsGraphConfig
}

// NewMsGraphClient creates a new GraphClient
func NewMsGraphClient(config MsGraphConfig, cacheHelper kvcachehelper.CacheInterface) (*MsGraphClient, error) {
	client := &MsGraphClient{config: config, GroupCache: cacheHelper}

	if cacheHelper == nil {
		client.GroupCache = cacheHelper
	}

	cred, err := azidentity.NewClientSecretCredential(client.config.TenantID, client.config.ClientID, client.config.ClientSecret, nil)
	if err != nil {
		return nil, err
	}

	conn, err := msgraphsdk.NewGraphServiceClientWithCredentials(
		cred, []string{"https://graph.microsoft.com/.default"},
	)
	if err != nil {
		return nil, err
	}
	client.Client = conn
	return client, nil
}

// GetUsersWithGroups gets a user and the name of the groups the user is a member of
// TODO: Implement isExpired
// TODO: Implement isDisabled...
func (g *MsGraphClient) GetUser(userId string) (*identitymodels.User, error) {
	var ret *identitymodels.User
	var groupnames []string = []string{}
	var user models.Userable

	groupsChan := make(chan []string)
	userChan := make(chan models.Userable)
	errorChan := make(chan error)

	go g.getUser(userId, userChan, errorChan)
	go g.getGroups(userId, groupsChan, errorChan)

	for i := 0; i < 2; i++ {
		select {
		case returnedgroupids := <-groupsChan:
			returnedgroups, err := g.getGroupDisplayNames(returnedgroupids, g.GroupCache)
			if err != nil {
				return nil, err
			}
			groupnames = returnedgroups
		case returneUser := <-userChan:
			user = returneUser
		case err := <-errorChan:
			return nil, err
		}
	}

	ret = &identitymodels.User{
		Email:           *user.GetUserPrincipalName(),
		Name:            *user.GetDisplayName(),
		IsEmailVerified: true,
		Groups:          groupnames,
	}

	return ret, nil
}

// getUser gets a user from the graph api
func (g *MsGraphClient) getUser(userId string, userChan chan<- models.Userable, errorChan chan<- error) {
	user, err := g.Client.Users().ByUserId(userId).Get(context.Background(), nil)
	if err != nil {
		errorChan <- err
	}
	userChan <- user
}

// getGroups gets the groups a user is a member of from the graph api
// It returns a list of group ids
func (g *MsGraphClient) getGroups(userId string, groupsChan chan<- []string, errorChan chan<- error) {
	requestBody := graphusers.NewItemGetMemberGroupsPostRequestBody()
	securityEnabledOnly := true
	requestBody.SetSecurityEnabledOnly(&securityEnabledOnly)

	groups, err := g.Client.Users().ByUserId(userId).GetMemberGroups().Post(context.Background(), requestBody, nil)
	if err != nil {
		errorChan <- err
	}
	fmt.Print(groups.GetValue())
	groupsChan <- groups.GetValue()
}

// getGroupDisplayNames gets the display names of a list of groups in parallel
// It returns a list of group names based on the group ids
func (g *MsGraphClient) getGroupDisplayNames(groups []string, groupCache CacheInterface) ([]string, error) {
	groupsNameChan := make(chan string, len(groups))
	groupsErrorChan := make(chan error)
	for _, value := range groups {
		go g.getGroupDisplayName(value, groupsNameChan, groupsErrorChan, groupCache)
	}
	groupNames := make([]string, len(groups))

	for i := 0; i < len(groups); i++ {
		select {
		case groupname := <-groupsNameChan:
			groupNames[i] = groupname
		case err := <-groupsErrorChan:
			return nil, err
		}

	}
	return groupNames, nil
}

// getGroupDisplayName gets the display name of a group
// If the group is not in the cache, it will fetch it from the graph api
// and add it to the cache
func (g *MsGraphClient) getGroupDisplayName(groupId string, groupsNameChan chan<- string, groupsErrorChan chan<- error, groupCache CacheInterface) {
	name, cached := groupCache.Get(groupId)
	if cached {
		groupsNameChan <- name
		return
	}
	group, err := g.Client.Groups().ByGroupId(groupId).Get(context.Background(), nil)
	if err != nil {
		groupsErrorChan <- err
		return
	}
	groupCache.Add(groupId, *group.GetDisplayName())
	groupsNameChan <- *group.GetDisplayName()
}
