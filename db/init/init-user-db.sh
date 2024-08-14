#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username ${POSTGRES_USER} --dbname ${POSTGRES_DB} <<-EOSQL
	CREATE TABLE Users(
	  id        SERIAL PRIMARY KEY,
	  name      VARCHAR(20)
	);

	INSERT INTO Users (name) VALUES ('aiueo');
EOSQL
