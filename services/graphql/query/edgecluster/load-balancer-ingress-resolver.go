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

type loadBalancerIngressResolver struct {
	logger              *zap.Logger
	resolverCreator     types.ResolverCreatorContract
	loadBalancerIngress *edgeclusterGrpcContract.LoadBalancerIngress
}

// NewLoadBalancerIngressResolver creates new instance of the loadBalancerIngressResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// loadBalancerIngress: Mandatory. The load balancer ingress details
// Returns the new instance or error if something goes wrong
func NewLoadBalancerIngressResolver(
	ctx context.Context,
	logger *zap.Logger,
	resolverCreator types.ResolverCreatorContract,
	loadBalancerIngress *edgeclusterGrpcContract.LoadBalancerIngress) (edgecluster.LoadBalancerIngressResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if loadBalancerIngress == nil {
		return nil, commonErrors.NewArgumentNilError("loadBalancerIngress", "loadBalancerIngress is required")
	}

	return &loadBalancerIngressResolver{
		logger:              logger,
		resolverCreator:     resolverCreator,
		loadBalancerIngress: loadBalancerIngress,
	}, nil
}

// IP is set for load-balancer ingress points that are IP based
// ctx: Mandatory. Reference to the context
// Returns the IP that is set for load-balancer ingress points that are IP based
func (r *loadBalancerIngressResolver) IP(ctx context.Context) string {
	return r.loadBalancerIngress.Ip
}

// Hostname is set for load-balancer ingress points that are DNS based
// ctx: Mandatory. Reference to the context
// Returns the hostname that is set for load-balancer ingress points that are DNS based
func (r *loadBalancerIngressResolver) Hostname(ctx context.Context) string {
	return r.loadBalancerIngress.Hostname
}

// PortStatus is a list of records of service ports
// If used, every port defined in the service should have an entry in it
// ctx: Mandatory. Reference to the context
// Returns an array of port status resolver or error if something goes wrong.
func (r *loadBalancerIngressResolver) PortStatus(ctx context.Context) ([]edgecluster.PortStatusResolverContract, error) {
	response := []edgecluster.PortStatusResolverContract{}

	for _, portStatus := range r.loadBalancerIngress.PortStatus {
		if resolver, err := r.resolverCreator.NewPortStatusResolverContract(ctx, portStatus); err != nil {
			return nil, err
		} else {
			response = append(response, resolver)
		}
	}

	return response, nil
}
