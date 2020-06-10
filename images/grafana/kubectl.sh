#!/usr/bin/env bash
set -eo pipefail

DISABLE_MD_LINTING=1
DISABLE_MD_LINK_CHECK=1
export GO111MODULE=on

#in the prowjob  /usr/local/bin/runner.sh (1_14 | 1_15 | 1_16)
export KUBECONFIG="$@"_cluster.config

kubectl get pods

echo "******************************************************"
echo "installing kubectl manifest for Grafana"
echo "******************************************************"

kubectl apply -f templates/grafana.yaml --wait

echo "******************************************************"
echo "checking logs"
echo "******************************************************"

kubectl logs -l app.kubernetes.io/instance=grafana

kubectl get pod -l app.kubernetes.io/instance=grafana --field-selector=status.phase!=Running

echo "******************************************************"
echo "Running Tests"
echo "******************************************************"

kubectl apply -f templates/tests/tests.yaml --wait

echo "******************************************************"
echo "deleting Grafana"
echo "******************************************************"

kubectl delete -f templates/grafana.yaml --wait

