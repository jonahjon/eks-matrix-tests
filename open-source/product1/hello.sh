#!/bin/bash

echo $KUBECONFIG

export KUBECONFIG=$KUBECONFIG && cd ../../go && go test -timeout=0 -v -ginkgo.v ./e2e_test.go
