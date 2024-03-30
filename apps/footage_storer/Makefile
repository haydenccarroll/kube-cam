include ../../Makefile

build:
	docker build -t footage_storer .
	kind load docker-image footage_storer:latest --name $(CLUSTER_NAME)
