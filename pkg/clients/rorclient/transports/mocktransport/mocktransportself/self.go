package mocktransportself

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/v2/apicontractsv2self"
)

type V2Client struct{}

func NewV2Client() *V2Client {
	return &V2Client{}
}

func (c *V2Client) Get() (apicontractsv2self.SelfData, error) {
	selfData := apicontractsv2self.SelfData{
		User: apicontractsv2self.SelfUser{
			Name:  "Mock Self",
			Email: "mock@example.com",
		},
		Type: "mock-type",
	}
	return selfData, nil
}

func (c *V2Client) CreateOrUpdateApiKey(name string, ttl int64) (string, error) {
	// Mock implementation - return a fake API key
	return "mock-api-key-" + name + "-12345", nil
}
