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

// GetHttpHost retrieves HTTP host name
// Returns the HTTP host name or error if something goes wrong
func (service *envConfigurationService) GetHttpHost() (string, error) {
	return os.Getenv("HTTP_HOST"), nil
}

// GetHttpPort retrieves HTTP port number
// Returns the HTTP port number or error if something goes wrong
func (service *envConfigurationService) GetHttpPort() (int, error) {
	portNumberString := os.Getenv("HTTP_PORT")
	if strings.Trim(portNumberString, " ") == "" {
		return 0, NewUnknownError("HTTP_PORT is required")
	}

	portNumber, err := strconv.Atoi(portNumberString)
	if err != nil {
		return 0, NewUnknownErrorWithError("Failed to convert HTTPS_PORT to integer", err)
	}

	return portNumber, nil
}

// GetProjectServiceAddress retrieves project service full gRPC address and returns it.
// The address will be used to dial the gRPC client to connect to the project service.
// Returns the project service address or error if something goes wrong
func (service *envConfigurationService) GetProjectServiceAddress() (string, error) {
	address := os.Getenv("PROJECT_ADDRESS")
	if strings.Trim(address, " ") == "" {
		return "", NewUnknownError("PROJECT_ADDRESS is required")
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

// GetJwksURL retrieves the JWKS URL
// Returns the JWKS URL or error if something goes wrong
func (service *envConfigurationService) GetJwksURL() (string, error) {
	jwksURL := os.Getenv("JWKS_URL")

	if strings.Trim(jwksURL, " ") == "" {
		return "", NewUnknownError("JWKS_URL is required")
	}

	return jwksURL, nil
}
