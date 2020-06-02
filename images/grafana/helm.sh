#!/usr/bin/env bash
set -eo pipefail

DISABLE_MD_LINTING=1
DISABLE_MD_LINK_CHECK=1
export GO111MODULE=on

#in the prowjob  /usr/local/bin/runner.sh (1_14 | 1_15 | 1_16)
export KUBECONFIG="$@"_cluster.config

kubectl get pods

echo "******************************************************"
echo "adding helm chart repo"
echo "******************************************************"

helm repo add stable https://kubernetes-charts.storage.googleapis.com/

helm repo update 

helm version 

echo "******************************************************"
echo "installing helm chart"
echo "******************************************************"

helm install grafana stable/grafana --wait

echo "******************************************************"
echo "checking logs"
echo "******************************************************"

kubectl logs -l app.kubernetes.io/instance=grafana

kubectl get pod -l app.kubernetes.io/instance=grafana --field-selector=status.phase!=Running

echo "******************************************************"
echo "deleting helm charts"
echo "******************************************************"

helm delete grafana --timeout 10m0s

