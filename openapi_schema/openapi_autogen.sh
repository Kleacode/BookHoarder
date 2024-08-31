#!/bin/bash
SCRIPT_DIR=$(cd "$(dirname "$0")" && pwd)
BACKEND_CONTAINER_NAME=bookhoarder-back
FRONTEND_CONTAINER_NAME=bookhoarder-front

echo 1. Bundle openapi with redoc

${SCRIPT_DIR}/redoc_bundle.sh

echo 2. Generate backend files with oapi-gen

docker exec -w /app/tools ${BACKEND_CONTAINER_NAME} go generate

echo 3. Generate frontend files with openapi-typescript

docker exec -w /app ${FRONTEND_CONTAINER_NAME} npm run generate:api

echo Finish!