// Package edgecluster implements different edge cluster GraphQL query resovlers required by the GraphQL transport layer
package edgecluster

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	"github.com/graph-gophers/graphql-go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type edgeClusterResolver struct {
	logger                   *zap.Logger
	resolverCreator          types.ResolverCreatorContract
	edgeclusterID            string
	edgeClusterDetail        *edgecluster.EdgeClusterDetail
	edgeClusterClientService edgecluster.EdgeClusterClientContract
}

// NewEdgeClusterResolver creates new instance of the edgeClusterResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// logger: Mandatory. Reference to the logger service
// edgeClusterID: Mandatory. the edge cluster unique identifier
// edgeClusterDetail: Optional. The edge cluster details, if provided, the value be used instead of contacting  the edge cluster service
// Returns the new instance or error if something goes wrong
func NewEdgeClusterResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger,
	edgeClusterClientService edgecluster.EdgeClusterClientContract,
	edgeClusterID string,
	edgeClusterDetail *edgecluster.EdgeClusterDetail) (edgecluster.EdgeClusterResolverContract, error) {
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
		logger:                   logger,
		resolverCreator:          resolverCreator,
		edgeclusterID:            edgeClusterID,
		edgeClusterClientService: edgeClusterClientService,
	}

	if edgeClusterDetail == nil {
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

		resolver.edgeClusterDetail = &edgecluster.EdgeClusterDetail{
			EdgeCluster:      response.EdgeCluster,
			ProvisionDetails: response.ProvisionDetail,
		}
	} else {
		resolver.edgeClusterDetail = edgeClusterDetail
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
	return r.edgeClusterDetail.EdgeCluster.Name
}

// ClusterSecret returns edge cluster secret
// ctx: Mandatory. Reference to the context
// Returns the edge cluster secret
func (r *edgeClusterResolver) ClusterSecret(ctx context.Context) string {
	return r.edgeClusterDetail.EdgeCluster.ClusterSecret
}

// ClusterType returns the edge cluster current type
// ctx: Mandatory. Reference to the context
// Returns the edge cluster current type or error if something went wrong
func (r *edgeClusterResolver) ClusterType(ctx context.Context) (clusterType string, err error) {
	if r.edgeClusterDetail.EdgeCluster.ClusterType == edgeclusterGrpcContract.ClusterType_K3S {
		clusterType = "K3S"

		return
	}

	err = fmt.Errorf("cluster type is not supported. Cluster type: %v", r.edgeClusterDetail.EdgeCluster.ClusterType)

	return
}

// Project returns edge cluster project
// ctx: Mandatory. Reference to the context
// Returns the edge cluster project resolver or error if something goes wrong.
func (r *edgeClusterResolver) Project(ctx context.Context) (edgecluster.EdgeClusterProjectResolverContract, error) {
	return r.resolverCreator.NewEdgeClusterProjectResolver(ctx, r.edgeClusterDetail.EdgeCluster.ProjectID)
}

// ProvisionDetails returns edge cluster provisioning detail
// ctx: Mandatory. Reference to the context
// Returns the edge cluster provisioning detail resolver or error if something goes wrong.
func (r *edgeClusterResolver) ProvisionDetails(ctx context.Context) (edgecluster.ProvisionDetailsResolverContract, error) {
	return r.resolverCreator.NewProvisionDetailsResolver(ctx, r.edgeClusterDetail.ProvisionDetails)
}

