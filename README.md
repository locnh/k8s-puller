## k8s Images Puller
The puller periodically pulls image(s) to k8s cluster nodes to save the time of pulling images when launching new pods.

These are the Docker Hub autobuild images located [here](https://hub.docker.com/r/locnh/k8s-puller/).

[![License](https://img.shields.io/github/license/locnh/k8s-puller)](/LICENSE)
[![Build Status](https://travis-ci.org/locnh/k8s-puller.svg?branch=master)](https://travis-ci.org/locnh/k8s-puller)
[![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/locnh/k8s-puller?sort=semver)](/Dockerfile)
[![Docker Image Version (latest semver)](https://img.shields.io/docker/v/locnh/k8s-puller?sort=semver)](/Dockerfile)
[![Docker](https://img.shields.io/docker/pulls/locnh/k8s-puller)](https://hub.docker.com/r/locnh/k8s-puller)
[![codecov](https://codecov.io/gh/locnh/k8s-puller/branch/master/graph/badge.svg?token=22M1LNHEEM)](https://codecov.io/gh/locnh/k8s-puller)

## Usage
### Helm

**Note**: Helm chart has been moved to [HowDevOps/helm-charts](https://github.com/HowDevOps/helm-charts/tree/main/charts/k8s-puller) repository.


### Use as Docker container
#### Parameters as ENV variables

| Variable | Description | Mandatory | Default |
|-----|-----|-----|-----|
| `IMAGES` | `List` of images to be pulled, separated by `,` | Yes | `null` |
| `INTERVAL` | Time `interval`, eg: `30s`, `5m`, `1h`, ... [more](http://golang.org/pkg/time/#ParseDuration) | No | `60m` |
| `JSONLOG` | Toggle for `JSON` log format | No | `false` |
| `DOCKER_USERNAME` | `username` to login to docker registry | No | `""` |
| `DOCKER_PASSWORD` | `password` to login to docker registry | No | `""` |
| `DOCKER_SERVER` | [server](https://docs.docker.com/engine/reference/commandline/login/#login-to-a-self-hosted-registry) to login to docker registry | No | `""` |

#### Run a Docker container

```sh
docker run --name puller -e IMAGES=busybox,alpine -e INTERVAL=60m -v /var/run/docker.sock:/var/run/docker.sock -d locnh/k8s-puller
```

## Contribute
1. Fork me
2. Make changes
3. Create pull request
4. Grab a cup of tee and enjoy
