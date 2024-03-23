# kube-cam
A simulated security camera system made with kubernetes

## Description
KubeCam is a simulated security camera system made with kubernetes. Each camera within the system will be its own kubernetes deployment. These `footage-getter` deployments will simply push raw video data (that is either pulled from the internet, or premounted onto the image) to a queue. 
A different `footage-alter` deployment will read from this queue and perform some basic video processing on it.
Examples include adding a watermark, changing the grayscale, upping the contrast, etc. After the footage gets altered 
it will be pushed to a different queue. Finally, a `footage-storer` deployment will read from this queue and store the altered footage in some way.

## Architecture
`footage-getter` - sends fake camera data to a queue

`footage-alterer` - reads from camera deployment queue and adds a watermark ontop. puts the altered footage into a queue.

`footage-storer` - reads from the `altered-feed-queue` queue and stores it on the cluster. Since this is storing data, a `StatefulSet` will need to be used. 
    The `rabbitmq` kubernetes manifest file is a good example of a StatefulSet and how to store data on the cluster persistently.

`rabbitmq` - the queueing system that the different deployments will be used to communicate with each other. It will have 2 queues:
    1. camera-feed-queue (footage getter to footage alterer)
    2. altered-feed-queue (footage alterer to footage storer)

## Development
### Overview
* Good documentation
* Good testing

## Local Development
First, install these applications
- `docker`
- `kind`
- `k9s`
- `make`
- `kubectl`

All implementation for each subsystem goes in each `apps/subdir` subdirectory. These should all have unit tests and a nice README.md. They should all have a `make build` command defined in them that will build a docker image for the cluster to use.

Creating and using the local kubernetes cluster
1. Run `make create-cluster` to create a new kubernetes cluster from scratch
2. Run `make apply` to apply all of the kubernetes manifests defined in `k8s/`
3. Run `k9s` to view deployments, pods, configmaps, logs, etc for the cluster
