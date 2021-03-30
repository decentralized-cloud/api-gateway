// Package edgecluster implements different edge cluster GraphQL query resovlers required by the GraphQL transport layer
package edgecluster

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type edgeClusterPodSpecResolver struct {
	logger *zap.Logger
	spec   *edgeclusterGrpcContract.EdgeClusterPodSpec
}

// NewEdgeClusterPodSpecResolver creates new instance of the edgeClusterPodSpecResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// spec: Mandatory. Contains information about the edge cluster pod spec.
// Returns the new instance or error if something goes wrong
func NewEdgeClusterPodSpecResolver(
	ctx context.Context,
	logger *zap.Logger,
	spec *edgeclusterGrpcContract.EdgeClusterPodSpec) (edgecluster.EdgeClusterPodSpecResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if spec == nil {
		return nil, commonErrors.NewArgumentNilError("spec", "spec is required")
	}

	return &edgeClusterPodSpecResolver{
		logger: logger,
		spec:   spec,
	}, nil
}

// NodeName returns the name of the node where the Pod is deployed into.
// ctx: Mandatory. Reference to the context
// Returns the name of the node where the Pod is deployed into.
func (r *edgeClusterPodSpecResolver) NodeName(ctx context.Context) string {
	return r.spec.NodeName
}
