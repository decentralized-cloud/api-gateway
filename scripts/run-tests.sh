#!/usr/bin/env sh

set -e
set -x

cleanup() {
	docker-compose  -f docker/docker-compose-test.yml down
}

trap 'cleanup' EXIT
if [ $# -eq 0 ]; then
	current_directory=$(dirname "$0")
else
	current_directory="$1"
fi

cd "$current_directory"/..

docker-compose  -f docker/docker-compose-test.yml build
docker-compose  -f docker/docker-compose-test.yml run --rm -e COVERALLS_SERVICE_NAME=$COVERALLS_SERVICE_NAME -e COVERALLS_REPO_TOKEN=$COVERALLS_REPO_TOKEN tests

