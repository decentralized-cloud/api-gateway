// packae edgecluster implements used edge cluster related types in the GraphQL transport layer
package edgecluster

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types/relay"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	"github.com/graph-gophers/graphql-go"
)

type QueryResolverCreatorContract interface {
	// NewEdgeClusterResolver creates new EdgeClusterResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// edgeClusterID: Mandatory. The edge cluster unique identifier
	// edgeClusterDetail: Optional. The edge cluster details, if provided, the value be used instead of contacting  the edge cluster service
	// Returns the EdgeClusterResolverContract or error if something goes wrong
	NewEdgeClusterResolver(
		ctx context.Context,
		edgeClusterID string,
		edgeClusterDetail *EdgeClusterDetail) (EdgeClusterResolverContract, error)

	// NewEdgeClusterObjectMetadataResolverContract creates new instance of the EdgeClusterObjectMetadataResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// metadata: Mandatory. Contains the object metadata.
	// Returns the new instance or error if something goes wrong
	NewEdgeClusterObjectMetadataResolverContract(
		ctx context.Context,
		metadata *edgeclusterGrpcContract.EdgeClusterObjectMetadata) (EdgeClusterObjectMetadataResolverContract, error)

	// NewEdgeClusterTypeEdgeResolver creates new EdgeClusterTypeEdgeResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// edgeClusterID: Mandatory. The edge cluster unique identifier
	// cursor: Mandatory. The cursor
	// edgeClusterDetail: Optional. The edge cluster details, if provided, the value be used instead of contacting  the edge cluster service
	// Returns the EdgeClusterTypeEdgeResolverContract or error if something goes wrong
	NewEdgeClusterTypeEdgeResolver(
		ctx context.Context,
		edgeClusterID string,
		cursor string,
		edgeClusterDetail *EdgeClusterDetail) (EdgeClusterTypeEdgeResolverContract, error)

	// NewEdgeClusterTypeConnectionResolver creates new EdgeClusterTypeConnectionResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// edgeClusters: Mandatory. Reference the list of edge clusters
	// hasPreviousPage: Mandatory. Indicates whether more edges exist prior to the set defined by the clients arguments
	// hasNextPage: Mandatory. Indicates whether more edges exist following the set defined by the clients arguments
	// totalCount: Mandatory. The total count of matched edge clusters
	// Returns the EdgeClusterTypeConnectionResolverContract or error if something goes wrong
	NewEdgeClusterTypeConnectionResolver(
		ctx context.Context,
		edgeClusters []*edgeclusterGrpcContract.EdgeClusterWithCursor,
		hasPreviousPage bool,
		hasNextPage bool,
		totalCount int32) (EdgeClusterTypeConnectionResolverContract, error)

	// NewEdgeClusterProjectResolver creates new EdgeClusterTenatnResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// projectID: Mandatory. The project unique identifier
	// Returns the EdgeClusterTenatnResolverContract or error if something goes wrong
	NewEdgeClusterProjectResolver(
		ctx context.Context,
		projectID string) (EdgeClusterProjectResolverContract, error)

	// NewEdgeClusterProvisionDetailResolver creates new EdgeClusterProvisionDetailResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// provisionDetail: Optional. The edge cluster provisioning details
	// Returns the EdgeClusterProvisionDetailResolverContract or error if something goes wrong
	NewEdgeClusterProvisionDetailResolver(
		ctx context.Context,
		provisionDetail *edgeclusterGrpcContract.EdgeClusterProvisionDetail) (EdgeClusterProvisionDetailResolverContract, error)

	// NewIngressResolver creates new instance of the ingressResolver, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// ingress: Mandatory. The ingress details
	// Returns the new instance or error if something goes wrong
	NewIngressResolver(
		ctx context.Context,
		ingress *edgeclusterGrpcContract.Ingress) (IngressResolverContract, error)

	// NewPortResolver creates new instance of the PortResolver, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// logger: Mandatory. Reference to the logger service
	// port: Mandatory. The PostStatus details
	// Returns the new instance or error if something goes wrong
	NewPortResolver(
		ctx context.Context,
		port *edgeclusterGrpcContract.Port) (PortResolverContract, error)

	// NewEdgeClusterNodeResolver creates new instance of the NewEdgeClusterNodeResolver, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// logger: Mandatory. Reference to the logger service
	// node: Mandatory. Contains information about the edge cluster node
	// Returns the new instance or error if something goes wrong
	NewEdgeClusterNodeResolver(
		ctx context.Context,
		node *edgeclusterGrpcContract.EdgeClusterNode) (EdgeClusterNodeResolverContract, error)

	// NewEdgeClusterNodeStatusResolver creates new instance of the EdgeClusterNodeStatusResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// logger: Mandatory. Reference to the logger service
	// status: Mandatory. Contains information about the edge cluster node status
	// Returns the new instance or error if something goes wrong
	NewEdgeClusterNodeStatusResolver(
		ctx context.Context,
		status *edgeclusterGrpcContract.EdgeClusterNodeStatus) (EdgeClusterNodeStatusResolverContract, error)

	// NewEdgeClusterNodeConditionResolver creates new instance of the EdgeClusterNodeConditionResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// condition: Mandatory. Contains condition information for a node.
	// Returns the new instance or error if something goes wrong
	NewEdgeClusterNodeConditionResolver(
		ctx context.Context,
		condition *edgeclusterGrpcContract.EdgeClusterNodeCondition) (EdgeClusterNodeConditionResolverContract, error)

	// NewEdgeClusterNodeAddressResolverContract creates new instance of the EdgeClusterNodeAddressResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// nodeAddresss: Mandatory. Contains information for the edge cluster node's address.
	// Returns the new instance or error if something goes wrong
	NewEdgeClusterNodeAddressResolverContract(
		ctx context.Context,
		nodeAddresss *edgeclusterGrpcContract.EdgeClusterNodeAddress) (EdgeClusterNodeAddressResolverContract, error)

	// NewEdgeClusterNodeSystemInfoResolverContract creates new instance of the EdgeClusterNodeSystemInfoResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// nodeInfo: Mandatory. Contains a set of ids/uuids to uniquely identify the node.
	// Returns the new instance or error if something goes wrong
	NewEdgeClusterNodeSystemInfoResolverContract(
		ctx context.Context,
		nodeInfo *edgeclusterGrpcContract.EdgeClusterNodeSystemInfo) (EdgeClusterNodeSystemInfoResolverContract, error)

	// NewEdgeClusterPodResolver creates new instance of the NewEdgeClusterPodResolver, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// logger: Mandatory. Reference to the logger service
	// pod: Mandatory. Contains information about the edge cluster pod
	// Returns the new instance or error if something goes wrong
	NewEdgeClusterPodResolver(
		ctx context.Context,
		pod *edgeclusterGrpcContract.EdgeClusterPod) (EdgeClusterPodResolverContract, error)

	// NewEdgeClusterPodStatusResolver creates new instance of the EdgeClusterPodStatusResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// logger: Mandatory. Reference to the logger service
	// status: Mandatory. Contains information about the edge cluster pod status
	// Returns the new instance or error if something goes wrong
	NewEdgeClusterPodStatusResolver(
		ctx context.Context,
		status *edgeclusterGrpcContract.EdgeClusterPodStatus) (EdgeClusterPodStatusResolverContract, error)

	// NewEdgeClusterPodSpecResolver creates new instance of the EdgeClusterPodSpecResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// logger: Mandatory. Reference to the logger service
	// spec: Mandatory. Contains information about the edge cluster pod specification
	// Returns the new instance or error if something goes wrong
	NewEdgeClusterPodSpecResolver(
		ctx context.Context,
		spec *edgeclusterGrpcContract.EdgeClusterPodSpec) (EdgeClusterPodSpecResolverContract, error)

	// NewEdgeClusterPodConditionResolver creates new instance of the EdgeClusterPodConditionResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// condition: Mandatory. Contains condition information for a pod.
	// Returns the new instance or error if something goes wrong
	NewEdgeClusterPodConditionResolver(
		ctx context.Context,
		condition *edgeclusterGrpcContract.EdgeClusterPodCondition) (EdgeClusterPodConditionResolverContract, error)
}

type EdgeClusterDetail struct {
	EdgeCluster     *edgeclusterGrpcContract.EdgeCluster
	ProvisionDetail *edgeclusterGrpcContract.EdgeClusterProvisionDetail
}

type EdgeClusterClusterEdgeClusterInputArgument struct {
	EdgeClusterID graphql.ID
}

type EdgeClusterEdgeClustersInputArgument struct {
	relay.ConnectionArgument
	EdgeClusterIDs *[]graphql.ID
	SortOption     *string
}
