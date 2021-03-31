// Package graphql implements functions to expose api-gateway service endpoint using GraphQL protocol.
package graphql

import (
	"context"

	mutationedgecluster "github.com/decentralized-cloud/api-gateway/services/graphql/mutation/edgecluster"
	mutationproject "github.com/decentralized-cloud/api-gateway/services/graphql/mutation/project"
	"github.com/decentralized-cloud/api-gateway/services/graphql/query"
	queryedgecluster "github.com/decentralized-cloud/api-gateway/services/graphql/query/edgecluster"
	queryproject "github.com/decentralized-cloud/api-gateway/services/graphql/query/project"
	queryrelay "github.com/decentralized-cloud/api-gateway/services/graphql/query/relay"
	"github.com/decentralized-cloud/api-gateway/services/graphql/root"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/project"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/relay"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	projectGrpcContract "github.com/decentralized-cloud/project/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type resolverCreator struct {
	logger                   *zap.Logger
	projectClientService     project.ProjectClientContract
	edgeClusterClientService edgecluster.EdgeClusterClientContract
}

// NewResolverCreator creates new instance of the resolverCreator, setting up all dependencies and returns the instance
// logger: Mandatory. Reference to the logger service
// configurationService: Mandatory. Reference to the configuration service
// projectClientService: Mandatory. the project client service that creates gRPC connection and client to the project
// Returns the new instance or error if something goes wrong
func NewResolverCreator(
	logger *zap.Logger,
	projectClientService project.ProjectClientContract,
	edgeClusterClientService edgecluster.EdgeClusterClientContract) (types.ResolverCreatorContract, error) {
	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if projectClientService == nil {
		return nil, commonErrors.NewArgumentNilError("projectClientService", "projectClientService is required")
	}

	if edgeClusterClientService == nil {
		return nil, commonErrors.NewArgumentNilError("edgeClusterClientService", "edgeClusterClientService is required")
	}

	return &resolverCreator{
		logger:                   logger,
		projectClientService:     projectClientService,
		edgeClusterClientService: edgeClusterClientService,
	}, nil
}

// NewPageInfoResolver creates new PageInfoResolverContract and returns it
// ctx: Mandatory. Reference to the context
// startCursor: Mandatory. Reference to the start cursor
// endCursor: Mandatory. Reference to the end cursor
// hasNextPage: Mandatory. Reference to the value indicates whether returned page has next page to be retrieved
// hasPreviousPage: Mandatory. Reference to the value indicates whether returned page has previous page to be retrieved
// Returns the PageInfoResolverContract or error if something goes wrong
func (creator *resolverCreator) NewPageInfoResolver(
	ctx context.Context,
	startCursor *string,
	endCursor *string,
	hasNextPage bool,
	hasPreviousPage bool) (relay.PageInfoResolverContract, error) {
	return queryrelay.NewPageInfoResolver(
		startCursor,
		endCursor,
		hasNextPage,
		hasPreviousPage)
}

// NewRootResolver creates new RootResolverContract and returns it
// ctx: Mandatory. Reference to the context
// Returns the RootResolverContract or error if something goes wrong
func (creator *resolverCreator) NewRootResolver(ctx context.Context) (types.RootResolverContract, error) {
	return root.NewRootResolver(
		ctx,
		creator,
		creator.logger)
}

// NewUserResolver creates new UserResolverContract and returns it
// ctx: Mandatory. Reference to the context
// userID: Mandatory. The user unique identifier
// Returns the UserResolverContract or error if something goes wrong
func (creator *resolverCreator) NewUserResolver(
	ctx context.Context,
	userID string) (types.UserResolverContract, error) {
	return query.NewUserResolver(
		ctx,
		creator,
		creator.logger,
		userID,
		creator.projectClientService,
		creator.edgeClusterClientService)
}

// NewProjectResolver creates new ProjectResolverContract and returns it
// ctx: Mandatory. Reference to the context
// projectID: Mandatory. The project unique identifier
// projectDetail: Optional. The tennat details, if provided, the value be used instead of contacting  the edge cluster service
// Returns the ProjectResolverContract or error if something goes wrong
func (creator *resolverCreator) NewProjectResolver(
	ctx context.Context,
	projectID string,
	projectDetail *project.ProjectDetail) (project.ProjectResolverContract, error) {
	return queryproject.NewProjectResolver(
		ctx,
		creator,
		creator.logger,
		creator.projectClientService,
		creator.edgeClusterClientService,
		projectID,
		projectDetail)
}

