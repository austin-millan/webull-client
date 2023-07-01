#!/bin/sh
# This script will codegen a go package for a server stub or client.

rm -rf openapi/

docker run \
    --rm \
    --user $(id -u):$(id -g) \
    -v ${PWD}:/local \
    openapitools/openapi-generator-cli:latest generate \
        -i /local/swagger.yml \
        -g go \
        -o /local/models
rm models/api_*.go \
    models/.travis.yml \
    models/client.go \
    models/configuration.go \
    models/git_push.sh \
    models/go.mod \
    models/go.sum \
    models/response.go \
    models/api/openapi.yaml \
    models/docs/*Api.md \
    models/test/*.go
