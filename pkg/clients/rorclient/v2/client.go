package rorclient

import (
	"context"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/clients"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/clientinterface"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/transports/transportinterface"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/v1clientset"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v2/v2clientset"

	"github.com/NorskHelsenett/ror/pkg/helpers/rorhealth"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/rorresourceowner"
)

type RorConfig struct {
	Host string
}

// Compile-time check to ensure RorClient implements RorClientInterface
var _ RorClientInterface = (*RorClient)(nil)

type RorClientInterface interface {
	transportinterface.RorCommonClientTransportInterface
	clientinterface.RorCommonClientApiInterfaceVersioned

	clientinterface.RorCommonClientOwnerInterface
	transportinterface.RorCommonClientTransportSetterInterface

	clients.CommonClient
}

type RorClient struct {
	ownerRef *rorresourceowner.RorResourceOwnerReference

	Transport transportinterface.RorTransport
	v1        clientinterface.RorCommonClientApiInterfaceV1
	v2        clientinterface.RorCommonClientApiInterfaceV2
}

func NewRorClient(transport transportinterface.RorTransport) *RorClient {
	return &RorClient{
		Transport: transport,
		v1:        v1clientset.NewV1ClientSet(transport),
		v2:        v2clientset.NewV2ClientSet(transport),
	}
}

func (c *RorClient) V1() clientinterface.RorCommonClientApiInterfaceV1 {
	return c.v1
}

func (c *RorClient) V2() clientinterface.RorCommonClientApiInterfaceV2 {
	return c.v2
}

// Ping checks the connection to the transport.
// Old version used error handling, use CheckConnection instead.
func (c *RorClient) Ping() bool {
	return c.PingWithContext(context.TODO())
}

// PingWithContext checks the connection to the transport with a context.
func (c *RorClient) PingWithContext(ctx context.Context) bool {
	return c.Transport.Ping(ctx)
}

// Transport related methods

// CheckConnection checks the connection to the transport.
func (c *RorClient) CheckConnection() error {
	return c.Transport.CheckConnection()
}

func (c *RorClient) GetRole() string {
	return c.Transport.GetRole()
}

// GetApiSecret gets the API secret from the transport.
func (c *RorClient) GetApiSecret() string {
	return c.Transport.GetApiSecret()
}

// SetTransport sets the transport for the RorClient.
func (c *RorClient) SetTransport(transport transportinterface.RorTransport) {
	c.Transport = transport
}

// GetOwnerref gets the owner reference for the RorClient.
func (c *RorClient) GetOwnerref() rorresourceowner.RorResourceOwnerReference {
	if c.ownerRef == nil {
		return rorresourceowner.RorResourceOwnerReference{Scope: aclmodels.Acl2ScopeUnknown, Subject: aclmodels.Acl2RorSubjecUnknown}
	}
	return *c.ownerRef
}

// SetOwnerref sets the owner reference for the RorClient.
func (c *RorClient) SetOwnerref(ownerref rorresourceowner.RorResourceOwnerReference) {
	c.ownerRef = &ownerref
}

// CheckHealth checks the health of the RorClient.
func (c *RorClient) CheckHealth(ctx context.Context) []rorhealth.Check {
	healthChecks := []rorhealth.Check{}
	if !c.Transport.Ping(ctx) {
		healthChecks = append(healthChecks, rorhealth.Check{
			ComponentID: "Transport",
			Status:      rorhealth.StatusFail,
			Output:      fmt.Sprintf("%s could not be connected", c.Transport.GetTransportName()),
		})
	}
	return healthChecks
}

// CheckHealthWithoutContext checks the health of the RorClient without a context.
func (c *RorClient) CheckHealthWithoutContext() []rorhealth.Check {
	healthChecks := []rorhealth.Check{}
	if !c.Transport.Ping(context.Background()) {
		healthChecks = append(healthChecks, rorhealth.Check{
			ComponentID: "Transport",
			Status:      rorhealth.StatusFail,
			Output:      fmt.Sprintf("%s could not be connected", c.Transport.GetTransportName()),
		})
	}
	return healthChecks
}
