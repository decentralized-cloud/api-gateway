// Package edgelcuster implements different edge cluster GraphQL query resovlers required by the GraphQL transport layer
package edgeclster

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type edgeClusterProvisioningDetailResolver struct {
	logger             *zap.Logger
	provisioningDetail *edgeclusterGrpcContract.EdgeClusterProvisioningDetail
}

// NewEdgeClusterProvisioningDetailResolver creates new instance of the edgeClusterProvisioningDetailResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// provisioningDetail: Optional. The edge cluster provisioning details
// Returns the new instance or error if something goes wrong
func NewEdgeClusterProvisioningDetailResolver(
	ctx context.Context,
	logger *zap.Logger,
	provisioningDetail *edgeclusterGrpcContract.EdgeClusterProvisioningDetail) (edgecluster.EdgeClusterProvisioningDetailResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	return &edgeClusterProvisioningDetailResolver{
		logger:             logger,
		provisioningDetail: provisioningDetail,
	}, nil
}

// Status returns the edge cluster current status
// ctx: Mandatory. Reference to the context
// Returns the edge cluster current status
func (r *edgeClusterProvisioningDetailResolver) Status(ctx context.Context) *edgecluster.EdgeClusterStatus {
	if r.provisioningDetail == nil {
		return nil
	}

	var status edgecluster.EdgeClusterStatus

	switch r.provisioningDetail.Status {
	case edgeclusterGrpcContract.EdgeClusterStatus_Provisioning:
		status = edgecluster.Provisioning
	case edgeclusterGrpcContract.EdgeClusterStatus_Ready:
		status = edgecluster.Ready
	case edgeclusterGrpcContract.EdgeClusterStatus_Deleting:
		status = edgecluster.Deleting
	default:
		return nil
	}

	return &status
}

// PublicIPAddress returns the edge cluster public IP address
// ctx: Mandatory. Reference to the context
// Returns the edge cluster public IP address
func (r *edgeClusterProvisioningDetailResolver) PublicIPAddress(ctx context.Context) *string {
	if r.provisioningDetail == nil {
		return nil
	}

	return &r.provisioningDetail.PublicIPAddress
}

// KubeconfigContent returns the edge cluster Kubeconfig content
// ctx: Mandatory. Reference to the context
// Returns the edge cluster Kubeconfig content
func (r *edgeClusterProvisioningDetailResolver) KubeconfigContent(ctx context.Context) *string {
	if r.provisioningDetail == nil {
		return nil
	}

	return &r.provisioningDetail.KubeConfigContent
}
