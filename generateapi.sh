#!/bin/bash -eu

curl -L -o openapi-generator-cli.jar 'https://search.maven.org/remotecontent?filepath=org/openapitools/openapi-generator-cli/4.3.0/openapi-generator-cli-4.3.0.jar'
java -jar openapi-generator-cli.jar generate -i openapi.yaml -g go \
    -o pkg/instana/openapi \
    --skip-validate-spec
gofmt -s -w pkg/instana/openapi
#swagger generate client -A instana -f https://instana.github.io/openapi/openapi.yaml --skip-validation

