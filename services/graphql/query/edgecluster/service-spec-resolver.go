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

type serviceSpecResolver struct {
	logger          *zap.Logger
	resolverCreator types.ResolverCreatorContract
	serviceSpec     *edgeclusterGrpcContract.ServiceSpec
}

// NewServiceSpecResolver creates new instance of the serviceSpecResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// serviceSpec: Optional. The service spec
// Returns the new instance or error if something goes wrong
func NewServiceSpecResolver(
	ctx context.Context,
	logger *zap.Logger,
	resolverCreator types.ResolverCreatorContract,
	serviceSpec *edgeclusterGrpcContract.ServiceSpec) (edgecluster.ServiceSpecResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if serviceSpec == nil {
		return nil, commonErrors.NewArgumentNilError("serviceSpec", "serviceSpec is required")
	}

	return &serviceSpecResolver{
		logger:          logger,
		resolverCreator: resolverCreator,
		serviceSpec:     serviceSpec,
	}, nil
}

// Ports is the list of ports that are exposed by this service
// ctx: Mandatory. Reference to the context
// Returns an array of service port resolver or error if something goes wrong.
func (r *serviceSpecResolver) Ports(ctx context.Context) ([]edgecluster.ServicePortResolverContract, error) {
	response := []edgecluster.ServicePortResolverContract{}
	for _, port := range r.serviceSpec.Ports {
		if resolver, err := r.resolverCreator.NewServicePortResolver(ctx, port); err != nil {
			return nil, err
		} else {
			response = append(response, resolver)
		}
	}

	return response, nil
}

// ClusterIPs is a list of IP addresses assigned to this service
// ctx: Mandatory. Reference to the context
// Returns the list of IP addresses assigned to this service
func (r *serviceSpecResolver) ClusterIPs(ctx context.Context) []string {
	return r.serviceSpec.ClusterIPs
}

// Type determines how the Service is exposed
// ctx: Mandatory. Reference to the context
// Returns how the service is exposed
func (r *serviceSpecResolver) Type(ctx context.Context) string {
	return r.serviceSpec.Type.String()

}

// ExternalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service
// ctx: Mandatory. Reference to the context
// Returns the list of IP addresses for which nodes in the cluster will also accept traffic for this service
func (r *serviceSpecResolver) ExternalIPs(ctx context.Context) []string {
	return r.serviceSpec.ExternalIPs
}

// ExternalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record)
// ctx: Mandatory. Reference to the context
// Returns the the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record)
func (r *serviceSpecResolver) ExternalName(ctx context.Context) *string {
	if r.serviceSpec.ExternalName == "" {
		return nil
	}

	return &r.serviceSpec.ExternalName
}