// NewProjectTypeEdgeResolver creates new ProjectTypeEdgeResolverContract and returns it
// ctx: Mandatory. Reference to the context
// projectID: Mandatory. The project unique identifier
// projectDetail: Optional. The tennat details, if provided, the value be used instead of contacting  the edge cluster service
// cursor: Mandatory. The cursor
// Returns the ProjectTypeEdgeResolverContract or error if something goes wrong
func (creator *resolverCreator) NewProjectTypeEdgeResolver(
	ctx context.Context,
	projectID string,
	cursor string,
	projectDetail *project.ProjectDetail) (project.ProjectTypeEdgeResolverContract, error) {
	return queryproject.NewProjectTypeEdgeResolver(
		ctx,
		creator,
		projectID,
		projectDetail,
		cursor)
}

// NewProjectTypeConnectionResolver creates new ProjectTypeConnectionResolverContract and returns it
// ctx: Mandatory. Reference to the context
// projects: Mandatory. Reference the list of projects
// hasPreviousPage: Mandatory. Indicates whether more edges exist prior to the set defined by the clients arguments
// hasNextPage: Mandatory. Indicates whether more edges exist following the set defined by the clients arguments
// totalCount: Mandatory. The total count of matched projects
// Returns the ProjectTypeConnectionResolverContract or error if something goes wrong
func (creator *resolverCreator) NewProjectTypeConnectionResolver(
	ctx context.Context,
	projects []*projectGrpcContract.ProjectWithCursor,
	hasPreviousPage bool,
	hasNextPage bool,
	totalCount int32) (project.ProjectTypeConnectionResolverContract, error) {
	return queryproject.NewProjectTypeConnectionResolver(
		ctx,
		creator,
		projects,
		hasPreviousPage,
		hasNextPage,
		totalCount)
}

// NewEdgeClusterResolver creates new EdgeClusterResolverContract and returns it
// ctx: Mandatory. Reference to the context
// edgeClusterID: Mandatory. The edge cluster unique identifier
// edgeClusterDetail: Optional. The edge cluster details, if provided, the value be used instead of contacting  the edge cluster service
// Returns the EdgeClusterResolverContract or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterResolver(
	ctx context.Context,
	edgeClusterID string,
	edgeClusterDetail *edgecluster.EdgeClusterDetail) (edgecluster.EdgeClusterResolverContract, error) {
	return queryedgecluster.NewEdgeClusterResolver(
		ctx,
		creator,
		creator.logger,
		creator.edgeClusterClientService,
		edgeClusterID,
		edgeClusterDetail)
}

// NewEdgeClusterObjectMetadataResolverContract creates new instance of the EdgeClusterObjectMetadataResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// metadata: Mandatory. Contains the object metadata.
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterObjectMetadataResolverContract(
	ctx context.Context,
	metadata *edgeclusterGrpcContract.EdgeClusterObjectMetadata) (edgecluster.EdgeClusterObjectMetadataResolverContract, error) {
	return queryedgecluster.NewEdgeClusterObjectMetadataResolver(
		ctx,
		creator.logger,
		metadata)
}

// NewEdgeClusterTypeEdgeResolver creates new EdgeClusterTypeEdgeResolverContract and returns it
// ctx: Mandatory. Reference to the context
// edgeClusterID: Mandatory. The edge cluster unique identifier
// cursor: Mandatory. The cursor
// edgeClusterDetail: Optional. The edge cluster details, if provided, the value be used instead of contacting  the edge cluster service
// Returns the EdgeClusterTypeEdgeResolverContract or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterTypeEdgeResolver(
	ctx context.Context,
	edgeClusterID string,
	cursor string,
	edgeClusterDetail *edgecluster.EdgeClusterDetail) (edgecluster.EdgeClusterTypeEdgeResolverContract, error) {
	return queryedgecluster.NewEdgeClusterTypeEdgeResolver(
		ctx,
		creator,
		edgeClusterID,
		edgeClusterDetail,
		cursor)
}

