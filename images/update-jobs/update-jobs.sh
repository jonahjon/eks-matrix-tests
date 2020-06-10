#!/usr/bin/env bash
set -eo pipefail
set -o errexit

DISABLE_MD_LINTING=1
DISABLE_MD_LINK_CHECK=1
export GO111MODULE=off
export GOPATH="/workspace"

go get -u github.com/golang/dep/cmd/dep

#in the prowjob  /usr/local/bin/runner.sh (1_14 | 1_15 | 1_16)
export KUBECONFIG=/workspace/"$@".config
export PULL_PULL_SHA=$PULL_PULL_SHA

echo "******************************************************"
echo "DryRun of updating job configs for github PR hash $PULL_PULL_SHA"
echo "******************************************************"

git clone https://github.com/jonahjon/eks-matrix-tests $GOPATH/src/eks-matrix-tests

cd $GOPATH/src/eks-matrix-tests/prow/jobs

git checkout $PULL_PULL_SHA

echo "******************************************************"
echo "Updating Go Dep and running job update"
echo "******************************************************"

dep ensure -v

go run main.go --kubeconfig $KUBECONFIG --jobs-config-path .

echo "******************************************************"
echo "Updated jobs"
echo "******************************************************"

exit 0
