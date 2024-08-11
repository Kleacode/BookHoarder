FROM node:22.6.0-bookworm-slim

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"

RUN npm i -g pnpm

RUN set -ex \
    && apt update \
    && apt install -y \
    xdg-utils

WORKDIR /app

COPY ./package.json ./
RUN pnpm install

ENV NODE_ENV=development

CMD ["pnpm", "run", "dev"]
