package mongodbseeding

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/NorskHelsenett/ror/cmd/api/services/rulesetsService"
	"github.com/NorskHelsenett/ror/internal/mongodbrepo/mongoTypes"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/apicontracts/messages"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"
	"github.com/NorskHelsenett/ror/pkg/rlog"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/models/providers"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckAndSeed(ctx context.Context) {
	seedPrices(ctx)

	if viper.GetBool(configconsts.DEVELOPMENT) {
		seedDatacenters(ctx)
		seedAclv2Items(ctx)
		seedDevelopmentRulesets(ctx)
		seedApiKeys(ctx)
		seedProjects(ctx)
	}

	//seedInternalRuleset(ctx)
	seedTasks(ctx)
	seedOperatorConfigs(ctx)
}

func seedInternalRuleset(ctx context.Context) {
	db := mongodb.GetMongoDb()
	collection := db.Collection("messagerulesets")

	switchboardCount, err := collection.CountDocuments(ctx, bson.D{{Key: "identity.type", Value: messages.RulesetIdentityTypeInternal}, {Key: "identity.id", Value: "internal-primary"}})
	if err != nil {
		rlog.Errorc(ctx, "could not check switchboard doc count", err)
		return
	}

	if switchboardCount != 0 {
		return
	}

	channelId := "C03U5CGFYQ6"

	if viper.GetBool(configconsts.DEVELOPMENT) {
		channelId = "C059W2B3Y4F"
	}

	switchboard, err := rulesetsService.CreateInternal(ctx)
	if err != nil {
		rlog.Errorc(ctx, "could not create internal switchboard", err)
		return
	}

	{
		input := messages.RulesetResourceInput{
			Uid: "ror-api",
		}

		resource, err := rulesetsService.AddResource(ctx, switchboard.ID, &input)
		if err != nil {
			rlog.Errorc(ctx, "could not add resource", err)
			return
		}

		if _, err := rulesetsService.AddResourceRule(ctx, switchboard.ID, resource.Id, &messages.RulesetRuleInput{
			Service:  messages.RulesetServiceTypeSlack,
			Lifetime: messages.RulesetLifetimeTypeRegular,
			Type:     messages.RulesetRuleTypeCrashed,

			Slack: messages.RulesetSlackModel{
				ChannelId: channelId,
			},
		}); err != nil {
			rlog.Errorc(ctx, "could not add resource event", err)
			return
		}

		if _, err := rulesetsService.AddResourceRule(ctx, switchboard.ID, resource.Id, &messages.RulesetRuleInput{
			Service:  messages.RulesetServiceTypeIgnore,
			Lifetime: messages.RulesetLifetimeTypeRegular,
			Type:     messages.RulesetRuleTypeCreated,
		}); err != nil {
			rlog.Errorc(ctx, "could not add resource event", err)
			return
		}

		if _, err := rulesetsService.AddResourceRule(ctx, switchboard.ID, resource.Id, &messages.RulesetRuleInput{
			Service:  messages.RulesetServiceTypeSlack,
			Lifetime: messages.RulesetLifetimeTypeRegular,
			Type:     messages.RulesetRuleTypeStarted,

			Slack: messages.RulesetSlackModel{
				ChannelId: channelId,
			},
		}); err != nil {
			rlog.Errorc(ctx, "could not add resource event", err)
			return
		}
	}

	{
		input := messages.RulesetResourceInput{
			Uid: "ror-ms-nhn",
		}

		resource, err := rulesetsService.AddResource(ctx, switchboard.ID, &input)
		if err != nil {
			rlog.Errorc(ctx, "could not add resource", err)
			return
		}

		_ = resource
	}

	{
		input := messages.RulesetResourceInput{
			Uid: "ror-ms-switchboard",
		}

		resource, err := rulesetsService.AddResource(ctx, switchboard.ID, &input)
		if err != nil {
			rlog.Errorc(ctx, "could not add resource", err)
			return
		}

		_ = resource
	}
}

