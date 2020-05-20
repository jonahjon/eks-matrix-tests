.PHONY: test

test:
	cd ../../go; \
		go test KUBECONFIG=$$KUBECONFIG -timeout=0 -v -ginkgo.v ./e2e_test.go