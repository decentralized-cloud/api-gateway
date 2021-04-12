// Package edgecluster implements different edge cluster GraphQL query resovlers required by the GraphQL transport layer
package edgecluster

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type edgeClusterPodStatusResolver struct {
	logger          *zap.Logger
	resolverCreator types.ResolverCreatorContract
	status          *edgeclusterGrpcContract.PodStatus
}

// NewPodStatusResolver creates new instance of the edgeClusterPodStatusResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// status: Mandatory. Contains information about the edge cluster pod status.
// Returns the new instance or error if something goes wrong
func NewPodStatusResolver(
	ctx context.Context,
	logger *zap.Logger,
	resolverCreator types.ResolverCreatorContract,
	status *edgeclusterGrpcContract.PodStatus) (edgecluster.PodStatusResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if status == nil {
		return nil, commonErrors.NewArgumentNilError("status", "status is required")
	}

	return &edgeClusterPodStatusResolver{
		logger:          logger,
		resolverCreator: resolverCreator,
		status:          status,
	}, nil
}

// HostIP returns the IP address allocated to the pod. Routable at least within the cluster.
// ctx: Mandatory. Reference to the context
// Returns the IP address allocated to the pod. Routable at least within the cluster.
func (r *edgeClusterPodStatusResolver) HostIP(ctx context.Context) string {
	return r.status.HostIP
}

// PodIP returns the IP address allocated to the pod. Routable at least within the cluster.
// ctx: Mandatory. Reference to the context
// Returns the IP address allocated to the pod. Routable at least within the cluster.
func (r *edgeClusterPodStatusResolver) PodIP(ctx context.Context) string {
	return r.status.PodIP
}

// Conditions is an array of current observed pod conditions.
// ctx: Mandatory. Reference to the context
// Returns an array of current observed pod conditions resolver or error if something goes wrong.
func (r *edgeClusterPodStatusResolver) Conditions(ctx context.Context) ([]edgecluster.PodConditionResolverContract, error) {
	response := []edgecluster.PodConditionResolverContract{}
	for _, condition := range r.status.Conditions {
		if resolver, err := r.resolverCreator.NewPodConditionResolver(ctx, condition); err != nil {
			return nil, err
		} else {
			response = append(response, resolver)
		}
	}

	return response, nil
}