func seedDevelopmentRulesets(ctx context.Context) {
	db := mongodb.GetMongoDb()
	collection := db.Collection("messagerulesets")
	switchboardCount, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		rlog.Errorc(ctx, "could not check switchboard doc count", err)
		return
	}

	if switchboardCount != 0 {
		return
	}
}

func seedPrices(ctx context.Context) {
	db := mongodb.GetMongoDb()
	collection := db.Collection("prices")
	priceCount, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		rlog.Errorc(ctx, "could not check prics doc count", err)
		return
	}

	if priceCount != 0 {
		return
	}

	insertResult, err := collection.InsertMany(ctx, []interface{}{
		mongoTypes.MongoPrice{
			ID:           primitive.NewObjectID(),
			Provider:     providers.ProviderTypeTanzu,
			MachineClass: "best-effort-medium",
			Cpu:          2,
			Memory:       int64(8),
			MemoryBytes:  int64(8238813184),
			Price:        900,
			From:         time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC),
		},
		mongoTypes.MongoPrice{
			ID:           primitive.NewObjectID(),
			Provider:     providers.ProviderTypeTanzu,
			MachineClass: "best-effort-large",
			Cpu:          4,
			Memory:       int64(16),
			MemoryBytes:  int64(16681451520),
			Price:        1800,
			From:         time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC),
		},
		mongoTypes.MongoPrice{
			ID:           primitive.NewObjectID(),
			Provider:     providers.ProviderTypeTanzu,
			MachineClass: "best-effort-xlarge",
			Cpu:          4,
			Memory:       int64(32),
			MemoryBytes:  int64(33567711232),
			Price:        2232,
			From:         time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC),
		},
	})

	if err != nil {
		rlog.Errorc(ctx, "could not insert prices", err)
		panic(err)
	}

	insertCount := len(insertResult.InsertedIDs)
	if insertCount > 0 {
		rlog.Infoc(ctx, "Seeded prices", rlog.Int("insert count", insertCount))
	}
}

func seedDatacenters(ctx context.Context) {
	db := mongodb.GetMongoDb()
	collection := db.Collection("datacenters")

	if !viper.GetBool(configconsts.DEVELOPMENT) {
		return
	}

	datacenters := []mongoTypes.MongoDatacenter{
		{
			ID:       primitive.NewObjectID(),
			Provider: providers.ProviderTypeUnknown,
			Location: mongoTypes.MongoDatacenterLocation{
				Country: "Norway",
				Region:  "Trøndelag",
			},
			Name:        "local",
			APIEndpoint: "localhost",
		},
		{
			ID:       primitive.NewObjectID(),
			Provider: providers.ProviderTypeK3d,
			Location: mongoTypes.MongoDatacenterLocation{
				Country: "Norway",
				Region:  "Trøndelag",
			},
			Name:        "local-k3d",
			APIEndpoint: "localhost",
		},
		{
			ID:       primitive.NewObjectID(),
			Provider: providers.ProviderTypeKind,
			Location: mongoTypes.MongoDatacenterLocation{
				Country: "Norway",
				Region:  "Trøndelag",
			},
			Name:        "local-kind",
			APIEndpoint: "localhost",
		},
		{
			ID:       primitive.NewObjectID(),
			Provider: providers.ProviderTypeTalos,
			Location: mongoTypes.MongoDatacenterLocation{
				Country: "Norway",
				Region:  "Trøndelag",
			},
			Name:        "local-talos",
			APIEndpoint: "localhost",
		},
		{
			ID:       primitive.NewObjectID(),
			Provider: providers.ProviderTypeTanzu,
			Location: mongoTypes.MongoDatacenterLocation{
				Country: "Norway",
				Region:  "Trøndelag",
			},
			Name:        "trd1",
			APIEndpoint: "ptr1-w02-cl01-api.sdi.nhn.no",
		},
		{
			ID:       primitive.NewObjectID(),
			Provider: providers.ProviderTypeTanzu,
			Location: mongoTypes.MongoDatacenterLocation{
				Country: "Norway",
				Region:  "Trøndelag",
			},
			Name:        "trd1-cl02",
			APIEndpoint: "ptr1-w02-cl02-api.sdi.nhn.no",
		},
		{
			ID:       primitive.NewObjectID(),
			Provider: providers.ProviderTypeTanzu,
			Location: mongoTypes.MongoDatacenterLocation{
				Country: "Norway",
				Region:  "Trøndelag",
			},
			Name:        "trd1cl02",
			APIEndpoint: "ptr1-w02-cl02-api.sdi.nhn.no",
		},
		{
			ID:       primitive.NewObjectID(),
			Provider: providers.ProviderTypeTanzu,
			Location: mongoTypes.MongoDatacenterLocation{
				Country: "Norway",
				Region:  "Oslo",
			},
			Name:        "osl1",
			APIEndpoint: "pos1-w02-cl01-api.sdi.nhn.no",
		},
	}

	for i := 0; i < len(datacenters); i++ {
		datacenterInput := datacenters[i]
		var datacenter *mongoTypes.MongoDatacenter
		findError := collection.FindOne(ctx, bson.M{"name": datacenterInput.Name}).Decode(&datacenter)
		if findError != nil {
			rlog.Errorc(ctx, "could not find datacenter", findError)
		}

		if datacenter != nil {
			continue
		}

		insertResult, err := collection.InsertOne(ctx, datacenterInput)
		errorMsg := fmt.Sprintf("could not insert datacenter of type: %s", datacenterInput.Provider)
		if err != nil {
			rlog.Errorc(ctx, errorMsg, err)
			panic(err)
		}
		if insertResult == nil {
			rlog.Errorc(ctx, errorMsg, err)
			panic(errors.New(errorMsg))
		} else {
			rlog.Infoc(ctx, "Inserted datacenter", rlog.String("datacenter", datacenterInput.Name), rlog.String("provider", string(datacenterInput.Provider)))
		}
	}
}

