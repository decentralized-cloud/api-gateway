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

type serviceStatusResolver struct {
	logger          *zap.Logger
	resolverCreator types.ResolverCreatorContract
	serviceStatus   *edgeclusterGrpcContract.ServiceStatus
}

// NewServiceStatusResolver creates new instance of the serviceStatusResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// serviceStatus: Optional. The service status
// Returns the new instance or error if something goes wrong
func NewServiceStatusResolver(
	ctx context.Context,
	logger *zap.Logger,
	resolverCreator types.ResolverCreatorContract,
	serviceStatus *edgeclusterGrpcContract.ServiceStatus) (edgecluster.ServiceStatusResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if serviceStatus == nil {
		return nil, commonErrors.NewArgumentNilError("serviceStatus", "serviceStatus is required")
	}

	return &serviceStatusResolver{
		logger:          logger,
		resolverCreator: resolverCreator,
		serviceStatus:   serviceStatus,
	}, nil
}

// LoadBalancer contains the current status of the load-balancer
// ctx: Mandatory. Reference to the context
// Returns the load balancer status resolver or error if something goes wrong.
func (r *serviceStatusResolver) LoadBalancer(ctx context.Context) (edgecluster.LoadBalancerStatusResolverContract, error) {
	return r.resolverCreator.NewLoadBalancerStatusResolver(ctx, r.serviceStatus.LoadBalancer)
}

// KubeconfigContent returns the edge cluster Kubeconfig content
// ctx: Mandatory. Reference to the context
// Returns the edge cluster Kubeconfig content
func (r *serviceStatusResolver) Conditions(ctx context.Context) ([]edgecluster.ServiceConditionResolverContract, error) {
	response := []edgecluster.ServiceConditionResolverContract{}
	for _, condition := range r.serviceStatus.Conditions {
		if resolver, err := r.resolverCreator.NewServiceConditionResolver(ctx, condition); err != nil {
			return nil, err
		} else {
			response = append(response, resolver)
		}
	}

	return response, nil
}
