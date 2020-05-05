.PHONY: test-e2e

test-e2e:
	cd ./go; \
		go test -timeout=0 -v -ginkgo.v ./e2e_test.go

local-test-e2e:
	cd ./go; \
		export KUBECONFIG=~/.kube/config && go test -timeout=0 -v -ginkgo.v ./e2e_test.go

#testing
