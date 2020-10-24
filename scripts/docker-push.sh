#!/bin/bash
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin

if [ $TRAVIS_BRANCH != "master" ]; then
    docker tag locnh/k8s-puller locnh/k8s-puller:devel
    docker push locnh/k8s-puller:devel
elif [ -z $TRAVIS_TAG ]; then
    docker tag locnh/k8s-puller locnh/k8s-puller:$TRAVIS_TAG
    docker push locnh/k8s-puller:$TRAVIS_TAG
else
    docker push locnh/k8s-puller
fi