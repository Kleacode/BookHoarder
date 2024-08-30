#!/bin/bash
SCRIPT_DIR=$(cd "$(dirname "$0")" && pwd)
BACKEND_CONTAINER_NAME=bookhoarder-back

echo 1. Bundle openapi with redoc

${SCRIPT_DIR}/redoc_bundle.sh

echo 2. Generate backend files with oapi-gen

docker exec -w /app/tools ${BACKEND_CONTAINER_NAME} go generate

echo Finish!