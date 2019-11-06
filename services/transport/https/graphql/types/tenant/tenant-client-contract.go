// packae tenant implements used tenant related types in the GraphQL transport layer
package tenant

import (
	tenantGrpcContract "github.com/decentralized-cloud/tenant/contract/grpc/go"
	"google.golang.org/grpc"
)

// TenantClientContract wraps the tennat gRPC client to make it easy for testing
type TenantClientContract interface {
	// CreateClient creats a new tenant gRPC client and returns the connection
	// and the client to the caller.
	// Returns connection and the tenant gRPC client or error if something goes wrong.
	CreateClient() (*grpc.ClientConn, tenantGrpcContract.TenantServiceClient, error)
}
