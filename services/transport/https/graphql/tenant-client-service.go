// Package graphql implements functions to expose api-gateway service endpoint using GraphQL protocol.
package graphql

import (
	"github.com/decentralized-cloud/api-gateway/services/configuration"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/tenant"
	tenantGrpcContract "github.com/decentralized-cloud/tenant/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"google.golang.org/grpc"
)

type tenantClientService struct {
	serviceAddress string
}

// NewTenantClientService creates new instance of the tenantClientService, setting up all dependencies and returns the instance
// configurationService: Mandatory. Reference to the configuration service
// Returns the new instance or error if something goes wrong
func NewTenantClientService(
	configurationService configuration.ConfigurationContract) (tenant.TenantClientContract, error) {
	if configurationService == nil {
		return nil, commonErrors.NewArgumentNilError("configurationService", "configurationService is required")
	}

	serviceAddress, err := configurationService.GetTenantServiceAddress()
	if err != nil {
		return nil, err
	}

	return &tenantClientService{
		serviceAddress: serviceAddress,
	}, nil
}

// CreateClient creats a new tenant gRPC client and returns the connection
// and the client to the caller.
// Returns connection and the tenant gRPC client or error if something goes wrong.
func (service *tenantClientService) CreateClient() (*grpc.ClientConn, tenantGrpcContract.TenantServiceClient, error) {
	connection, err := grpc.Dial(service.serviceAddress, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	return connection, tenantGrpcContract.NewTenantServiceClient(connection), nil
}