func seedTasks(ctx context.Context) {
	db := mongodb.GetMongoDb()
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	collection := db.Collection("tasks")
	defer cancel()
	taskCount, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		rlog.Errorc(ctx, "could not check tasks count", err)
		return
	}

	if taskCount != 0 {
		return
	}

	script1 := `#!/bin/bash
echo "task1"
helm repo add argo https://argoproj.github.io/argo-helm
helm install argocd argo/argo-cd --version $ARGOCD_VERSION --create-namespace -n argocd -f values.yaml
kubectl apply -f /app/rolebinding.yaml
`
	insertResult, err := collection.InsertMany(ctx, []interface{}{
		apicontracts.Task{
			Id:   primitive.NewObjectID(),
			Name: "argocd-installer",
			Config: apicontracts.TaskSpec{
				ImageName: "devops-base",
				Cmd:       "/app/entrypoint.sh",
				EnvVars: []apicontracts.KeyValue{
					{
						Key:   "ARGOCD_VERSION",
						Value: "5.55.0",
					},
				},
				BackOffLimit:     3,
				TimeOutInSeconds: 180,
				Version:          "1.0.0",
				Scripts: &apicontracts.TaskScripts{
					ScriptDirectory: "/scripts",
					FileNameAndData: []apicontracts.FileNameAndData{
						{
							FileName: "task1.sh",
							Data:     script1,
						},
						{
							FileName: "task2.sh",
							Data:     "echo 'task2'",
						},
					},
				},
				Secret: &apicontracts.TaskSecret{
					Path: "/data/",
					FileNameAndData: []apicontracts.FileNameAndData{
						{
							FileName: "values.yaml",
							Data:     "",
						},
					},
					GitSource: &apicontracts.TaskGitSource{
						Type:        apicontracts.Git,
						ContentPath: "config/config.yaml",
						GitConfig: apicontracts.GitConfig{
							Token:      "",
							User:       "",
							Repository: "https://helsegitlab.nhn.no/sdi/SDI-Infrastruktur/ror-jobs/argocd.git",
							Branch:     "feature/argocd-install",
							ProjectId:  413,
						},
					},
				},
			},
		},
		apicontracts.Task{
			Id:   primitive.NewObjectID(),
			Name: "cluster-agent-installer",
			Config: apicontracts.TaskSpec{
				ImageName:        "devops-base",
				Cmd:              "/app/entrypoint.sh",
				EnvVars:          make([]apicontracts.KeyValue, 0),
				BackOffLimit:     3,
				TimeOutInSeconds: 180,
				Version:          "1.0.0",
				Secret:           nil,
			},
		},
		apicontracts.Task{
			Id:   primitive.NewObjectID(),
			Name: "nhn-tooling-installer",
			Config: apicontracts.TaskSpec{
				ImageName:        "devops-base",
				Cmd:              "/app/entrypoint.sh",
				EnvVars:          make([]apicontracts.KeyValue, 0),
				BackOffLimit:     3,
				TimeOutInSeconds: 180,
				Version:          "1.0.0",
				Secret:           nil,
			},
		},
	})

	if err != nil {
		rlog.Errorc(ctx, "could not insert tasks: ", err)
		panic(err)
	}

	insertCount := len(insertResult.InsertedIDs)
	if insertCount > 0 {
		rlog.Infoc(ctx, "Seeded tasks", rlog.Int("insert count", insertCount))
	}
}

