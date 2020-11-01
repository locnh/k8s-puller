#!/bin/sh

# Check if docker registry credentials to login
if [ ! -z $DOCKER_USERNAME ] && [ ! -z $DOCKER_PASSWORD ]; then
    if [ -z $DOCKER_SERVER ]; then
        echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
    else
        echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin $DOCKER_SERVER
    fi
fi

# Main program
/docker-puller