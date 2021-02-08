// packae project implements used project related types in the GraphQL transport layer
package project

import (
	projectGrpcContract "github.com/decentralized-cloud/project/contract/grpc/go"
	"google.golang.org/grpc"
)

// ProjectClientContract wraps the tennat gRPC client to make it easy for testing
type ProjectClientContract interface {
	// CreateClient creats a new project gRPC client and returns the connection
	// and the client to the caller.
	// Returns connection and the project gRPC client or error if something goes wrong.
	CreateClient() (*grpc.ClientConn, projectGrpcContract.ProjectServiceClient, error)
}
