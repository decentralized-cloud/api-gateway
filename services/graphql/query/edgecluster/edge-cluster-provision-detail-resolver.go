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

type edgeClusterProvisionDetailResolver struct {
	logger          *zap.Logger
	resolverCreator types.ResolverCreatorContract
	provisionDetail *edgeclusterGrpcContract.EdgeClusterProvisionDetail
}

// NewEdgeClusterProvisionDetailResolver creates new instance of the edgeClusterProvisionDetailResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// provisionDetail: Optional. The edge cluster provisioning details
// Returns the new instance or error if something goes wrong
func NewEdgeClusterProvisionDetailResolver(
	ctx context.Context,
	logger *zap.Logger,
	resolverCreator types.ResolverCreatorContract,
	provisionDetail *edgeclusterGrpcContract.EdgeClusterProvisionDetail) (edgecluster.EdgeClusterProvisionDetailResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if provisionDetail == nil {
		return nil, commonErrors.NewArgumentNilError("provisionDetail", "provisionDetail is required")
	}

	return &edgeClusterProvisionDetailResolver{
		logger:          logger,
		resolverCreator: resolverCreator,
		provisionDetail: provisionDetail,
	}, nil
}

// Ingress returns the ingress details of the edge cluster master node
// ctx: Mandatory. Reference to the context
// Returns the ingress details of the edge cluster master node
func (r *edgeClusterProvisionDetailResolver) Ingress(ctx context.Context) (*[]edgecluster.IngressResolverContract, error) {
	if r.provisionDetail.Ingress == nil {
		return nil, nil
	}

	response := []edgecluster.IngressResolverContract{}

	for _, item := range r.provisionDetail.Ingress {
		resolver, err := r.resolverCreator.NewIngressResolver(ctx, item)

		if err != nil {
			return nil, err
		}

		response = append(response, resolver)
	}

	return &response, nil
}

// Ingress returns the ingress details of the edge cluster master node
// ctx: Mandatory. Reference to the context
// Returns the ingress details of the edge cluster master node
func (r *edgeClusterProvisionDetailResolver) Ports(ctx context.Context) (*[]edgecluster.PortResolverContract, error) {
	if r.provisionDetail.Ports == nil {
		return nil, nil
	}

	response := []edgecluster.PortResolverContract{}

	for _, item := range r.provisionDetail.Ports {
		resolver, err := r.resolverCreator.NewPortResolver(ctx, item)

		if err != nil {
			return nil, err
		}

		response = append(response, resolver)
	}

	return &response, nil
}

// KubeconfigContent returns the edge cluster Kubeconfig content
// ctx: Mandatory. Reference to the context
// Returns the edge cluster Kubeconfig content
func (r *edgeClusterProvisionDetailResolver) KubeconfigContent(ctx context.Context) *string {
	return &r.provisionDetail.KubeConfigContent
}
