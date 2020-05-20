#!/bin/bash

#!/bin/bash

echo $KUBECONFIG

cat $KUBECONFIG

cd ../../go

go test -timeout=0 -v -ginkgo.v ./e2e_test.go
