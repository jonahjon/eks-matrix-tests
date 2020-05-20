#!/bin/bash

echo $KUBECONFIG

export KUBECONFIG=$KUBECONFIG && ../../go test -timeout=0 -v -ginkgo.v ./e2e_test.go