func seedOperatorConfigs(ctx context.Context) {
	db := mongodb.GetMongoDb()
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	collection := db.Collection("operatorconfigs")
	defer cancel()
	totalCount, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		rlog.Errorc(ctx, "could not check operator config count", err)
		return
	}

	if totalCount != 0 {
		return
	}

	insertResult, err := collection.InsertMany(ctx, []interface{}{
		mongoTypes.MongoOperatorConfig{
			Id:         primitive.NewObjectID(),
			ApiVersion: "github.com/NorskHelsenett/ror/v1/config",
			Kind:       "ror-operator",
			Spec: &mongoTypes.MongoOperatorSpec{
				ImagePostfix: "ror-operator:0.0.1",
				Tasks: []mongoTypes.MongoOperatorTask{
					{
						Index:   0,
						Name:    "argocd-installer",
						Version: "0.0.1",
						RunOnce: true,
					},
					{
						Index:   1,
						Name:    "cluster-agent-installer",
						Version: "1.0.0",
						RunOnce: false,
					},
					{
						Index:   1,
						Name:    "nhn-tooling-installer",
						Version: "1.0.2",
						RunOnce: false,
					},
				},
			},
		},
	})

	if err != nil {
		rlog.Errorc(ctx, "could not insert tasks", err)
		panic(err)
	}

	insertCount := len(insertResult.InsertedIDs)
	if insertCount > 0 {
		rlog.Infoc(ctx, "seeded tasks", rlog.Int("insert count", insertCount))
	}
}

