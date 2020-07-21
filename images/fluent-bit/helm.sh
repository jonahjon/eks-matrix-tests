#!/usr/bin/env bash
set -eo pipefail

DISABLE_MD_LINTING=1
DISABLE_MD_LINK_CHECK=1
export GO111MODULE=on

#in the prowjob (1_14 | 1_15 | 1_16)
export KUBECONFIG=/workspace/"$@"_cluster.config

echo "******************************************************"
echo "adding helm chart repo"
echo "******************************************************"

helm repo add stable https://kubernetes-charts.storage.googleapis.com/

helm repo update 

helm version 

echo "******************************************************"
echo "installing helm chart"
echo "******************************************************"

helm install fluent-bit stable/fluent-bit --wait

echo "******************************************************"
echo "checking logs"
echo "******************************************************"

kubectl logs -l app.kubernetes.io/instance=fluent-bit

kubectl get pod -l app.kubernetes.io/instance=fluent-bit --field-selector=status.phase!=Running

echo "******************************************************"
echo "Running Tests"
echo "******************************************************"

helm test fluent-bit

echo "******************************************************"
echo "deleting helm charts"
echo "******************************************************"

helm delete fluent-bit --timeout 10m0s