// NewEdgeClusterTypeConnectionResolver creates new EdgeClusterTypeConnectionResolverContract and returns it
// ctx: Mandatory. Reference to the context
// edgeClusters: Mandatory. Reference the list of edge clusters
// hasPreviousPage: Mandatory. Indicates whether more edges exist prior to the set defined by the clients arguments
// hasNextPage: Mandatory. Indicates whether more edges exist following the set defined by the clients arguments
// totalCount: Mandatory. The total count of matched edge clusters
// Returns the EdgeClusterTypeConnectionResolverContract or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterTypeConnectionResolver(
	ctx context.Context,
	edgeclusters []*edgeclusterGrpcContract.EdgeClusterWithCursor,
	hasPreviousPage bool,
	hasNextPage bool,
	totalCount int32) (edgecluster.EdgeClusterTypeConnectionResolverContract, error) {
	return queryedgecluster.NewEdgeClusterTypeConnectionResolver(
		ctx,
		creator,
		edgeclusters,
		hasPreviousPage,
		hasNextPage,
		totalCount)
}

// NewEdgeClusterProjectResolver creates new EdgeClusterTenatnResolverContract and returns it
// ctx: Mandatory. Reference to the context
// projectID: Mandatory. The project unique identifier
// Returns the EdgeClusterTenatnResolverContract or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterProjectResolver(
	ctx context.Context,
	projectID string) (edgecluster.EdgeClusterProjectResolverContract, error) {
	return queryedgecluster.NewEdgeClusterProjectResolver(
		ctx,
		creator,
		creator.logger,
		projectID)
}

// NewEdgeClusterProvisionDetailResolver creates new EdgeClusterProvisionDetailResolverContract and returns it
// ctx: Mandatory. Reference to the context
// ProvisionDetail: Mandatory. The edge cluster provisioning details
// Returns the EdgeClusterProvisionDetailResolverContract or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterProvisionDetailResolver(
	ctx context.Context,
	provisionDetail *edgeclusterGrpcContract.EdgeClusterProvisionDetail) (edgecluster.EdgeClusterProvisionDetailResolverContract, error) {
	return queryedgecluster.NewEdgeClusterProvisionDetailResolver(
		ctx,
		creator.logger,
		creator,
		provisionDetail)
}

// NewIngressResolver creates new instance of the ingressResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// ingress: Mandatory. The ingress details
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewIngressResolver(
	ctx context.Context,
	ingress *edgeclusterGrpcContract.Ingress) (edgecluster.IngressResolverContract, error) {
	return queryedgecluster.NewIngressResolver(
		ctx,
		creator.logger,
		ingress)
}

// NewPortResolver creates new instance of the PortResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// Port: Mandatory. The PostStatus details
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewPortResolver(
	ctx context.Context,
	port *edgeclusterGrpcContract.Port) (edgecluster.PortResolverContract, error) {
	return queryedgecluster.NewPortResolver(
		ctx,
		creator.logger,
		port)
}

// NewEdgeClusterNodeResolver creates new instance of the edgeClusterNodeResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// node: Mandatory. Contains information about the edge cluster node.
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterNodeResolver(
	ctx context.Context,
	node *edgeclusterGrpcContract.EdgeClusterNode) (edgecluster.EdgeClusterNodeResolverContract, error) {
	return queryedgecluster.NewEdgeClusterNodeResolver(
		ctx,
		creator.logger,
		creator,
		node)
}

// NewEdgeClusterNodeStatusResolver creates new instance of the edgeClusterNodeStatusResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// status: Mandatory. Contains information about the edge cluster node status.
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterNodeStatusResolver(
	ctx context.Context,
	status *edgeclusterGrpcContract.EdgeClusterNodeStatus) (edgecluster.EdgeClusterNodeStatusResolverContract, error) {
	return queryedgecluster.NewEdgeClusterNodeStatusResolver(
		ctx,
		creator.logger,
		creator,
		status)
}

// NewEdgeClusterNodeConditionResolver creates new instance of the edgeClusterNodeConditionResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// condition: Mandatory. Contains condition information for a node.
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterNodeConditionResolver(
	ctx context.Context,
	condition *edgeclusterGrpcContract.EdgeClusterNodeCondition) (edgecluster.EdgeClusterNodeConditionResolverContract, error) {
	return queryedgecluster.NewEdgeClusterNodeConditionResolver(
		ctx,
		creator.logger,
		condition)
}

// NewEdgeClusterNodeAddressResolverContract creates new instance of the edgeClusterNodeAddressResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// nodeAddresss: Mandatory. Contains information for the edge cluster node's address.
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterNodeAddressResolverContract(
	ctx context.Context,
	nodeAddresss *edgeclusterGrpcContract.EdgeClusterNodeAddress) (edgecluster.EdgeClusterNodeAddressResolverContract, error) {
	return queryedgecluster.NewEdgeClusterNodeAddressResolverContract(
		ctx,
		creator.logger,
		nodeAddresss)
}

