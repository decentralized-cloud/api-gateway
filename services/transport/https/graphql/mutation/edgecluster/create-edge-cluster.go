// Package edgecluster implements edge cluster mutation required by the GraphQL transport layer
package edgeclster

import (
	"context"
	"errors"

	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/edgecluster"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type createEdgeCluster struct {
	logger                   *zap.Logger
	resolverCreator          types.ResolverCreatorContract
	edgeClusterClientService edgecluster.EdgeClusterClientContract
}

type createEdgeClusterPayloadResolver struct {
	resolverCreator  types.ResolverCreatorContract
	clientMutationId *string
	edgeClusterID    string
	edgeCluster      *edgeclusterGrpcContract.EdgeCluster
}

// NewCreateEdgeCluster creates new instance of the createEdgeCluster, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// logger: Mandatory. Reference to the logger service
// Returns the new instance or error if something goes wrong
func NewCreateEdgeCluster(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger,
	edgeClusterClientService edgecluster.EdgeClusterClientContract) (edgecluster.CreateEdgeClusterContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if edgeClusterClientService == nil {
		return nil, commonErrors.NewArgumentNilError("edgeClusterClientService", "edgeClusterClientService is required")
	}

	return &createEdgeCluster{
		logger:                   logger,
		resolverCreator:          resolverCreator,
		edgeClusterClientService: edgeClusterClientService,
	}, nil
}

// NewCreateEdgeClusterPayloadResolver creates new instance of the createEdgeClusterPayloadResolvere, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// Returns the new instance or error if something goes wrong
func NewCreateEdgeClusterPayloadResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	clientMutationId *string,
	edgeClusterID string,
	edgeCluster *edgeclusterGrpcContract.EdgeCluster) (edgecluster.CreateEdgeClusterPayloadResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	return &createEdgeClusterPayloadResolver{
		resolverCreator:  resolverCreator,
		clientMutationId: clientMutationId,
		edgeClusterID:    edgeClusterID,
		edgeCluster:      edgeCluster,
	}, nil
}

// MutateAndGetPayload creates a new edge cluster and returns the payload contains the result of creating a new edge cluster
// ctx: Mandatory. Reference to the context
// args: Mandatory. Reference to the input argument contains edge cluster information to create
// Returns the new edge cluster payload or error if something goes wrong
func (m *createEdgeCluster) MutateAndGetPayload(
	ctx context.Context,
	args edgecluster.CreateEdgeClusterInputArgument) (edgecluster.CreateEdgeClusterPayloadResolverContract, error) {
	connection, edgeClusterServiceClient, err := m.edgeClusterClientService.CreateClient()
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = connection.Close()
	}()

	response, err := edgeClusterServiceClient.CreateEdgeCluster(
		ctx,
		&edgeclusterGrpcContract.CreateEdgeClusterRequest{
			EdgeCluster: &edgeclusterGrpcContract.EdgeCluster{
				TenantID: string(args.Input.TenantID),
				Name:     args.Input.Name,
			}})
	if err != nil {
		return nil, err
	}

	if response.Error != edgeclusterGrpcContract.Error_NO_ERROR {
		return nil, errors.New(response.ErrorMessage)
	}

	return m.resolverCreator.NewCreateEdgeClusterPayloadResolver(
		ctx,
		args.Input.ClientMutationId,
		response.EdgeClusterID,
		response.EdgeCluster)
}

// EdgeCluster returns the new edge cluster inforamtion
// ctx: Mandatory. Reference to the context
// Returns the new edge cluster inforamtion
func (r *createEdgeClusterPayloadResolver) EdgeCluster(ctx context.Context) (edgecluster.EdgeClusterTypeEdgeResolverContract, error) {
	resolver, err := r.resolverCreator.NewEdgeClusterTypeEdgeResolver(ctx, r.edgeClusterID, "Not implemented", r.edgeCluster)

	return resolver, err
}

// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
// ctx: Mandatory. Reference to the context
// Returns the provided clientMutationId as part of mutation request
func (r *createEdgeClusterPayloadResolver) ClientMutationId(ctx context.Context) *string {
	return r.clientMutationId
}
