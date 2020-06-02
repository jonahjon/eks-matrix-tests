#!/usr/bin/env bash
set -eo pipefail

DISABLE_MD_LINTING=1
DISABLE_MD_LINK_CHECK=1
export GO111MODULE=on

echo "******************************************************"
echo "installing helm chart"
echo "******************************************************"

helm install grafana stable/grafana --wait

echo "******************************************************"
echo "checking logs"
echo "******************************************************"

k logs -l app.kubernetes.io/instance=grafana

k get pod -l app.kubernetes.io/instance=grafana --field-selector=status.phase!=Running

echo "******************************************************"
echo "deleting logs"
echo "******************************************************"

helm delete grafana --timeout duration 10m0s



[Service]
Environment="AWS_ACCESS_KEY_ID=AKIASMRP2L3QI5PZQPOJ"
Environment="AWS_SECRET_ACCESS_KEY=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"