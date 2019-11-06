// Package configuration implements configuration service required by the api-gateway service
package configuration

import (
	"os"
	"strconv"
	"strings"
)

type envConfigurationService struct {
}

// NewEnvConfigurationService creates new instance of the EnvConfigurationService, setting up all dependencies and returns the instance
// Returns the new service or error if something goes wrong
func NewEnvConfigurationService() (ConfigurationContract, error) {
	return &envConfigurationService{}, nil
}

// GetHttpsHost retrieves HTTPS host name
// Returns the HTTPS host name or error if something goes wrong
func (service *envConfigurationService) GetHttpsHost() (string, error) {
	return os.Getenv("HTTPS_HOST"), nil
}

// GetHttpsPort retrieves HTTPS port number
// Returns the HTTPS port number or error if something goes wrong
func (service *envConfigurationService) GetHttpsPort() (int, error) {
	portNumberString := os.Getenv("HTTPS_PORT")
	if strings.Trim(portNumberString, " ") == "" {
		return 0, NewUnknownError("HTTPS_PORT is required")
	}

	portNumber, err := strconv.Atoi(portNumberString)
	if err != nil {
		return 0, NewUnknownErrorWithError("Failed to convert HTTPS_PORT to integer", err)
	}

	return portNumber, nil
}

// GetTenantServiceAddress retrieves tenant service full gRPC address and returns it.
// The address will be used to dial the gRPC client to connect to the tenant service.
// Returns the tenant service address or error if something goes wrong
func (service *envConfigurationService) GetTenantServiceAddress() (string, error) {
	address := os.Getenv("TENANT_ADDRESS")
	if strings.Trim(address, " ") == "" {
		return "", NewUnknownError("TENANT_ADDRESS is required")
	}

	return address, nil
}

// GetEdgeClusterServiceAddress retrieves edge cluster service full gRPC address and returns it.
// The address will be used to dial the gRPC client to connect to the edge cluster service.
// Returns the edge cluster service address or error if something goes wrong
func (service *envConfigurationService) GetEdgeClusterServiceAddress() (string, error) {
	address := os.Getenv("EDGE_CLUSTER_ADDRESS")
	if strings.Trim(address, " ") == "" {
		return "", NewUnknownError("EDGE_CLUSTER_ADDRESS is required")
	}

	return address, nil
}
