# kube-cam
A simulated security camera system made with kubernetes

## Description
KubeCam is a simulated security camera system made with kubernetes. Each camera within the system will be it's own kubernetes deployment. These camera deployments will simply push raw video data (that is pre-mounted on the container) to a `camera-uuidhere-queue` topic. A different "FootageAlterer" deployment will read from these N different topics and perform some basic video processing on them. Examples include adding a water mark, changing the grayscale, upping the contrast, etc. After the footage gets altered with the basic video processing techniques, it will then get archived. This archive will also live on the kubernetes cluster.

## Architecture
CameraDeployment - sends fake camera data to a queue

FootageAlterer - reads from camera deployment queue and adds a watermark ontop. puts the altered footage into a queue.

FootageStorer - reads from the FootageAlterer queue and stores it in some way

FootageQueuer - queue implementation to send footage data from one application to another

## Development
### Overview
* I want to use this project to practice writing good documentation. Documentation should be clear, and be simple for any user from most any platform to go from cloning the repo, to start contributing.
* Testing. For this project, I want 100% (or close to it) unit test coverage. I have not been in the practice of writing tests for my personal projects, and I understand it is an invaluable skill. I also some end-to-end tests that will take a video feed as input, and check for the altered footage in the database.
* I am going to use GitHub to host this project, and keep everything open source. With this project having a few different subprojects, each subproject will be a subdirectory in 1 repository (versus many).
### Tools
Here's a list of tools I think we should use
* Programming Languages - any backend language. Each microservice can be written in whatever language the "champion" of the service wants. I am writing the service I am championing in Rust.
* Git. We should be using git to version our changes
* GitHub. I'm choosing to host this on GitHub. Every person must submit a PR (which needs a review from me personally). No merging directly to the main branch. Let's use best practice here.
* `k9s` is a nice little GUI that lives in your terminal. We can use this to view deployments, config maps, etc on a kubernetes cluster.
* `kind`is a kubernetes emulator(?) that can be used to create a kubernetes cluster on your local machine for local development. We can also use this to create a kubernetes cluster for CI stuff.
* IDE - whatever. I use VSCode, Goland, or whatever other Jetbrains IDE

## Contributing
### Local Development

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
