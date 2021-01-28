// Package configuration implements configuration service required by the api-gateway service
package configuration

// ConfigurationContract declares the service that provides configuration required by different Tenat modules
type ConfigurationContract interface {
	// GetHttpHost retrieves HTTP host name
	// Returns the HTTP host name or error if something goes wrong
	GetHttpHost() (string, error)

	// GetHttpPort retrieves HTTP port number
	// Returns the HTTP port number or error if something goes wrong
	GetHttpPort() (int, error)

	// GetTenantServiceAddress retrieves tenant service full gRPC address and returns it.
	// The address will be used to dial the gRPC client to connect to the tenant service.
	// Returns the tenant service address or error if something goes wrong
	GetTenantServiceAddress() (string, error)

	// GetEdgeClusterServiceAddress retrieves edge cluster service full gRPC address and returns it.
	// The address will be used to dial the gRPC client to connect to the edge cluster service.
	// Returns the edge cluster service address or error if something goes wrong
	GetEdgeClusterServiceAddress() (string, error)
}