func seedAclv2Items(ctx context.Context) {
	db := mongodb.GetMongoDb()
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	collection := db.Collection("acl")
	defer cancel()
	totalCount, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		rlog.Errorc(ctx, "could not check acl count", err)
		return
	}

	if totalCount != 0 {
		return
	}

	insertResult, err := collection.InsertMany(ctx, []interface{}{
		aclmodels.AclV2ListItem{
			Version:    2,
			Group:      "A-T1-SDI-DevOps-Operators@ror.dev",
			Scope:      aclmodels.Acl2ScopeRor,
			Subject:    aclmodels.Acl2Subject(aclmodels.Acl2RorSubjectGlobal),
			Access:     aclmodels.AclV2ListItemAccess{Read: true, Create: true, Update: true, Delete: true, Owner: true},
			Kubernetes: aclmodels.AclV2ListItemKubernetes{Logon: true},
			IssuedBy:   "system@ror.dev",
			Created:    time.Now(),
		},
		aclmodels.AclV2ListItem{
			Version:    2,
			Group:      "service-nhn@ror.system",
			Scope:      aclmodels.Acl2ScopeRor,
			Subject:    aclmodels.Acl2Subject(aclmodels.Acl2ScopeCluster),
			Access:     aclmodels.AclV2ListItemAccess{Read: true, Create: false, Update: true, Delete: false, Owner: false},
			Kubernetes: aclmodels.AclV2ListItemKubernetes{Logon: false},
			IssuedBy:   "system@ror.dev",
			Created:    time.Now(),
		},
		aclmodels.AclV2ListItem{
			Version:    2,
			Group:      "service-audit@ror.system",
			Scope:      aclmodels.Acl2ScopeRor,
			Subject:    aclmodels.Acl2Subject(aclmodels.Acl2RorSubjectGlobal),
			Access:     aclmodels.AclV2ListItemAccess{Read: true, Create: false, Update: false, Delete: false, Owner: false},
			Kubernetes: aclmodels.AclV2ListItemKubernetes{Logon: false},
			IssuedBy:   "system@ror.dev",
			Created:    time.Now(),
		},
		aclmodels.AclV2ListItem{
			Version:    2,
			Group:      "service-msswitchboard@ror.system",
			Scope:      aclmodels.Acl2ScopeRor,
			Subject:    aclmodels.Acl2Subject(aclmodels.Acl2RorSubjectGlobal),
			Access:     aclmodels.AclV2ListItemAccess{Read: true, Create: true, Update: true, Delete: false, Owner: false},
			Kubernetes: aclmodels.AclV2ListItemKubernetes{Logon: false},
			IssuedBy:   "system@ror.dev",
			Created:    time.Now(),
		},
		aclmodels.AclV2ListItem{
			Version:    2,
			Group:      "service-mstanzu@ror.system",
			Scope:      aclmodels.Acl2ScopeRor,
			Subject:    aclmodels.Acl2Subject(aclmodels.Acl2RorSubjectGlobal),
			Access:     aclmodels.AclV2ListItemAccess{Read: true, Create: true, Update: true, Delete: true, Owner: false},
			Kubernetes: aclmodels.AclV2ListItemKubernetes{Logon: false},
			IssuedBy:   "system@ror.dev",
			Created:    time.Now(),
		},
		aclmodels.AclV2ListItem{
			Version:    2,
			Group:      "service-mskind@ror.system",
			Scope:      aclmodels.Acl2ScopeRor,
			Subject:    aclmodels.Acl2Subject(aclmodels.Acl2RorSubjectGlobal),
			Access:     aclmodels.AclV2ListItemAccess{Read: true, Create: true, Update: true, Delete: true, Owner: false},
			Kubernetes: aclmodels.AclV2ListItemKubernetes{Logon: false},
			IssuedBy:   "system@ror.dev",
			Created:    time.Now(),
		},
		aclmodels.AclV2ListItem{
			Version:    2,
			Group:      "service-msvulnerability@ror.system",
			Scope:      aclmodels.Acl2ScopeRor,
			Subject:    aclmodels.Acl2Subject(aclmodels.Acl2RorSubjectGlobal),
			Access:     aclmodels.AclV2ListItemAccess{Read: true, Create: true, Update: true, Delete: false, Owner: false},
			Kubernetes: aclmodels.AclV2ListItemKubernetes{Logon: false},
			IssuedBy:   "system@ror.dev",
			Created:    time.Now(),
		},
		aclmodels.AclV2ListItem{
			Version:    2,
			Group:      "service-msslack@ror.system",
			Scope:      aclmodels.Acl2ScopeRor,
			Subject:    aclmodels.Acl2Subject(aclmodels.Acl2RorSubjectGlobal),
			Access:     aclmodels.AclV2ListItemAccess{Read: true, Create: true, Update: true, Delete: false, Owner: false},
			Kubernetes: aclmodels.AclV2ListItemKubernetes{Logon: false},
			IssuedBy:   "system@ror.dev",
			Created:    time.Now(),
		},
		aclmodels.AclV2ListItem{
			Version:    2,
			Group:      "service-mstalos@ror.system",
			Scope:      aclmodels.Acl2ScopeRor,
			Subject:    aclmodels.Acl2Subject(aclmodels.Acl2RorSubjectGlobal),
			Access:     aclmodels.AclV2ListItemAccess{Read: true, Create: true, Update: true, Delete: true, Owner: false},
			Kubernetes: aclmodels.AclV2ListItemKubernetes{Logon: false},
			IssuedBy:   "system@ror.dev",
			Created:    time.Now(),
		},
	})

	if err != nil {
		rlog.Errorc(ctx, "could not insert acl items", err)
		panic(err)
	}

	insertCount := len(insertResult.InsertedIDs)
	if insertCount > 0 {
		rlog.Infoc(ctx, "seeded acl items", rlog.Int("insert count", insertCount))
	}
}

