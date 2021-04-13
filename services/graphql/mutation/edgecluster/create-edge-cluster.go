// Package edgecluster implements edge cluster mutation required by the GraphQL transport layer
package edgecluster

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
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
	resolverCreator   types.ResolverCreatorContract
	clientMutationId  *string
	edgeClusterID     string
	edgeClusterDetail *edgecluster.EdgeClusterDetail
	cursor            string
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
// clientMutationId: Optional. Reference to the client mutation ID to correlate the request and response
// edgeClusterID: Mandatory. The edge cluster unique identifier
// edgeClusterDetail: Mandatory. The edge cluster details
// cursor: Mandatory. The edge cluster cursor
// Returns the new instance or error if something goes wrong
func NewCreateEdgeClusterPayloadResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	clientMutationId *string,
	edgeClusterID string,
	edgeClusterDetail *edgecluster.EdgeClusterDetail,
	cursor string) (edgecluster.CreateEdgeClusterPayloadResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if strings.Trim(edgeClusterID, " ") == "" {
		return nil, commonErrors.NewArgumentError("edgeClusterID", "edgeClusterID is required")
	}

	if edgeClusterDetail == nil {
		return nil, commonErrors.NewArgumentNilError("edgeClusterDetail", "edgeClusterDetail is required")
	}

	if strings.Trim(cursor, " ") == "" {
		return nil, commonErrors.NewArgumentError("cursor", "cursor is required")
	}

	return &createEdgeClusterPayloadResolver{
		resolverCreator:   resolverCreator,
		clientMutationId:  clientMutationId,
		edgeClusterID:     edgeClusterID,
		edgeClusterDetail: edgeClusterDetail,
		cursor:            cursor,
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

	var clusterType edgeclusterGrpcContract.ClusterType

	if args.Input.ClusterType == "K3S" {
		clusterType = edgeclusterGrpcContract.ClusterType_K3S
	} else {
		return nil, fmt.Errorf("cluster type is not supported. Cluster type: %v", args.Input.ClusterType)
	}

	response, err := edgeClusterServiceClient.CreateEdgeCluster(
		ctx,
		&edgeclusterGrpcContract.CreateEdgeClusterRequest{
			EdgeCluster: &edgeclusterGrpcContract.EdgeCluster{
				ProjectID:     string(args.Input.ProjectID),
				Name:          args.Input.Name,
				ClusterSecret: args.Input.ClusterSecret,
				ClusterType:   clusterType,
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
		&edgecluster.EdgeClusterDetail{
			EdgeCluster:      response.EdgeCluster,
			ProvisionDetails: &edgeclusterGrpcContract.ProvisionDetail{},
		},
		response.Cursor)
}

// EdgeCluster returns the new edge cluster inforamtion
// ctx: Mandatory. Reference to the context
// Returns the new edge cluster inforamtion
func (r *createEdgeClusterPayloadResolver) EdgeCluster(ctx context.Context) (edgecluster.EdgeClusterTypeEdgeResolverContract, error) {
	return r.resolverCreator.NewEdgeClusterTypeEdgeResolver(
		ctx,
		r.edgeClusterID,
		r.cursor,
		r.edgeClusterDetail)
}

// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
// ctx: Mandatory. Reference to the context
// Returns the provided clientMutationId as part of mutation request
func (r *createEdgeClusterPayloadResolver) ClientMutationId(ctx context.Context) *string {
	return r.clientMutationId
}
