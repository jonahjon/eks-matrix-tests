#!/bin/bash

kubectl create configmap plugins --from-file "./prow/cluster/components/plugins.yaml" --dry-run -o yaml | kubectl replace configmap plugins -f -
kubectl create configmap config --from-file "./prow/cluster/components/config.yaml" --dry-run -o yaml | kubectl replace configmap config -f -
kubectl create configmap branding --from-file "./prow/branding" --dry-run -o yaml | kubectl replace configmap branding -f -
kubectl create secret generic hmac-token --from-file=hmac=./prow/cluster/components/bot_hmac --dry-run -o yaml | kubectl replace secret hmac-token -f -
kubectl create secret generic oauth-token --from-file=oauth=./prow/cluster/components/bot_oauth --dry-run -o yaml | kubectl replace secret oauth-token -f -
kubectl create secret generic github-oauth-config --from-file=secret=./prow/cluster/components/deck_oauth --dry-run -o yaml | kubectl replace secret github-oauth-config -f -
