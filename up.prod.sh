#!/bin/sh

docker-compose \
    -f docker-compose.yml \
    -f docker-compose.production.yml \
    up -d \
    --remove-orphans \
    --force-recreate \
    --build