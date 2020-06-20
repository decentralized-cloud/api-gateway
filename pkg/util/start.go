// Package util implements different utilities required by the API Gateway service
package util

import (
	"log"
	"os"
	"os/signal"

	"github.com/decentralized-cloud/api-gateway/services/configuration"
	"github.com/decentralized-cloud/api-gateway/services/endpoint"
	"github.com/decentralized-cloud/api-gateway/services/graphql"
	"github.com/decentralized-cloud/api-gateway/services/transport/https"
	"github.com/micro-business/gokit-core/middleware"
	"go.uber.org/zap"
)

var configurationService configuration.ConfigurationContract
var endpointCreatorService endpoint.EndpointCreatorContract
var middlewareProviderService middleware.MiddlewareProviderContract

// StartService setups all dependecies required to start the API Gateway service and
// start the service
func StartService() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = logger.Sync()
	}()

	err = setupDependencies(logger)
	if err != nil {
		logger.Fatal("Failed to setup dependecies", zap.Error(err))
	}

	httpsTransportService, err := https.NewTransportService(
		logger,
		configurationService,
		endpointCreatorService,
		middlewareProviderService)
	if err != nil {
		logger.Fatal("Failed to create GraphQL transport service", zap.Error(err))
	}

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan struct{})
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		if serviceErr := httpsTransportService.Start(); serviceErr != nil {
			logger.Fatal("Failed to start HTTPS transport service", zap.Error(serviceErr))
		}
	}()

	go func() {
		<-signalChan
		logger.Info("Received an interrupt, stopping services...")

		if err := httpsTransportService.Stop(); err != nil {
			logger.Error("Failed to stop HTTPS transport service", zap.Error(err))
		}

		close(cleanupDone)
	}()
	<-cleanupDone
}

func setupDependencies(logger *zap.Logger) (err error) {
	if configurationService, err = configuration.NewEnvConfigurationService(); err != nil {
		return
	}

	if middlewareProviderService, err = middleware.NewMiddlewareProviderService(logger, true, ""); err != nil {
		return
	}

	tenantClientService, err := graphql.NewTenantClientService(configurationService)
	if err != nil {
		return
	}

	edgeClusterClientService, err := graphql.NewEdgeClusterClientService(configurationService)
	if err != nil {
		return
	}

	resolverCreator, err := graphql.NewResolverCreator(
		logger,
		tenantClientService,
		edgeClusterClientService)
	if err != nil {
		return
	}

	if endpointCreatorService, err = endpoint.NewEndpointCreatorService(resolverCreator); err != nil {
		return
	}

	return
}
