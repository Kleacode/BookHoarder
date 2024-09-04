FROM postgres:16.4-bookworm

COPY ./init/init-user-db.sh /docker-entrypoint-initdb.d/

RUN apt-get update && apt-get install -y curl

RUN curl -fsSL \
    https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
    sh

WORKDIR /app

COPY ./migrations /app/migrations/
COPY ./migration.sh /app/