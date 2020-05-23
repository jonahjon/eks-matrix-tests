.PHONY: test

test:
	cd ../../go; \
		export KUBECONFIG=$$KUBECONFIG && go test -timeout=0 -v -ginkgo.v ./e2e_test.go