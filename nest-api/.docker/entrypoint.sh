#!/bin/bash

if [ ! -f ".env" ]; then
  cp .env.example .env
fi

mkdir ~/.npm-global
npm config set prefix '~/.npm-global'
export PATH=~/.npm-global/bin:$PATH

echo "installing"
npm ci --legacy-peer-dependencies

echo "running"
npm run start:dev