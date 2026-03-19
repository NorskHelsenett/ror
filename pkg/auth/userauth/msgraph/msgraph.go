// This package contains the msgraph package which is used to interact with the Microsoft Graph API
// it implements the UserWithGroup struct and the GraphClient struct
package msgraph

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/NorskHelsenett/ror/pkg/auth/authtools"
	"github.com/NorskHelsenett/ror/pkg/config/rorconfig"
	"github.com/NorskHelsenett/ror/pkg/helpers/kvcachehelper"
	"github.com/NorskHelsenett/ror/pkg/helpers/kvcachehelper/memorycache"
	"github.com/NorskHelsenett/ror/pkg/helpers/rorhealth"
	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	graphusers "github.com/microsoftgraph/msgraph-sdk-go/users"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

var ApiEndpoint = "https://graph.microsoft.com/.default"

type MsGraphConfig struct {
	Domain       string `json:"domain"`
	TenantID     string `json:"tenantId"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type MsGraphClient struct {
	Client *msgraphsdk.GraphServiceClient
	Cache  kvcachehelper.CacheInterface
	config MsGraphConfig
}

// NewMsGraphClient creates a new GraphClient
func NewMsGraphClient(config MsGraphConfig, cacheHelper kvcachehelper.CacheInterface) (*MsGraphClient, error) {
	client := &MsGraphClient{config: config, Cache: cacheHelper}
	rlog.Infof("Connecting to msgraph api for domain %s.", config.Domain)
	if cacheHelper != nil {
		client.Cache = cacheHelper
	} else {
		client.Cache = memorycache.NewKvCache()
	}
	connectionStart := time.Now()
	cred, err := azidentity.NewClientSecretCredential(client.config.TenantID, client.config.ClientID, client.config.ClientSecret, nil)
	if err != nil {
		return nil, err
	}

	conn, err := msgraphsdk.NewGraphServiceClientWithCredentials(
		cred, []string{ApiEndpoint},
	)

	if err != nil {
		authtools.ServerConnectionHistogram.WithLabelValues("msgraph", config.Domain, ApiEndpoint, "443", "500").Observe(time.Since(connectionStart).Seconds())
		return nil, err
	}
	authtools.ServerConnectionHistogram.WithLabelValues("msgraph", config.Domain, ApiEndpoint, "443", "200").Observe(time.Since(connectionStart).Seconds())
	rlog.Infof("Connected to msgraph api for domain %s.", config.Domain)
	client.Client = conn
	return client, nil
}

// GetUsersWithGroups gets a user and the name of the groups the user is a member of
// TODO: Implement isExpired
// TODO: Implement isDisabled...
func (g *MsGraphClient) GetUser(ctx context.Context, userId string) (*identitymodels.User, error) {
	ctx, span := otel.GetTracerProvider().Tracer(rorconfig.GetString(rorconfig.TRACER_ID)).Start(ctx, "msgraph.MsGraphClient.GetUser")
	defer span.End()
	if g == nil {
		return nil, fmt.Errorf("msgraph client is nil")
	}
	if g.Client == nil {
		return nil, fmt.Errorf("msgraph graph service client is nil")
	}
	if g.Cache == nil {
		g.Cache = memorycache.NewKvCache()
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var ret *identitymodels.User
	var groupnames []string = []string{}
	var user models.Userable

	groupsChan := make(chan []string, 1)
	userChan := make(chan models.Userable, 1)
	errorChan := make(chan error, 2)
	queryStart := time.Now()
	go g.getUser(ctx, userId, userChan, errorChan)
	go g.getGroups(ctx, userId, groupsChan, errorChan)

	hasUser := false
	hasGroups := false

	for !hasUser || !hasGroups {
		select {
		case returnedgroupids := <-groupsChan:
			returnedgroups, err := g.getGroupDisplayNames(ctx, returnedgroupids)
			if err != nil {
				cancel()
				return nil, err
			}
			groupnames = returnedgroups
			hasGroups = true
		case returneUser := <-userChan:
			user = returneUser
			hasUser = true
		case err := <-errorChan:
			cancel()
			authtools.UserLookupHistogram.WithLabelValues("msgraph", g.config.Domain, ApiEndpoint, "500").Observe(time.Since(queryStart).Seconds())
			return nil, err
		case <-ctx.Done():
			authtools.UserLookupHistogram.WithLabelValues("msgraph", g.config.Domain, ApiEndpoint, "500").Observe(time.Since(queryStart).Seconds())
			return nil, ctx.Err()
		}
	}

	addDomainpartToGroups(&groupnames, userId)
	authtools.UserLookupHistogram.WithLabelValues("msgraph", g.config.Domain, ApiEndpoint, "200").Observe(time.Since(queryStart).Seconds())
	if user == nil {
		return nil, fmt.Errorf("msgraph returned nil user for userId: %s", userId)
	}
	userPrincipalName := user.GetUserPrincipalName()
	if userPrincipalName == nil {
		return nil, fmt.Errorf("msgraph returned nil user principal name for userId: %s", userId)
	}
	displayName := user.GetDisplayName()
	if displayName == nil {
		return nil, fmt.Errorf("msgraph returned nil display name for userId: %s", userId)
	}

	ret = &identitymodels.User{
		Email:           *userPrincipalName,
		Name:            *displayName,
		IsEmailVerified: true,
		Groups:          groupnames,
	}
	rlog.Debug(fmt.Sprintf("Got user %s with %d groups", userId, len(ret.Groups)))
	return ret, nil
}

func addDomainpartToGroups(groupnames *[]string, userId string) {

	_, domain, err := authtools.SplitUserId(userId)
	if err != nil {
		domain = ""
	}

	// TODO: Add check if domainpart is already part of the group name
	for i, group := range *groupnames {
		(*groupnames)[i] = group + "@" + domain
	}
}

// getUser gets a user from the graph api
func (g *MsGraphClient) getUser(ctx context.Context, userId string, userChan chan<- models.Userable, errorChan chan<- error) {
	ctx, span := otel.GetTracerProvider().Tracer(rorconfig.GetString(rorconfig.TRACER_ID)).Start(ctx, "msgraph.MsGraphClient.getUser")
	defer span.End()
	if g == nil {
		err := fmt.Errorf("msgraph client is nil")
		select {
		case errorChan <- err:
		case <-ctx.Done():
		}
		return
	}
	if g.Client == nil {
		err := fmt.Errorf("msgraph graph service client is nil")
		select {
		case errorChan <- err:
		case <-ctx.Done():
		}
		return
	}
	if g.Cache == nil {
		g.Cache = memorycache.NewKvCache()
	}

	cachedValue, cached := g.Cache.Get(ctx, userId, kvcachehelper.CacheGetOptions{Prefix: "user_"})
	if cached {
		cachedUser, ok := cachedValue.(models.Userable)
		if ok && cachedUser != nil {
			span.SetStatus(codes.Ok, "Cache hit")
			select {
			case userChan <- cachedUser:
			case <-ctx.Done():
			}
			return
		}
	}
	fetchedUser, err := g.Client.Users().ByUserId(userId).Get(ctx, nil)
	if err != nil {
		span.RecordError(err)
		select {
		case errorChan <- err:
		case <-ctx.Done():
		}
		return
	}

	if fetchedUser == nil {
		err := fmt.Errorf("msgraph returned nil user for userId: %s", userId)
		span.RecordError(err)
		select {
		case errorChan <- err:
		case <-ctx.Done():
		}
		return
	}

	// Optional: cache write-back here if cache supports this type with prefix
	g.Cache.Set(ctx, userId, fetchedUser, kvcachehelper.CacheSetOptions{Prefix: "user_", Timeout: 15 * time.Minute})

	select {
	case userChan <- fetchedUser:
	case <-ctx.Done():
	}
}

// getGroups gets the groups a user is a member of from the graph api
// It returns a list of group ids
func (g *MsGraphClient) getGroups(ctx context.Context, userId string, groupsChan chan<- []string, errorChan chan<- error) {
	ctx, span := otel.GetTracerProvider().Tracer(rorconfig.GetString(rorconfig.TRACER_ID)).Start(ctx, "msgraph.MsGraphClient.getGroups")
	defer span.End()
	if g == nil {
		err := fmt.Errorf("msgraph client is nil")
		select {
		case errorChan <- err:
		case <-ctx.Done():
		}
		return
	}
	if g.Client == nil {
		err := fmt.Errorf("msgraph graph service client is nil")
		select {
		case errorChan <- err:
		case <-ctx.Done():
		}
		return
	}
	if g.Cache == nil {
		g.Cache = memorycache.NewKvCache()
	}
	// MS Fu%¤d up, change back if they fix their api
	//requestBody := graphusers.NewItemGetmembergroupsGetMemberGroupsPostRequestBody()
	requestBody := graphusers.NewItemGetMemberGroupsPostRequestBody()
	securityEnabledOnly := true
	requestBody.SetSecurityEnabledOnly(&securityEnabledOnly)

	cachedValue, cached := g.Cache.Get(ctx, userId, kvcachehelper.CacheGetOptions{Prefix: "usergroups_"})
	if cached {
		cachedUserGroups, ok := cachedValue.([]string)
		if ok && cachedUserGroups != nil {
			span.SetStatus(codes.Ok, "Cache hit")
			select {
			case groupsChan <- cachedUserGroups:
			case <-ctx.Done():
			}
			return
		}
	}

	fetchedUserGroups, err := g.Client.Users().ByUserId(userId).GetMemberGroups().Post(ctx, requestBody, nil)
	if err != nil {
		span.RecordError(err)
		select {
		case errorChan <- err:
		case <-ctx.Done():
		}
		return
	}

	if fetchedUserGroups == nil {
		err := fmt.Errorf("msgraph returned nil user for userId: %s", userId)
		span.RecordError(err)
		select {
		case errorChan <- err:
		case <-ctx.Done():
		}
		return
	}

	g.Cache.Set(ctx, userId, fetchedUserGroups.GetValue(), kvcachehelper.CacheSetOptions{Prefix: "usergroups_", Timeout: 2 * time.Minute})
	select {
	case groupsChan <- fetchedUserGroups.GetValue():
	case <-ctx.Done():
	}
}

// getGroupDisplayNames gets the display names of a list of groups in parallel
// It returns a list of group names based on the group ids
func (g *MsGraphClient) getGroupDisplayNames(ctx context.Context, groups []string) ([]string, error) {
	ctx, span := otel.GetTracerProvider().Tracer(rorconfig.GetString(rorconfig.TRACER_ID)).Start(ctx, "msgraph.MsGraphClient.getGroupDisplayNames")
	defer span.End()
	if len(groups) == 0 {
		return []string{}, nil
	}
	groupsNameChan := make(chan string, len(groups))
	groupsErrorChan := make(chan error, len(groups))
	for _, value := range groups {
		go g.getGroupDisplayName(ctx, value, groupsNameChan, groupsErrorChan)
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
	span.SetAttributes(
		attribute.Int("numberOfGroups", len(groupNames)),
	)
	span.SetStatus(codes.Ok, "Got group display names")
	return groupNames, nil
}

// getGroupDisplayName gets the display name of a group
// If the group is not in the cache, it will fetch it from the graph api
// and add it to the cache
func (g *MsGraphClient) getGroupDisplayName(ctx context.Context, groupId string, groupsNameChan chan<- string, groupsErrorChan chan<- error) {
	if g == nil {
		err := fmt.Errorf("msgraph client is nil")
		select {
		case groupsErrorChan <- err:
		case <-ctx.Done():
		}
		return
	}
	if g.Client == nil {
		err := fmt.Errorf("msgraph graph service client is nil")
		select {
		case groupsErrorChan <- err:
		case <-ctx.Done():
		}
		return
	}
	if g.Cache == nil {
		g.Cache = memorycache.NewKvCache()
	}

	name, cached := g.Cache.Get(ctx, groupId)
	if cached {
		cachedName, ok := name.(string)
		if ok {
			select {
			case groupsNameChan <- cachedName:
			case <-ctx.Done():
			}
			return
		}
	}

	group, err := g.Client.Groups().ByGroupId(groupId).Get(ctx, nil)
	if err != nil {
		select {
		case groupsErrorChan <- err:
		case <-ctx.Done():
		}
		return
	}
	if group == nil {
		err := fmt.Errorf("msgraph returned nil group for groupId: %s", groupId)
		select {
		case groupsErrorChan <- err:
		case <-ctx.Done():
		}
		return
	}
	groupDisplayName := group.GetDisplayName()
	if groupDisplayName == nil {
		err := fmt.Errorf("msgraph returned nil group display name for groupId: %s", groupId)
		select {
		case groupsErrorChan <- err:
		case <-ctx.Done():
		}
		return
	}
	g.Cache.Set(ctx, groupId, *groupDisplayName)
	select {
	case groupsNameChan <- *groupDisplayName:
	case <-ctx.Done():
	}
}

func (g *MsGraphClient) CheckHealthWithoutContext() []rorhealth.Check {
	var status rorhealth.Status = rorhealth.StatusPass
	if g.Client == nil {
		status = rorhealth.StatusFail
	}
	return []rorhealth.Check{
		{
			ComponentID:   g.config.Domain,
			ComponentType: "msGrapDomainResolver",
			Status:        status,
		},
	}
}

// TODO: Implement
func (g *MsGraphClient) CheckHealth(_ context.Context) []rorhealth.Check {
	return g.CheckHealthWithoutContext()
}
