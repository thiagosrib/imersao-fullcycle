FROM node:12.14.0-alpine3.11
# FROM node:18.14-alpine3.16
# FROM node:18.14-slim

WORKDIR /home/node/app

RUN apk add --no-cache bash

RUN npm install -g @nestjs/cli@7.5.6
# RUN npm install -g npm@9.4.2
# RUN npm install @node-rs/xxhash-linux-arm64-musl

USER node
