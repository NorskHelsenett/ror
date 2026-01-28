package rorclientv2self

import "github.com/NorskHelsenett/ror/pkg/apicontracts/v2/apicontractsv2self"

type SelfInterface interface {
	Get() (apicontractsv2self.SelfData, error)
	CreateOrUpdateApiKey(name string, ttl int64) (string, error)
}