func seedApiKeys(ctx context.Context) {
	db := mongodb.GetMongoDb()
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	collection := db.Collection("apikeys")
	defer cancel()
	totalCount, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		rlog.Errorc(ctx, "could not check apikey count", err)
		return
	}

	tanzuKey := apicontracts.ApiKey{
		Identifier:  "mstanzu",
		DisplayName: "mstanzu",
		Type:        "Service",
		ReadOnly:    false,
		Expires:     time.Time{},
		Created:     time.Now(),
		LastUsed:    time.Time{},
		Hash:        "246bd9a1958f8a52e5c31f0b832d2243a72d210472e3daa19449d21bed25664cbd4076864b5ea1c8732a4aeaf81d82c24653eb4cb593d6e8e876f6d2a996c629",
	}
	talosKey := apicontracts.ApiKey{
		Identifier:  "mstalos",
		DisplayName: "mstalos",
		Type:        "Service",
		ReadOnly:    false,
		Expires:     time.Time{},
		Created:     time.Now(),
		LastUsed:    time.Time{},
		Hash:        "14f9a160e00b172f8da8ad487dac537b8399f3e5f7696d6d9e4d10a96d1cc9992fcbcf96579df0fa5323c0bb89f6dab875956899d822ddf695b6a6b592a21874",
	}
	kindKey := apicontracts.ApiKey{
		Identifier:  "mskind",
		DisplayName: "mskind",
		Type:        "Service",
		ReadOnly:    false,
		Expires:     time.Time{},
		Created:     time.Now(),
		LastUsed:    time.Time{},
		Hash:        "5b52ea5f512b1630efa24b9a86dbb23a6b97174220c262b0c6d6af11120149f06c1aae8afaeaedc4e1cf4a20cbf1ab81031df33a9af35573c1e5795b01b5f9d2",
	}
	msvulnerabilityKey := apicontracts.ApiKey{
		Identifier:  "msvulnerability",
		DisplayName: "msvulnerability",
		Type:        "Service",
		ReadOnly:    false,
		Expires:     time.Time{},
		Created:     time.Now(),
		LastUsed:    time.Time{},
		Hash:        "af0342b0a470675ab5a526b7a3db18faf3781cacafa82474bed940d9e35c3aa1f99fcff21da3b0fee7010962bb2722d5c6a65ace0eca871acbd61c586da6bb47",
	}
	msslackKey := apicontracts.ApiKey{
		Identifier:  "msslack",
		DisplayName: "msslack",
		Type:        "Service",
		ReadOnly:    false,
		Expires:     time.Time{},
		Created:     time.Now(),
		LastUsed:    time.Time{},
		Hash:        "93e1613a8c9cbff6724a0935b81d2611a1afd8ebf42ffe0a4c529923baff5186d2a91d5ece0f348936e471181ca8f7228872c1014bf24623e53de66a53d040b1",
	}
	msswitchboardKey := apicontracts.ApiKey{
		Identifier:  "msswitchboard",
		DisplayName: "msswitchboard",
		Type:        "Service",
		ReadOnly:    false,
		Expires:     time.Time{},
		Created:     time.Now(),
		LastUsed:    time.Time{},
		Hash:        "dc9874d499431e92eb30f607b87e19efa3806c57344358f9bd392ba72ef5ffde80f4a942c3398bf8379ac0364bffba0ce24b9344ce183b4fd33807be9046d2fa",
	}

	if totalCount == 0 {
		insertResult, err := collection.InsertMany(ctx, []interface{}{
			tanzuKey,
			kindKey,
			talosKey,
			msvulnerabilityKey,
			msslackKey,
			msswitchboardKey,
		})
		if err != nil {
			rlog.Errorc(ctx, "could not insert apikeys items", err)
			panic(err)
		}

		insertCount := len(insertResult.InsertedIDs)
		if insertCount > 0 {
			rlog.Infoc(ctx, "seeded apikey items", rlog.Int("insert count", insertCount))
		}
	}

	var tanzuResult apicontracts.ApiKey
	err = collection.FindOne(ctx, bson.M{"identifier": "mstanzu"}).Decode(&tanzuResult)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			_, err = collection.InsertOne(ctx, tanzuKey)
			if err != nil {
				rlog.Errorc(ctx, "could not insert mstanzu key", err)
				panic(err)
			}
		} else {
			rlog.Errorc(ctx, "could not find mstanzu key", err)
			return
		}
	}

	var talosResult apicontracts.ApiKey
	err = collection.FindOne(ctx, bson.M{"identifier": "mstalos"}).Decode(&talosResult)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			_, err = collection.InsertOne(ctx, talosKey)
			if err != nil {
				rlog.Errorc(ctx, "could not insert mstalos key", err)
				panic(err)
			}
		} else {
			rlog.Errorc(ctx, "could not find mstalos key", err)
			return
		}
	}

	var kindResult apicontracts.ApiKey
	err = collection.FindOne(ctx, bson.M{"identifier": "mskind"}).Decode(&kindResult)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			_, err = collection.InsertOne(ctx, kindKey)
			if err != nil {
				rlog.Errorc(ctx, "could not insert mskind key", err)
				panic(err)
			}
		} else {
			rlog.Errorc(ctx, "could not find mskind key", err)
			return
		}
	}

	var vulnerabilityResult apicontracts.ApiKey
	err = collection.FindOne(ctx, bson.M{"identifier": "msvulnerability"}).Decode(&vulnerabilityResult)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			_, err = collection.InsertOne(ctx, msvulnerabilityKey)
			if err != nil {
				rlog.Errorc(ctx, "could not insert msvulnerability key", err)
				panic(err)
			}
		} else {
			rlog.Errorc(ctx, "could not find msvulnerability key", err)
			return
		}
	}
}

