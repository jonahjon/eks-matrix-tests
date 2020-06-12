.PHONY: test-e2e

test-e2e:
	cd ./go; \
		go test -timeout=0 -v -ginkgo.v ./e2e_test.go

local-test-e2e:
	cd ./go; \
		export KUBECONFIG=~/.kube/config && go test -timeout=0 -v -ginkgo.v ./e2e_test.go

reload-components:
	kubectl delete -f "./prow/cluster/components/04-hook.yaml"
	kubectl delete -f "./prow/cluster/components/05-plank.yaml"
	kubectl delete -f "./prow/cluster/components/06-sinker.yaml"
	kubectl delete -f "./prow/cluster/components/07-deck.yaml"
	kubectl create configmap config --from-file=config.yaml=prow/cluster/components/config.yaml --dry-run -o yaml | kubectl replace configmap config -f -
	kubectl apply -f "./prow/cluster/components/04-hook.yaml"
	kubectl apply -f "./prow/cluster/components/05-plank.yaml"
	kubectl apply -f "./prow/cluster/components/06-sinker.yaml"
	kubectl apply -f "./prow/cluster/components/07-deck.yaml"

update-config:
	kubectl create configmap config --from-file=config.yaml=prow/cluster/components/config.yaml --dry-run -o yaml | kubectl replace configmap config -f -

update-clusters:
	kubectl create secret generic kubeconfig --from-file=config=./prow/cluster/components/workload_clusters.yaml --dry-run -o yaml | kubectl replace secret kubeconfig -f -

update-jobs:
	go run prow/jobs/main.go --kubeconfig $$HOME/.kube/config --jobs-config-path prow/jobs/