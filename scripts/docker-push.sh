#!/bin/bash
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker tag locnh/k8s-puller:devel locnh/k8s-puller:latest
docker tag locnh/k8s-puller:devel locnh/k8s-puller:$TRAVIS_TAG
docker push locnh/k8s-puller:latest
docker push locnh/k8s-puller:$TRAVIS_TAG