// Nodes returns the resolver that resolves the nodes that are part of the given edge cluster or error if something goes wrong.
// ctx: Mandatory. Reference to the context
// Returns the resolver that resolves the nodes that are part of the given edge cluster or error if something goes wrong.
func (r *edgeClusterResolver) Nodes(ctx context.Context) ([]edgecluster.NodeResolverContract, error) {
	connection, edgeClusterServiceClient, err := r.edgeClusterClientService.CreateClient()
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = connection.Close()
	}()

	listEdgeClusterNodesResponse, err := edgeClusterServiceClient.ListEdgeClusterNodes(
		ctx,
		&edgeclusterGrpcContract.ListEdgeClusterNodesRequest{
			EdgeClusterID: r.edgeclusterID,
		})
	if err != nil {
		// if failed to retrieve nodes, return empty list
		return []edgecluster.NodeResolverContract{}, err
	}

	if listEdgeClusterNodesResponse.Error != edgeclusterGrpcContract.Error_NO_ERROR {
		return nil, errors.New(listEdgeClusterNodesResponse.ErrorMessage)
	}

	response := []edgecluster.NodeResolverContract{}
	for _, node := range listEdgeClusterNodesResponse.Nodes {
		if resolver, err := r.resolverCreator.NewEdgeClusterNodeResolver(ctx, node); err != nil {
			return nil, err
		} else {
			response = append(response, resolver)
		}
	}

	return response, nil
}

// Pods returns the resolver that resolves the pods that are part of the given edge cluster or error if something goes wrong.
// ctx: Mandatory. Reference to the context
// args: Mandatory. Reference to the query argument
// Returns the resolver that resolves the pods that are part of the given edge cluster or error if something goes wrong.
func (r *edgeClusterResolver) Pods(ctx context.Context, args edgecluster.EdgeClusterPodInputArgument) ([]edgecluster.PodResolverContract, error) {
	connection, edgeClusterServiceClient, err := r.edgeClusterClientService.CreateClient()
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = connection.Close()
	}()

	request := &edgeclusterGrpcContract.ListEdgeClusterPodsRequest{
		EdgeClusterID: r.edgeclusterID,
	}

	if args.NodeName != nil {
		request.NodeName = *args.NodeName
	}

	if args.Namespace != nil {
		request.Namespace = *args.Namespace
	}

	listEdgeClusterNodesResponse, err := edgeClusterServiceClient.ListEdgeClusterPods(
		ctx,
		request)
	if err != nil {
		return nil, err
	}

	if listEdgeClusterNodesResponse.Error != edgeclusterGrpcContract.Error_NO_ERROR {
		// if failed to retrieve pods, return empty list
		return []edgecluster.PodResolverContract{}, nil
	}

	response := []edgecluster.PodResolverContract{}
	for _, pod := range listEdgeClusterNodesResponse.Pods {
		if resolver, err := r.resolverCreator.NewEdgeClusterPodResolver(ctx, pod); err != nil {
			return nil, err
		} else {
			response = append(response, resolver)
		}
	}

	return response, nil
}

// Services returns the resolver that resolves the services that are part of the given edge cluster or error if something goes wrong.
// ctx: Mandatory. Reference to the context
// args: Mandatory. Reference to the query argument
// Returns the resolver that resolves the services that are part of the given edge cluster or error if something goes wrong.
func (r *edgeClusterResolver) Services(ctx context.Context, args edgecluster.EdgeClusterServiceInputArgument) ([]edgecluster.ServiceResolverContract, error) {
	connection, edgeClusterServiceClient, err := r.edgeClusterClientService.CreateClient()
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = connection.Close()
	}()

	request := &edgeclusterGrpcContract.ListEdgeClusterServicesRequest{
		EdgeClusterID: r.edgeclusterID,
	}

	if args.Namespace != nil {
		request.Namespace = *args.Namespace
	}

	listEdgeClusterNodesResponse, err := edgeClusterServiceClient.ListEdgeClusterServices(
		ctx,
		request)
	if err != nil {
		return nil, err
	}

	if listEdgeClusterNodesResponse.Error != edgeclusterGrpcContract.Error_NO_ERROR {
		// if failed to retrieve services, return empty list
		return []edgecluster.ServiceResolverContract{}, nil
	}

	response := []edgecluster.ServiceResolverContract{}
	for _, service := range listEdgeClusterNodesResponse.Services {
		if resolver, err := r.resolverCreator.NewEdgeClusterServiceResolver(ctx, service); err != nil {
			return nil, err
		} else {
			response = append(response, resolver)
		}
	}

	return response, nil
}