func seedProjects(ctx context.Context) {
	db := mongodb.GetMongoDb()
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	collection := db.Collection("projects")
	defer cancel()
	totalCount, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		rlog.Errorc(ctx, "could not check projects count", err)
		return
	}

	if totalCount != 0 {
		return
	}

	roles := make([]mongoTypes.MongoProjectRole, 0)
	roles = append(roles, mongoTypes.MongoProjectRole{
		ContactInfo: mongoTypes.MongoProjectContactInfo{
			UPN:   "p1@p1.no",
			Email: "p1@p1.no",
			Phone: "12345678",
		},
		RoleDefinition: apicontracts.ProjectRoleOwner,
	})
	roles = append(roles, mongoTypes.MongoProjectRole{
		ContactInfo: mongoTypes.MongoProjectContactInfo{
			UPN:   "p1@p1.no",
			Email: "p1@p1.no",
			Phone: "12345678",
		},
		RoleDefinition: apicontracts.ProjectRoleResponsible,
	})
	tags := map[string]string{}

	insertResult, err := collection.InsertMany(ctx, []interface{}{
		mongoTypes.MongoProject{
			ID:          primitive.NewObjectID(),
			Name:        "Project 1",
			Description: "Project 1 description",
			Created:     time.Now(),
			Updated:     time.Now(),
			Active:      true,
			ProjectMetadata: mongoTypes.MongoProjectMetadata{
				Roles: roles,
				Billing: mongoTypes.MongoBilling{
					Workorder: "w-p1-123456",
				},
				ServiceTags: tags,
			},
		},
	})
	if err != nil {
		rlog.Errorc(ctx, "could not insert apikeys items", err)
		panic(err)
	}

	insertCount := len(insertResult.InsertedIDs)
	if insertCount > 0 {
		rlog.Infoc(ctx, "seeded apikey items", rlog.Int("insert count", insertCount))
	}
}
