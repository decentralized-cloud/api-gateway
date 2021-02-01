#!/usr/bin/env sh

set -e
set -x

cleanup() {
	docker rm extract-api-gateway-contract-graphql-builder
}

trap 'cleanup' EXIT

if [ $# -eq 0 ]; then
	current_directory=$(dirname "$0")
else
	current_directory="$1"
fi

cd "$current_directory"/..

docker build -f docker/Dockerfile.buildGraphQLContract -t api-gateway-contract-graphql-builder .
docker create --name extract-api-gateway-contract-graphql-builder api-gateway-contract-graphql-builder
docker cp extract-api-gateway-contract-graphql-builder:/src/contract/graphql/schema/schema.graphql ./contract/graphql/schema
