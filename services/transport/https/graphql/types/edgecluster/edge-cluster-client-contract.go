// packae edgecluster implements used edge cluster related types in the GraphQL transport layer
package edgecluster

import (
	edgeClusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	"google.golang.org/grpc"
)

// EdgeClusterClientContract wraps the edge cluser gRPC client to make it easy for testing
type EdgeClusterClientContract interface {
	// CreateClient creats a new edge cluster gRPC client and returns the connection
	// and the client to the caller.
	// Returns connection and the edge cluster gRPC client or error if something goes wrong.
	CreateClient() (*grpc.ClientConn, edgeClusterGrpcContract.EdgeClusterServiceClient, error)
}
