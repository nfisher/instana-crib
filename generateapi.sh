#!/bin/bash -eu

curl -L -o openapi-generator-cli.jar 'https://search.maven.org/remotecontent?filepath=org/openapitools/openapi-generator-cli/4.2.3/openapi-generator-cli-4.2.3.jar'
java -jar openapi-generator-cli.jar generate -i https://instana.github.io/openapi/openapi.yaml -g go \
    -o pkg/instana/openapi \
    --skip-validate-spec
gofmt -s -w pkg/instana/openapi

