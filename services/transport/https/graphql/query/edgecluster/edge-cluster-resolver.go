// Package edgelcuster implements different edge cluster GraphQL query resovlers required by the GraphQL transport layer
package edgeclster

import (
	"context"
	"errors"
	"strings"

	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/edgecluster"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	"github.com/graph-gophers/graphql-go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type edgeClusterResolver struct {
	logger          *zap.Logger
	resolverCreator types.ResolverCreatorContract
	edgeclusterID   string
	name            string
	edgeCluster     *edgeclusterGrpcContract.EdgeCluster
}

// NewEdgeClusterResolver creates new instance of the edgeClusterResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// logger: Mandatory. Reference to the logger service
// edgeClusterID: Mandatory. the edge cluster unique identifier
// Returns the new instance or error if something goes wrong
func NewEdgeClusterResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger,
	edgeClusterClientService edgecluster.EdgeClusterClientContract,
	edgeClusterID string,
	edgeCluster *edgeclusterGrpcContract.EdgeCluster) (edgecluster.EdgeClusterResolverContract, error) {
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

	if strings.Trim(edgeClusterID, " ") == "" {
		return nil, commonErrors.NewArgumentError("edgeClusterID", "edgeClusterID is required")
	}

	resolver := edgeClusterResolver{
		logger:          logger,
		resolverCreator: resolverCreator,
		edgeclusterID:   edgeClusterID,
	}

	if edgeCluster == nil {
		connection, edgeClusterServiceClient, err := edgeClusterClientService.CreateClient()
		if err != nil {
			return nil, err
		}

		defer func() {
			_ = connection.Close()
		}()

		response, err := edgeClusterServiceClient.ReadEdgeCluster(
			ctx,
			&edgeclusterGrpcContract.ReadEdgeClusterRequest{
				EdgeClusterID: edgeClusterID,
			})
		if err != nil {
			return nil, err
		}

		if response.Error != edgeclusterGrpcContract.Error_NO_ERROR {
			return nil, errors.New(response.ErrorMessage)
		}

		resolver.edgeCluster = response.EdgeCluster
	} else {
		resolver.edgeCluster = edgeCluster
	}

	return &resolver, nil
}

// ID returns edge cluster unique identifier
// ctx: Mandatory. Reference to the context
// Returns the edge cluster unique identifier
func (r *edgeClusterResolver) ID(ctx context.Context) graphql.ID {
	return graphql.ID(r.edgeclusterID)
}

// Name returns edge cluster name
// ctx: Mandatory. Reference to the context
// Returns the edge cluster name or error
func (r *edgeClusterResolver) Name(ctx context.Context) string {
	return r.edgeCluster.Name
}

// Tenant returns edge cluster tenant
// ctx: Mandatory. Reference to the context
// Returns the edge cluster tenant
func (r *edgeClusterResolver) Tenant(ctx context.Context) (edgecluster.EdgeClusterTenantResolverContract, error) {
	return r.resolverCreator.NewEdgeClusterTenantResolver(ctx, r.edgeCluster.TenantID)
}
