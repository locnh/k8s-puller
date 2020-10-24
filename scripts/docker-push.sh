#!/bin/bash
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin

if [ $TRAVIS_BRANCH != "master" ]; then
    docker tag locnh/k8s-puller locnh/k8s-puller:devel
    docker push locnh/k8s-puller:devel
else
    docker push locnh/k8s-puller