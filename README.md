# kube-cam
A simulated cloud camera system made with kubernetes

## Description
KubeCam is a simulated cloud camera system made with kubernetes. The `footage-getter` deployment will send fake camera data to a RabbitMQ `footage` queue.
A `footage-storer` deployment will read from this queue and store this data on the cluster using kubernetes persistent volumes.

## Architecture
`footage-getter` - sends fake camera data to a the RabbitMQ `footage` queue.

`footage-storer` - reads from the `footage` queue and stores it on the cluster.

`RabbitMQ` - the queueing system that the above deployments use to communicate with each other. There is only 1 topic/queue called `footage`.

## Local Development
First, install these applications
- `docker`
- `kind`
- `k9s`
- `make`
- `kubectl`


Creating and using the local kubernetes cluster
1. Run `make create-cluster` to create a new kubernetes cluster from scratch
2. Run `make apply` to apply all of the kubernetes manifests defined in `k8s/`
3. Run `k9s` to view deployments, pods, configmaps, logs, etc for the cluster