// NewEdgeClusterNodeSystemInfoResolverContract creates new instance of the edgeClusterNodeSystemInfoResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// nodeInfo: Mandatory. Contains a set of ids/uuids to uniquely identify the node.
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterNodeSystemInfoResolverContract(
	ctx context.Context,
	nodeInfo *edgeclusterGrpcContract.EdgeClusterNodeSystemInfo) (edgecluster.EdgeClusterNodeSystemInfoResolverContract, error) {
	return queryedgecluster.NewEdgeClusterNodeSystemInfoResolverContract(
		ctx,
		creator.logger,
		nodeInfo)
}

// NewCreateProject creates new instance of the createProject, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewCreateProject(ctx context.Context) (project.CreateProjectContract, error) {
	return mutationproject.NewCreateProject(
		ctx,
		creator,
		creator.logger,
		creator.projectClientService)
}

// NewCreateProjectPayloadResolver creates new instance of the createProjectPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// clientMutationId: Optional. Reference to the client mutation ID to correlate the request and response
// projectID: Mandatory. The project unique identifier
// projectDetail: Mandatory. The project details
// cursor: Mandatory. The edge cluster cursor
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewCreateProjectPayloadResolver(
	ctx context.Context,
	clientMutationId *string,
	projectID string,
	projectDetail *project.ProjectDetail,
	cursor string) (project.CreateProjectPayloadResolverContract, error) {
	return mutationproject.NewCreateProjectPayloadResolver(
		ctx,
		creator,
		clientMutationId,
		projectID,
		projectDetail,
		cursor)
}

// NewUpdateProject creates new instance of the updateProject, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewUpdateProject(ctx context.Context) (project.UpdateProjectContract, error) {
	return mutationproject.NewUpdateProject(
		ctx,
		creator,
		creator.logger,
		creator.projectClientService)
}

// NewUpdateProjectPayloadResolver creates new instance of the updateProjectPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// clientMutationId: Optional. Reference to the client mutation ID to correlate the request and response
// projectID: Mandatory. The project unique identifier
// projectDetail: Mandatory. The project details
// cursor: Mandatory. The edge cluster cursor
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewUpdateProjectPayloadResolver(
	ctx context.Context,
	clientMutationId *string,
	projectID string,
	projectDetail *project.ProjectDetail,
	cursor string) (project.UpdateProjectPayloadResolverContract, error) {
	return mutationproject.NewUpdateProjectPayloadResolver(
		ctx,
		creator,
		clientMutationId,
		projectID,
		projectDetail,
		cursor)
}

// NewDeleteProject creates new instance of the deleteProject, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewDeleteProject(ctx context.Context) (project.DeleteProjectContract, error) {
	return mutationproject.NewDeleteProject(
		ctx,
		creator,
		creator.logger,
		creator.projectClientService)
}

// NewDeleteProjectPayloadResolver creates new instance of the deleteProjectPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// projectID: Mandatory. The project unique identifier
// clientMutationId: Optional. Reference to the client mutation ID to correlate the request and response
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewDeleteProjectPayloadResolver(
	ctx context.Context,
	projectID string,
	clientMutationId *string) (project.DeleteProjectPayloadResolverContract, error) {
	return mutationproject.NewDeleteProjectPayloadResolver(
		ctx,
		creator,
		projectID,
		clientMutationId)
}

// NewCreateEdgeCluster creates new instance of the createEdgeCluster, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewCreateEdgeCluster(ctx context.Context) (edgecluster.CreateEdgeClusterContract, error) {
	return mutationedgecluster.NewCreateEdgeCluster(
		ctx,
		creator,
		creator.logger,
		creator.edgeClusterClientService)
}

// NewCreateEdgeClusterPayloadResolver creates new instance of the createEdgeClusterPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// clientMutationId: Optional. Reference to the client mutation ID to correlate the request and response
// edgeClusterID: Mandatory. The edge cluster unique identifier
// edgeClusterDetail: Mandatory. The edge cluster details
// cursor: Mandatory. The edge cluster cursor
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewCreateEdgeClusterPayloadResolver(
	ctx context.Context,
	clientMutationId *string,
	edgeClusterID string,
	edgeClusterDetail *edgecluster.EdgeClusterDetail,
	cursor string) (edgecluster.CreateEdgeClusterPayloadResolverContract, error) {
	return mutationedgecluster.NewCreateEdgeClusterPayloadResolver(
		ctx,
		creator,
		clientMutationId,
		edgeClusterID,
		edgeClusterDetail,
		cursor)
}

