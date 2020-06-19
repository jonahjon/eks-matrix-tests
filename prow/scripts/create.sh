#!/bin/bash

kubectl apply -k "github.com/kubernetes-sigs/aws-ebs-csi-driver/deploy/kubernetes/overlays/stable/?ref=master"

kubectl create configmap plugins --from-file "./prow/cluster/components/plugins.yaml"
kubectl create configmap config --from-file "./prow/cluster/components/config.yaml"
kubectl create configmap job-config --from-file "./prow/jobs/config.yaml"
kubectl create configmap branding --from-file "./prow/branding"
kubectl create secret generic hmac-token --from-file "./prow/cluster/components/bot_hmac"
kubectl create secret generic oauth-token --from-file "./prow/cluster/components/bot_oauth"
kubectl create secret generic github-oauth-config --from-file=secret="./prow/cluster/components/deck_oauth"
kubectl create secret generic cookie --from-file=secret="./prow/cluster/components/cookie.txt"




#I think there is some amount of dependcy on starting these up so we go them one at a time
kubectl apply -f "./prow/cluster/components/01-ghproxy.yaml"
kubectl apply -f "./prow/cluster/components/02-cluster_config_maps.yaml"
kubectl apply -f "./prow/cluster/components/03-prowjob_custromresourcedefinition.yaml"
kubectl apply -f "./prow/cluster/components/04-hook.yaml"
kubectl apply -f "./prow/cluster/components/05-plank.yaml"
kubectl apply -f "./prow/cluster/components/06-sinker.yaml"
kubectl apply -f "./prow/cluster/components/07-deck.yaml"
kubectl apply -f "./prow/cluster/components/08-horologium.yaml"
kubectl apply -f "./prow/cluster/components/09-pushgateway.yaml"
kubectl apply -f "./prow/cluster/components/10-prow_addons_ctrlmanager.yaml"
kubectl apply -f "./prow/cluster/components/11-alb_ingress.yaml"
kubectl apply -f "./prow/cluster/components/12-crier.yaml"


kubectl create secret generic sa-s3-plank --from-file=service-account.json=./prow/cluster/components/service-account.json --dry-run -o yaml | kubectl replace secret sa-s3-plank -f -
