FROM node:22.6.0-bookworm-slim

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"

RUN npm i -g pnpm

ARG username=vscode
ARG useruid=1000
ARG usergid=${useruid}

RUN set -ex \
    && apt-get update \
    && apt-get install -y \
    sudo \
    --no-install-recommends \
    # Delete node user with uid=1000 and create vscode user with uid=1000
    && userdel -r node \
    && groupadd --gid ${usergid} ${username} \
    && useradd -s /bin/bash --uid ${useruid} --gid ${usergid} -m ${username} \
    && echo ${username} ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/${username} \
    && chmod 0440 /etc/sudoers.d/${username}

USER ${username}

WORKDIR /app

COPY ./package.json ./
RUN pnpm install

ENV NODE_ENV=development

EXPOSE 3000

CMD ["pnpm", "run", "dev"]
