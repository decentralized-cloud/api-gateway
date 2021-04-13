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

type provisionDetailsResolver struct {
	logger           *zap.Logger
	resolverCreator  types.ResolverCreatorContract
	provisionDetails *edgeclusterGrpcContract.ProvisionDetail
}

// NewProvisionDetailsResolver creates new instance of the provisionDetailsResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// provisionDetails: Optional. The edge cluster provisioning details
// Returns the new instance or error if something goes wrong
func NewProvisionDetailsResolver(
	ctx context.Context,
	logger *zap.Logger,
	resolverCreator types.ResolverCreatorContract,
	provisionDetails *edgeclusterGrpcContract.ProvisionDetail) (edgecluster.ProvisionDetailsResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if provisionDetails == nil {
		return nil, commonErrors.NewArgumentNilError("provisionDetails", "provisionDetails is required")
	}

	return &provisionDetailsResolver{
		logger:           logger,
		resolverCreator:  resolverCreator,
		provisionDetails: provisionDetails,
	}, nil
}

// LoadBalancer contains the current status of the load-balancer
// ctx: Mandatory. Reference to the context
// Returns the load balancer status resolver or error if something goes wrong.
func (r *provisionDetailsResolver) LoadBalancer(ctx context.Context) (edgecluster.LoadBalancerStatusResolverContract, error) {
	return r.resolverCreator.NewLoadBalancerStatusResolver(ctx, r.provisionDetails.LoadBalancer)
}

// KubeconfigContent returns the edge cluster Kubeconfig content
// ctx: Mandatory. Reference to the context
// Returns the edge cluster Kubeconfig content
func (r *provisionDetailsResolver) KubeconfigContent(ctx context.Context) *string {
	if r.provisionDetails.KubeConfigContent == "" {
		return nil
	}

	return &r.provisionDetails.KubeConfigContent
}

// Ports returns the ports that are exposed by the service
// ctx: Mandatory. Reference to the context
// Returns the ports that are exposed by the service
func (r *provisionDetailsResolver) Ports(ctx context.Context) []int32 {
	return r.provisionDetails.Ports
}
