// Package edgecluster implements edge cluster mutation required by the GraphQL transport layer
package edgeclster

import (
	"context"
	"errors"
	"strings"

	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/edgecluster"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type updateEdgeCluster struct {
	logger                   *zap.Logger
	resolverCreator          types.ResolverCreatorContract
	edgeClusterClientService edgecluster.EdgeClusterClientContract
}

type updateEdgeClusterPayloadResolver struct {
	resolverCreator  types.ResolverCreatorContract
	clientMutationId *string
	edgeClusterID    string
	edgeCluster      *edgeclusterGrpcContract.EdgeCluster
	cursor           string
}

// NewUpdateEdgeCluster updates new instance of the updateEdgeCluster, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can update new instances of resolvers
// logger: Mandatory. Reference to the logger service
// Returns the new instance or error if something goes wrong
func NewUpdateEdgeCluster(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger,
	edgeClusterClientService edgecluster.EdgeClusterClientContract) (edgecluster.UpdateEdgeClusterContract, error) {
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

	return &updateEdgeCluster{
		logger:                   logger,
		resolverCreator:          resolverCreator,
		edgeClusterClientService: edgeClusterClientService,
	}, nil
}

// EdgeCluster returns the updated edge cluster inforamtion
// ctx: Mandatory. Reference to the context
// clientMutationId: Optional. Reference to the client mutation ID to correlate the request and response
// edgeClusterID: Mandatory. The edge cluster unique identifier
// edgeCluster: Optional. The edge cluster details
// cursor: Mandatory. The edge cluster cursor
// Returns the updated edge cluster inforamtion
func NewUpdateEdgeClusterPayloadResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	clientMutationId *string,
	edgeClusterID string,
	edgeCluster *edgeclusterGrpcContract.EdgeCluster,
	cursor string) (edgecluster.UpdateEdgeClusterPayloadResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if strings.Trim(edgeClusterID, " ") == "" {
		return nil, commonErrors.NewArgumentError("edgeClusterID", "edgeClusterID is required")
	}

	if edgeCluster == nil {
		return nil, commonErrors.NewArgumentNilError("edgeCluster", "edgeCluster is required")
	}

	if strings.Trim(cursor, " ") == "" {
		return nil, commonErrors.NewArgumentError("cursor", "cursor is required")
	}

	return &updateEdgeClusterPayloadResolver{
		resolverCreator:  resolverCreator,
		clientMutationId: clientMutationId,
		edgeClusterID:    edgeClusterID,
		edgeCluster:      edgeCluster,
		cursor:           cursor,
	}, nil
}

// MutateAndGetPayload update an existing edge cluster and returns the payload contains the result of updating an existing edge cluster
// ctx: Mandatory. Reference to the context
// args: Mandatory. Reference to the input argument contains edge cluster information to update
// Returns the updated edge cluster payload or error if something goes wrong
func (m *updateEdgeCluster) MutateAndGetPayload(
	ctx context.Context,
	args edgecluster.UpdateEdgeClusterInputArgument) (edgecluster.UpdateEdgeClusterPayloadResolverContract, error) {
	edgeClusterID := string(args.Input.EdgeClusterID)
	connection, edgeClusterServiceClient, err := m.edgeClusterClientService.CreateClient()
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = connection.Close()
	}()

	response, err := edgeClusterServiceClient.UpdateEdgeCluster(
		ctx,
		&edgeclusterGrpcContract.UpdateEdgeClusterRequest{
			EdgeClusterID:    edgeClusterID,
			K3SClusterSecret: args.Input.K3SClusterSecret,
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

	return m.resolverCreator.NewUpdateEdgeClusterPayloadResolver(
		ctx,
		args.Input.ClientMutationId,
		edgeClusterID,
		response.EdgeCluster,
		response.Cursor)
}

// EdgeCluster returns the updated edge cluster inforamtion
// ctx: Mandatory. Reference to the context
// Returns the updated edge cluster inforamtion
func (r *updateEdgeClusterPayloadResolver) EdgeCluster(ctx context.Context) (edgecluster.EdgeClusterTypeEdgeResolverContract, error) {
	return r.resolverCreator.NewEdgeClusterTypeEdgeResolver(
		ctx,
		r.edgeClusterID,
		r.cursor,
		r.edgeCluster)
}

// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
// ctx: Mandatory. Reference to the context
// Returns the provided clientMutationId as part of mutation request
func (r *updateEdgeClusterPayloadResolver) ClientMutationId(ctx context.Context) *string {
	return r.clientMutationId
}
