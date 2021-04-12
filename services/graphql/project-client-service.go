// Package graphql implements functions to expose api-gateway service endpoint using GraphQL protocol.
package graphql

import (
	"github.com/decentralized-cloud/api-gateway/services/configuration"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/project"
	projectGrpcContract "github.com/decentralized-cloud/project/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"google.golang.org/grpc"
)

type projectClientService struct {
	serviceAddress string
}

// NewProjectClientService creates new instance of the projectClientService, setting up all dependencies and returns the instance
// configurationService: Mandatory. Reference to the configuration service
// Returns the new instance or error if something goes wrong
func NewProjectClientService(
	configurationService configuration.ConfigurationContract) (project.ProjectClientContract, error) {
	if configurationService == nil {
		return nil, commonErrors.NewArgumentNilError("configurationService", "configurationService is required")
	}

	serviceAddress, err := configurationService.GetProjectServiceAddress()
	if err != nil {
		return nil, err
	}

	return &projectClientService{
		serviceAddress: serviceAddress,
	}, nil
}

// CreateClient creats a new project gRPC client and returns the connection
// and the client to the caller.
// Returns connection and the project gRPC client or error if something goes wrong.
func (service *projectClientService) CreateClient() (*grpc.ClientConn, projectGrpcContract.ServiceClient, error) {
	connection, err := grpc.Dial(service.serviceAddress, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	return connection, projectGrpcContract.NewServiceClient(connection), nil
}
