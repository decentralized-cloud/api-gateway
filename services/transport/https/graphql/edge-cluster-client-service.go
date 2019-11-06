// Package graphql implements functions to expose api-gateway service endpoint using GraphQL protocol.
package graphql

import (
	"github.com/decentralized-cloud/api-gateway/services/configuration"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/edgecluster"
	edgeClusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"google.golang.org/grpc"
)

type edgeClusterClientService struct {
	serviceAddress string
}

// NewEdgeClusterClientService creates new instance of the edgeClusterClientService, setting up all dependencies and returns the instance
// configurationService: Mandatory. Reference to the configuration service
// Returns the new instance or error if something goes wrong
func NewEdgeClusterClientService(
	configurationService configuration.ConfigurationContract) (edgecluster.EdgeClusterClientContract, error) {
	if configurationService == nil {
		return nil, commonErrors.NewArgumentNilError("configurationService", "configurationService is required")
	}

	serviceAddress, err := configurationService.GetEdgeClusterServiceAddress()
	if err != nil {
		return nil, err
	}

	return &edgeClusterClientService{
		serviceAddress: serviceAddress,
	}, nil
}

// CreateClient creats a new edge cluster gRPC client and returns the connection
// and the client to the caller.
// Returns connection and the edge cluster gRPC client or error if something goes wrong.
func (service *edgeClusterClientService) CreateClient() (*grpc.ClientConn, edgeClusterGrpcContract.EdgeClusterServiceClient, error) {
	connection, err := grpc.Dial(service.serviceAddress, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	return connection, edgeClusterGrpcContract.NewEdgeClusterServiceClient(connection), nil
}
