// Package util implements different utilities required by the API Gateway service
package util

import (
	"log"
	"os"
	"os/signal"

	"github.com/decentralized-cloud/api-gateway/services/configuration"
	"github.com/decentralized-cloud/api-gateway/services/transport/https"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types"
	"go.uber.org/zap"
)

// StartService setups all dependecies required to start the API Gateway service and
// start the service
func StartService() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	defer logger.Sync()

	configurationService, resolverCreator, err := setupDependencies(logger)
	if err != nil {
		logger.Fatal("Failed to setup dependecies", zap.Error(err))
	}

	httpsTransportService, err := https.NewTransportService(
		logger,
		configurationService,
		resolverCreator)
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

func setupDependencies(logger *zap.Logger) (configuration.ConfigurationContract, types.ResolverCreatorContract, error) {
	configurationService, err := configuration.NewEnvConfigurationService()
	if err != nil {
		return nil, nil, err
	}

	tenantClientService, err := graphql.NewTenantClientService(configurationService)
	if err != nil {
		return nil, nil, err
	}

	edgeClusterClientService, err := graphql.NewEdgeClusterClientService(configurationService)
	if err != nil {
		return nil, nil, err
	}

	resolverCreator, err := graphql.NewResolverCreator(
		logger,
		tenantClientService,
		edgeClusterClientService)
	if err != nil {
		return nil, nil, err
	}

	return configurationService, resolverCreator, err
}
