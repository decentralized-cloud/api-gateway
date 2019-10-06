#!/usr/bin/env bash

set -e

# making sure we change the directory to where the current script resides
current_directory=$(dirname "$0")
cd "$current_directory"

build_source_branch="$1"

if [[ "$build_source_branch" == "refs/heads/master" ]]; then
    # building repository name triggered by a build for the master branch
    version="latest"
elif [[ "$build_source_branch" == refs/pull* ]]; then
    # building repository name triggered by Pull Request
    build_source_branch="${build_source_branch//refs\/pull\//}"
    build_source_branch="${build_source_branch//\/merge/}"
    build_source_branch="${build_source_branch//\//-}"
    version=merge-$build_source_branch
elif [[ "$build_source_branch" == refs/tags* ]]; then
    # building repository name triggered by tagging
    build_source_branch="${build_source_branch//refs\/tags\//}"
    build_source_branch="${build_source_branch//\//-}"
    version=$build_source_branch
else
    # building repository name triggered by a build for a working branch
    build_source_branch="${build_source_branch//refs\/heads\//}"
    build_source_branch="${build_source_branch//\//-}"
    version=$build_source_branch
fi

azure_command_to_set_variable="##vso[task.setvariable variable=VERSION]"
azure_command_with_value="$azure_command_to_set_variable$version"
echo "$azure_command_with_value"