// NewUpdateEdgeCluster creates new instance of the updateEdgeCluster, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewUpdateEdgeCluster(ctx context.Context) (edgecluster.UpdateEdgeClusterContract, error) {
	return mutationedgecluster.NewUpdateEdgeCluster(
		ctx,
		creator,
		creator.logger,
		creator.edgeClusterClientService)
}

// NewUpdateEdgeClusterPayloadResolver creates new instance of the updateEdgeClusterPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// clientMutationId: Optional. Reference to the client mutation ID to correlate the request and response
// edgeClusterID: Mandatory. The edge cluster unique identifier
// edgeClusterDetail: Mandatory. The edge cluster details
// cursor: Mandatory. The edge cluster cursor
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewUpdateEdgeClusterPayloadResolver(
	ctx context.Context,
	clientMutationId *string,
	edgeClusterID string,
	edgeClusterDetail *edgecluster.EdgeClusterDetail,
	cursor string) (edgecluster.UpdateEdgeClusterPayloadResolverContract, error) {
	return mutationedgecluster.NewUpdateEdgeClusterPayloadResolver(
		ctx,
		creator,
		clientMutationId,
		edgeClusterID,
		edgeClusterDetail,
		cursor)
}

// NewDeleteEdgeCluster creates new instance of the deleteEdgeCluster, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewDeleteEdgeCluster(ctx context.Context) (edgecluster.DeleteEdgeClusterContract, error) {
	return mutationedgecluster.NewDeleteEdgeCluster(
		ctx,
		creator,
		creator.logger,
		creator.edgeClusterClientService)
}

// NewDeleteEdgeClusterPayloadResolver creates new instance of the deleteEdgeClusterPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// edgeClusterID: Mandatory. The edge cluster unique identifier
// clientMutationId: Optional. Reference to the client mutation ID to correlate the request and response
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewDeleteEdgeClusterPayloadResolver(
	ctx context.Context,
	edgeClusterID string,
	clientMutationId *string) (edgecluster.DeleteEdgeClusterPayloadResolverContract, error) {
	return mutationedgecluster.NewDeleteEdgeClusterPayloadResolver(
		ctx,
		creator,
		edgeClusterID,
		clientMutationId)
}

// NewEdgeClusterPodResolver creates new instance of the NewEdgeClusterPodResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// pod: Mandatory. Contains information about the edge cluster pod
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterPodResolver(
	ctx context.Context,
	pod *edgeclusterGrpcContract.EdgeClusterPod) (edgecluster.EdgeClusterPodResolverContract, error) {
	return queryedgecluster.NewEdgeClusterPodResolver(
		ctx,
		creator.logger,
		creator,
		pod)
}

// NewEdgeClusterPodStatusResolver creates new instance of the EdgeClusterPodStatusResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// status: Mandatory. Contains information about the edge cluster pod status
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterPodStatusResolver(
	ctx context.Context,
	status *edgeclusterGrpcContract.EdgeClusterPodStatus) (edgecluster.EdgeClusterPodStatusResolverContract, error) {
	return queryedgecluster.NewEdgeClusterPodStatusResolver(
		ctx,
		creator.logger,
		creator,
		status)
}

// NewEdgeClusterPodSpecResolver creates new instance of the EdgeClusterPodSpecResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// spec: Mandatory. Contains information about the edge cluster pod specification
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterPodSpecResolver(
	ctx context.Context,
	spec *edgeclusterGrpcContract.EdgeClusterPodSpec) (edgecluster.EdgeClusterPodSpecResolverContract, error) {
	return queryedgecluster.NewEdgeClusterPodSpecResolver(
		ctx,
		creator.logger,
		spec)
}

// NewEdgeClusterPodConditionResolver creates new instance of the EdgeClusterPodConditionResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// condition: Mandatory. Contains condition information for a pod.
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterPodConditionResolver(
	ctx context.Context,
	condition *edgeclusterGrpcContract.EdgeClusterPodCondition) (edgecluster.EdgeClusterPodConditionResolverContract, error) {
	return queryedgecluster.NewEdgeClusterPodConditionResolver(
		ctx,
		creator.logger,
		condition)
}
