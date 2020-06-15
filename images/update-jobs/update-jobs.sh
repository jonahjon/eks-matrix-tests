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

# Using prowjob elements:
# decorate: true 
# path_alias: github.com/jonahjon/eks-matrix-tests
# Will pull in the PR git HASH into the image via the Initupload Sidecar
cd /home/prow/go/src/github.com/jonahjon/eks-matrix-tests

echo "******************************************************"
echo "Updating Go Dep and running job update"
echo "******************************************************"

dep ensure -v

go run main.go --kubeconfig $KUBECONFIG --jobs-config-path .

echo "******************************************************"
echo "Updated-jobs image"
echo "******************************************************"

exit 0
