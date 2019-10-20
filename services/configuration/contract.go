// Package configuration implements configuration service required by the api-gateway service
package configuration

// ConfigurationContract declares the service that provides configuration required by different Tenat modules
type ConfigurationContract interface {
	// GetHttpsHost retrieves HTTPS host name
	// Returns the HTTPS host name or error if something goes wrong
	GetHttpsHost() (string, error)

	// GetHttpsPort retrieves HTTPS port number
	// Returns the HTTPS port number or error if something goes wrong
	GetHttpsPort() (int, error)

	// GetTenantServiceAddress retrieves tenant service full gRPC address and returns it.
	// The address will be used to dial the gRPC client to connect to the tenant service.
	// Returns the tenant service address or error if something goes wrong
	GetTenantServiceAddress() (string, error)
}
