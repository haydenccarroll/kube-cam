CLUSTER_NAME=manic-monkey
CLUSTER_CONTEXT=kind-$(CLUSTER_NAME)
APP_DIR=apps
create-cluster:
	kind create cluster --name $(CLUSTER_NAME)
delete-cluster:
	kind delete cluster --name $(CLUSTER_NAME)
apply:
	kubectl apply -Rf ./k8s --context $(CLUSTER_CONTEXT)
build-all:
	@for dir in $(APP_DIR)/*; do \
		make -C $$dir --silent build || exit 1; \
	done
	@echo "Successfully built all applications"
test-e2e:
	echo "Do not have any tests implemented"