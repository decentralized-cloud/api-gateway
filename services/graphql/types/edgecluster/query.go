// packae edgecluster implements used edge cluster related types in the GraphQL transport layer
package edgecluster

import (
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/relay"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	"github.com/graph-gophers/graphql-go"
)

type QueryResolverCreatorContract interface {
	CommonResolverCreatorContract
	EdgeClusterResolverCreatorContract
	EdgeClusterNodeResolverCreatorContract
	EdgeClusterPodResolverCreatorContract
	EdgeClusterServiceResolverCreatorContract
}

type EdgeClusterDetail struct {
	EdgeCluster      *edgeclusterGrpcContract.EdgeCluster
	ProvisionDetails *edgeclusterGrpcContract.ProvisionDetail
}

type EdgeClusterClusterEdgeClusterInputArgument struct {
	EdgeClusterID graphql.ID
}

type EdgeClusterEdgeClustersInputArgument struct {
	relay.ConnectionArgument
	EdgeClusterIDs *[]graphql.ID
	SortOption     *string
}
