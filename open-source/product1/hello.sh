#!/bin/bash

echo $KUBECONFIG

cat $KUBECONFIG

cd ../../go

ls -la

go test -timeout=0 -v -ginkgo.v ./e2e_test.go

# export KUBECONFIG=$KUBECONFIG && cd ../../go && go test -timeout=0 -v -ginkgo.v ./e2e_test.go
