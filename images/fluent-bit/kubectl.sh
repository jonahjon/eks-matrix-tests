#!/usr/bin/env bash
set -eo pipefail

DISABLE_MD_LINTING=1
DISABLE_MD_LINK_CHECK=1
export GO111MODULE=on

#in the prowjob  (1_14 | 1_15 | 1_16)
export KUBECONFIG=/workspace/"$@"_cluster.config

kubectl get pods

echo "******************************************************"
echo "installing kubectl manifest for Grafana"
echo "******************************************************"

kubectl apply -f images/fluent-bit/templates/fluent-bit.yaml --wait

echo "******************************************************"
echo "checking logs"
echo "******************************************************"

kubectl logs -l app.kubernetes.io/instance=fluent-bit

echo "******************************************************"
echo "Running Tests"
echo "******************************************************"

kubectl apply -f images/fluent-bit/templates/tests/tests.yaml --wait

kubectl get pod -l app.kubernetes.io/name=fluent-bit --field-selector=status.phase!=Running

echo "******************************************************"
echo "Tests Passed ........ deleting Grafana"
echo "******************************************************"

kubectl delete -f images/fluent-bit/templates/tests/tests.yaml --wait

kubectl delete -f images/fluent-bit/templates/fluent-bit.yaml --wait

echo "******************************************************"
echo "Prowjob Finished Sucessfully"
echo "******************************************************"