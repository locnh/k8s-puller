## k8s Images Puller
The puller periodically pulls image(s) to k8s cluster nodes to save the time of pulling images when launching new pods.

These are the Docker Hub autobuild images located [here](https://hub.docker.com/r/locnh/k8s-puller/).

[![License](https://img.shields.io/github/license/locnh/k8s-puller)](/LICENSE)
[![Docker](https://img.shields.io/docker/pulls/locnh/k8s-puller)](https://hub.docker.com/r/locnh/k8s-puller)
[![Build Status](https://travis-ci.org/locnh/k8s-puller.svg?branch=master)](https://travis-ci.org/locnh/k8s-puller)

## Parameters

| Parameter | Description | Type | Default |
|-----|-----|-----|-----|
| `puller.images` | `List` of images to be pulled | `List` | `[alpine]` |
| `puller.interval` | Time interval | `String` | `5m` |

## Usage
### Create the settings file

Create an `values.yaml` file like the following content, change the images and the interval (in minutes):
```yaml
puller:
  images:
    - alpine
    - busybox
  interval: 5m
```
These settings will tell the `puller` to pull the images [alpine](https://hub.docker.com/_/alpine/) and [busybox](https://hub.docker.com/_/busybox/) for every 5 minutes, because the tags was ommitted, then `latest` by default.

### Install with Helm
#### Add helm repo
```sh
helm repo add k8s-puller https://locnh.github.io/k8s-puller
```

#### Update available charts
```sh
helm repo update
```

#### Install / Upgrade the chart
Install chart with `values.yaml` in previous step.
```sh
helm upgrade --install puller k8s-puller/puller -f values.yaml
```

**Note**: I use `upgrade --install` to install the chart if not installed, and upgrade the chart if the old version was installed.


### Use as Docker container
#### Parameters as ENV variables

| Variable | Description |
|-----|-----|
| `IMAGES` | `List` of images to be pulled, separated by `,` |
| `INTERVAL` | Time interval, eg: `30s`, `5m`, `1h`, ... [more](http://golang.org/pkg/time/#ParseDuration) |

#### Run a Docker container

```sh
docker run --name puller -e IMAGES=busybox,alpine -e INTERVAL=60m -v /var/run/docker.sock:/var/run/docker.sock -d locnh/k8s-puller
```

## Contribute
1. Fork me
2. Make changes
3. Create pull request
4. Grab a cup of tee and enjoy
