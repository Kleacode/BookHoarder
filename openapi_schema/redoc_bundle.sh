#!/bin/bash

SCRIPT_DIR=$(cd "$(dirname "$0")" && pwd)

docker run --rm \
    -v $SCRIPT_DIR:/spec \
    -w /spec \
    redocly/cli \
    bundle ./openapi.yaml -o ./gen/openapi.yaml