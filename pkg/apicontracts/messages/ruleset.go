package messages

import (
	"errors"
)

var (
	ErrNoResource = errors.New("ruleset does not have entry for resource")
	ErrNoRule     = errors.New("resource does not have entry for rule")
)

type RulesetIdentityType string
type RulesetRuleType string
type RulesetLifetimeType string
type RulesetServiceType string

const (
	RulesetIdentityTypeUnknown  RulesetIdentityType = "unknown"
	RulesetIdentityTypeInternal RulesetIdentityType = "internal"
	RulesetIdentityTypeCluster  RulesetIdentityType = "cluster"
)

const (
	RulesetRuleTypeUnknown RulesetRuleType = "unknown"
	RulesetRuleTypeStarted RulesetRuleType = "started"
	RulesetRuleTypeCreated RulesetRuleType = "created"
	RulesetRuleTypeUpdated RulesetRuleType = "updated"
	RulesetRuleTypeCrashed RulesetRuleType = "crashed"
	RulesetRuleTypeDeleted RulesetRuleType = "deleted"
)

const (
	RulesetLifetimeTypeUnknown RulesetLifetimeType = "unknown"
	RulesetLifetimeTypeRegular RulesetLifetimeType = "regular"
	RulesetLifetimeTypeOneshot RulesetLifetimeType = "oneshot"
)

const (
	RulesetServiceTypeUnknown RulesetServiceType = "unknown"
	RulesetServiceTypeIgnore  RulesetServiceType = "ignore"
	RulesetServiceTypeSlack   RulesetServiceType = "slack"
	RulesetServiceTypeSms     RulesetServiceType = "sms"
)

type RulesetSlackModel struct {
	ChannelId string `json:"channelId"`
}

type RulesetRuleModel struct {
	Id string `json:"id" bson:"id"`

	Type     RulesetRuleType     `json:"type" bson:"type"`
	Lifetime RulesetLifetimeType `json:"lifetime" bson:"lifetime"`
	Service  RulesetServiceType  `json:"service" bson:"service"`

	Slack RulesetSlackModel `json:"slack,omitempty" bson:"slack,omitempty"`
}

type RulesetRuleInput struct {
	Type     RulesetRuleType     `json:"type"`
	Lifetime RulesetLifetimeType `json:"lifetime"`
	Service  RulesetServiceType  `json:"service"`

	Slack RulesetSlackModel `json:"slack,omitempty"`
}

type RulesetResourceModel struct {
	Id  string `json:"id"`
	Ref struct {
		// if present, only this is used for lookup
		Uid string `json:"uid"`

		// used for constructing a unique situation where uid is wildcard
		ApiVersion string `json:"apiVersion"`
		Kind       string `json:"kind"`
		Namespace  string `json:"namespace"`

		// only for presenting
		Name string `json:"name"`
	} `json:"ref"`
	Rules []RulesetRuleModel `json:"rules"`
}

type RulesetResourceInput struct {
	Uid        string `json:"uid"`
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Namespace  string `json:"namespace"`
}

type RulesetIdentityModel struct {
	Id   string              `json:"id" bson:"id"`
	Type RulesetIdentityType `json:"type" bson:"type"`
}

type RulesetModel struct {
	ID string `json:"id" bson:"_id,omitempty"`

	Identity  RulesetIdentityModel   `json:"identity" bson:"identity"`
	Resources []RulesetResourceModel `json:"resources" bson:"resources"`
}

func (ruleset *RulesetModel) FindResourceByUid(uid string) (*RulesetResourceModel, error) {
	for _, r := range ruleset.Resources {
		if r.Ref.Uid == uid {
			return &r, nil
		}
	}

	return nil, ErrNoResource
}

func (ruleset *RulesetModel) FindResourceByCombination(namespace string, apiVersion string, kind string) (*RulesetResourceModel, error) {
	for _, r := range ruleset.Resources {
		if r.Ref.Namespace == namespace && r.Ref.Kind == kind && r.Ref.ApiVersion == apiVersion {
			return &r, nil
		}
	}

	return nil, ErrNoResource
}

func (ruleset *RulesetModel) FindResourceById(id string) (*RulesetResourceModel, error) {
	for _, r := range ruleset.Resources {
		if r.Id == id {
			return &r, nil
		}
	}

	return nil, ErrNoResource
}

func (resource *RulesetResourceModel) FindRule(rule RulesetRuleType) (*RulesetRuleModel, error) {
	for _, r := range resource.Rules {
		if r.Type == rule {
			return &r, nil
		}
	}

	return nil, ErrNoRule
}
