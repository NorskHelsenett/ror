package apikeysservice

import (
	"context"
	"errors"
	"fmt"
	clustersservice "github.com/NorskHelsenett/ror/cmd/api/services/clustersService"
	"time"

	"github.com/NorskHelsenett/ror/internal/auditlog"
	"github.com/NorskHelsenett/ror/internal/models"
	apikeyrepo "github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/apikeysRepo"
	datacenterRepo "github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/datacentersRepo"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/context/rorcontext"

	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"
	"github.com/NorskHelsenett/ror/pkg/models/providers"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/apicontracts/v2/apicontractsv2self"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/NorskHelsenett/ror/pkg/helpers/stringhelper"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

// TODO: Move and remove duplicate in repo
func mustGetApikeySalt() string {
	apisalt := viper.GetString(configconsts.API_KEY_SALT)
	if len(apisalt) == 0 {
		panic("api key salt is missing")
	}
	return apisalt
}

func VerifyApiKey(ctx context.Context, apikey string) (apicontracts.ApiKey, error) {
	apikeyhashed := stringhelper.HashSHA512(apikey, []byte(mustGetApikeySalt()))

	apikeys, err := apikeyrepo.GetByHash(ctx, apikeyhashed)
	if err != nil {
		return apicontracts.ApiKey{}, fmt.Errorf("error when getting apikeys by hash from repo")
	}

	if len(apikeys) == 0 {
		return apicontracts.ApiKey{}, fmt.Errorf("error no api key matched provided key")
	}

	if len(apikeys) > 1 {
		return apicontracts.ApiKey{}, fmt.Errorf("error duplicate hashes matching the provided hash")
	}

	if apikeys[0].IsExpired() {
		return apicontracts.ApiKey{}, fmt.Errorf("error apikey expired")
	}

	return apikeys[0], nil
}

func GetByFilter(ctx context.Context, filter *apicontracts.Filter) (*apicontracts.PaginatedResult[apicontracts.ApiKey], error) {
	apikeys, totalCount, err := apikeyrepo.GetByFilter(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error when getting apikeys by filter from repo: %v", err)
	}

	paginatedResult := apicontracts.PaginatedResult[apicontracts.ApiKey]{}

	paginatedResult.Data = apikeys
	paginatedResult.DataCount = int64(len(apikeys))
	paginatedResult.Offset = int64(filter.Skip)
	paginatedResult.TotalCount = int64(totalCount)

	return &paginatedResult, nil
}

func Delete(ctx context.Context, apikeyId string, identity *identitymodels.Identity) (bool, error) {
	if !identity.IsUser() {
		return false, fmt.Errorf("could not delete object, must be delete by a user")
	}

	deleted, deletedObject, err := apikeyrepo.Delete(ctx, apikeyId)
	if err != nil {
		return false, fmt.Errorf("could not delete object: %v", err)
	}

	_, err = auditlog.Create(ctx, "Apikey deleted", models.AuditCategoryApikey, models.AuditActionDelete, identity.User, nil, deletedObject)
	if err != nil {
		return false, fmt.Errorf("could not audit log delete action: %v", err)
	}

	return deleted, nil
}

func DeleteForUser(ctx context.Context, apikeyId string, identity *identitymodels.Identity) (bool, error) {
	if !identity.IsUser() {
		return false, fmt.Errorf("could not delete object, must be delete by a user")
	}

	apikeys, _, err := apikeyrepo.GetByFilter(ctx, &apicontracts.Filter{
		Filters: []apicontracts.FilterMetadata{
			{
				Field:     "identifier",
				Value:     identity.GetId(),
				MatchMode: apicontracts.MatchModeEquals,
			},
		},
	})
	if err != nil {
		return false, fmt.Errorf("could not check if api key exist for apikeyid: %s", apikeyId)
	}

	var apikey apicontracts.ApiKey
	for _, value := range apikeys {
		if value.Id == apikeyId {
			apikey = value
			break
		}
	}

	if len(apikey.Id) == 0 {
		return false, fmt.Errorf("could not delete api key for apikeyid: %s", apikeyId)
	}

	deleted, deletedObject, err := apikeyrepo.Delete(ctx, apikeyId)
	if err != nil {
		return false, fmt.Errorf("could not delete object: %v", err)
	}

	_, err = auditlog.Create(ctx, "Apikey deleted", models.AuditCategoryApikey, models.AuditActionDelete, identity.User, nil, deletedObject)
	if err != nil {
		return false, fmt.Errorf("could not audit log delete action: %v", err)
	}

	return deleted, nil
}

func Create(ctx context.Context, input *apicontracts.ApiKey, identity *identitymodels.Identity) (string, error) {
	getname := identity.GetId()

	_, totalUserApikeyCount, err := apikeyrepo.GetByFilter(ctx, &apicontracts.Filter{
		Filters: []apicontracts.FilterMetadata{
			{
				Field:     "identifier",
				Value:     getname,
				MatchMode: apicontracts.MatchModeEquals,
			},
		},
	})
	if err != nil {
		rlog.Errorc(ctx, "error when checking of apikey for identifier", err)
		return "", fmt.Errorf("error when checking of apikey for identifier already exist")
	}

	if totalUserApikeyCount >= 100 {
		return "", fmt.Errorf("too many apikeys")
	}

	_, dbcount, err := apikeyrepo.GetByFilter(ctx, &apicontracts.Filter{
		Filters: []apicontracts.FilterMetadata{
			{
				Field:     "identifier",
				Value:     getname,
				MatchMode: apicontracts.MatchModeEquals,
			},
			{
				Field:     "displayname",
				Value:     input.DisplayName,
				MatchMode: apicontracts.MatchModeEquals,
			},
		},
	})
	if err != nil {
		rlog.Errorc(ctx, "error when checking of apikey for identifier", err)
		return "", fmt.Errorf("error when checking of apikey for identifier already exist")
	}

	if dbcount > 0 {
		return "", fmt.Errorf("already a key for idenitifier: %s", input.Identifier)
	}

	uniqueId, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	universalId := uniqueId.String()
	hash := stringhelper.HashSHA512(universalId, []byte(viper.GetString(configconsts.API_KEY_SALT)))

	if identity.IsCluster() {
		input.Identifier = identity.GetId()
		if len(input.DisplayName) == 0 {
			input.DisplayName = identity.GetId()
		}
	}

	if input.Type == "" {
		if identity.IsCluster() {
			input.Type = apicontracts.ApiKeyTypeCluster
		}

		if identity.IsUser() {
			input.Type = apicontracts.ApiKeyTypeUser
		}
	}

	input.Hash = hash
	err = apikeyrepo.Create(ctx, *input)
	if err != nil {
		return "", err
	}

	_, err = auditlog.Create(ctx, "Migration of acl", models.AuditCategoryApikey, models.AuditActionCreate, identity.User, input, nil)
	if err != nil {
		return "", fmt.Errorf("could not audit log create action: %v", err)
	}

	return universalId, nil
}

func CreateForAgent(ctx context.Context, input *apicontracts.AgentApiKeyModel) (string, error) {
	if input == nil {
		return "", errors.New("input is nil")
	}

	if input.Type != apicontracts.ApiKeyTypeCluster {
		return "", errors.New("wrong api key type")
	}
	mongoctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	apikeys, err := apikeyrepo.GetByIdentifier(mongoctx, input.Identifier)
	if err != nil {
		return "", fmt.Errorf("error when checking apikey for identifier: %w", err)
	}

	if len(apikeys) > 0 {
		return "", fmt.Errorf("already a key for idenitifier: %s", input.Identifier)
	}

	uniqueId, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	apikey := apicontracts.ApiKey{}
	hash := stringhelper.HashSHA512(uniqueId.String(), []byte(mustGetApikeySalt()))
	existingCluster, err := clustersservice.FindByName(mongoctx, input.Identifier)
	if err != nil {
		return "", err
	}

	clusterExist := false
	if existingCluster != nil && existingCluster.ClusterId != "" {
		clusterExist = true
	}

	var clusterId string
	if !clusterExist {
		var datacenter *apicontracts.Datacenter
		if input.Provider == providers.ProviderTypeK3d || input.Provider == providers.ProviderTypeKind {
			input.DatacenterName = fmt.Sprintf("local-%s", input.Provider)
		}
		if input.Provider == providers.ProviderTypeTalos && viper.GetBool(configconsts.DEVELOPMENT) {
			input.DatacenterName = fmt.Sprintf("local-%s", input.Provider)
		}
		datacenter, err = datacenterRepo.FindByName(mongoctx, input.DatacenterName)
		if err != nil || datacenter == nil {
			return "could not find datacenter", err
		}

		clusterId, err = clustersservice.Create(ctx, &apicontracts.Cluster{
			ClusterName: input.Identifier,
			Workspace: apicontracts.Workspace{
				Name:         input.WorkspaceName,
				DatacenterID: datacenter.ID,
				Datacenter: apicontracts.Datacenter{
					Name: input.DatacenterName,
				},
			},
		})
		if err != nil {
			return "", err
		}
	} else {
		clusterId = existingCluster.ClusterId
	}

	if clusterId != "" {
		apikey.DisplayName = clusterId
		apikey.Identifier = clusterId
	} else {
		return "", errors.New("could not find cluster id")
	}
	apikey.ReadOnly = false
	apikey.Type = apicontracts.ApiKeyTypeCluster
	apikey.Hash = hash

	err = apikeyrepo.Create(mongoctx, apikey)
	if err != nil {
		return "", err
	}

	return uniqueId.String(), nil
}

func UpdateLastUsed(ctx context.Context, apikeyId string, identifier string) error {
	err := apikeyrepo.UpdateLastUsed(ctx, apikeyId, identifier)
	if err != nil {
		return err
	}

	return nil
}

func CreateOrRenew(ctx context.Context, req *apicontractsv2self.CreateOrRenewApikeyRequest) (*apicontractsv2self.CreateOrRenewApikeyResponse, error) {
	resp := &apicontractsv2self.CreateOrRenewApikeyResponse{}

	identity := rorcontext.GetIdentityFromRorContext(ctx)
	identifier := identity.GetId()

	expires := time.Now().Local().Add(time.Duration(req.Ttl) * time.Second)

	existing, _ := apikeyrepo.GetOwnByName(ctx, req.Name)

	token, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	hash := stringhelper.HashSHA512(token.String(), []byte(mustGetApikeySalt()))

	if existing != nil {

		err := apikeyrepo.UpdateOwnByName(ctx, req.Name, hash, expires)

		if err != nil {
			return nil, err
		}
		newkey := existing
		newkey.Hash = hash
		newkey.Expires = expires

		_, err = auditlog.Create(ctx, "Apikey updated", models.AuditCategoryApikey, models.AuditActionUpdate, identity.User, existing, newkey)
	} else {

		newkey := apicontracts.ApiKey{
			Identifier:  identifier,
			DisplayName: req.Name,
			Type:        apicontracts.ApiKeyType(identity.Type),
			Hash:        hash,
			Expires:     expires,
			Created:     time.Now(),
		}
		err = apikeyrepo.Create(ctx, newkey)
		if err != nil {
			return nil, err
		}
		_, err = auditlog.Create(ctx, "Apikey created", models.AuditCategoryApikey, models.AuditActionCreate, identity.User, newkey, nil)
	}

	resp.Token = token.String()
	resp.Expires = expires

	return resp, nil
}
