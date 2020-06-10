.PHONY: test-e2e

test-e2e:
	cd ./go; \
		go test -timeout=0 -v -ginkgo.v ./e2e_test.go

local-test-e2e:
	cd ./go; \
		export KUBECONFIG=~/.kube/config && go test -timeout=0 -v -ginkgo.v ./e2e_test.go

update-config:
	kubectl create configmap config --from-file=config.yaml=prow/cluster/components/config.yaml --dry-run -o yaml | kubectl replace configmap config -f -

update-jobs:
	go run prow/jobs/main.go --kubeconfig $$HOME/.kube/config --jobs-config-path prow/jobs/