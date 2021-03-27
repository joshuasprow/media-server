#!/bin/bash

# restarts any outdated containers with the latest images
docker-compose pull && \
    ./up.prod.sh