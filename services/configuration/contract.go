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
}
