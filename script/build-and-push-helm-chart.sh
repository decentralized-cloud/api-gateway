#!/usr/bin/env sh

set -e
set -x

if [ $# -eq 0 ]; then
	current_directory=$(dirname "$0")
else
	current_directory="$1"
fi

cd "$current_directory"/..

docker build --build-arg VERSION=$VERSION --build-arg GITHUB_ACCESS_CREDENTIAL=$GITHUB_ACCESS_CREDENTIAL -f docker/Dockerfile.buildAndPushHelmChart .

