## k8s Images Puller
The puller periodically pulls image(s) to k8s cluster nodes to save the time of pulling images when launching new pods.

These are the Docker Hub autobuild images located [here](https://hub.docker.com/r/locnh/k8s-puller/).

## Parameters

| Parameter | Description | Type | Default |
|-----|-----|-----|-----|-----|
| `puller.images` | `List` of images to be pulled | `List` | `[alpine]` |
| `puller.interval` | Time interval in minutes | `Int` | `60` |

## Usage
### Create the settings file

Create an `values.yaml` file like the following content, change the images and the interval (in minutes):
```yaml
puller:
  images:
    - alpine
    - busybox
  interval: 5
```
These settings will tell the `puller` to pull the images [alpine](https://hub.docker.com/_/alpine/) and [busybox](https://hub.docker.com/_/busybox/) for every 5 minutes, because the tags was ommitted, then `latest` by default.

### Install with Helm
#### Add helm repo
```sh
helm repo add locnh https://github.com/locnh/k8s-puller
```

#### Install / Upgrade the chart
Install chart with `values.yaml` in previous step.
```sh
helm upgrade --install puller locnh/puller -f values.yaml
```

Note: I used `upgrade --install` to install the chart if not installed, and upgrade the chart if the old version was installed.


### Use as Docker container
#### Parameters as ENV variables

| Variable | Description |
|-----|-----|
| `IMAGES` | `List` of images to be pulled, separated by `,` |
| `INTERVAL` | Time interval in minutes |

#### Run a Docker container

```sh
docker run --name puller -e IMAGES=busybox,alpine -e INTERVAL=60 -v /var/run/docker.sock:/var/run/docker.sock -d locnh/k8s-puller
```

## Contribute
1. Fork me
2. Make changes
3. Create pull request
4. Grab a cup of coffee