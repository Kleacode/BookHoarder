#!/bin/bash
SCRIPT_DIR=$(cd "$(dirname "$0")" && pwd)

MIGRATIONS_FILE_DIR="${SCRIPT_DIR}/migrations"

goose -dir $MIGRATIONS_FILE_DIR postgres "user=$POSTGRES_USER dbname=$POSTGRES_TEST_DB password=$POSTGRES_PASSWORD sslmode=disable" up
goose -dir $MIGRATIONS_FILE_DIR postgres "user=$POSTGRES_USER dbname=$POSTGRES_DEV_DB password=$POSTGRES_PASSWORD sslmode=disable" up