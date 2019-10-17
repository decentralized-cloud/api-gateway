// Package configuration implements configuration service required by the tenant service
package configuration

// ConfigurationContract declares the service that provides configuration required by different Tenat modules
type ConfigurationContract interface {
	// GetHost retrieves host name
	// Returns the host name or error if something goes wrong
	GetHost() (string, error)

	// GetPort retrieves port number
	// Returns the port number or error if something goes wrong
	GetPort() (int, error)

	// GetTenantServiceAddress retrieves tenant service full gRPC address and returns it.
	// The address will be used to dial the gRPC client to connect to the tenant service.
	// Returns the tenant service address or error if something goes wrong
	GetTenantServiceAddress() (string, error)
}
