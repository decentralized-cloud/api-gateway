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

type loadBalancerStatusResolver struct {
	logger             *zap.Logger
	resolverCreator    types.ResolverCreatorContract
	loadBalancerStatus *edgeclusterGrpcContract.LoadBalancerStatus
}

// NewLoadBalancerStatusResolver creates new instance of the loadBalancerStatusResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// loadBalancerStatus: Mandatory. The load balancer status details
// Returns the new instance or error if something goes wrong
func NewLoadBalancerStatusResolver(
	ctx context.Context,
	logger *zap.Logger,
	resolverCreator types.ResolverCreatorContract,
	loadBalancerStatus *edgeclusterGrpcContract.LoadBalancerStatus) (edgecluster.LoadBalancerStatusResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if loadBalancerStatus == nil {
		return nil, commonErrors.NewArgumentNilError("loadBalancerStatus", "loadBalancerStatus is required")
	}

	return &loadBalancerStatusResolver{
		logger:             logger,
		resolverCreator:    resolverCreator,
		loadBalancerStatus: loadBalancerStatus,
	}, nil
}

// Ingress is a list containing ingress points for the load-balancer.
// Traffic intended for the service should be sent to these ingress points.
// ctx: Mandatory. Reference to the context
// Returns an array of load balancer ingress resolver or error if something goes wrong.
func (r *loadBalancerStatusResolver) Ingress(ctx context.Context) ([]edgecluster.LoadBalancerIngressResolverContract, error) {
	response := []edgecluster.LoadBalancerIngressResolverContract{}
	for _, ingress := range r.loadBalancerStatus.Ingress {
		if resolver, err := r.resolverCreator.NewLoadBalancerIngressResolver(ctx, ingress); err != nil {
			return nil, err
		} else {
			response = append(response, resolver)
		}
	}

	return response, nil
}
