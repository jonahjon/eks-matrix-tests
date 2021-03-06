#!/bin/bash

eksctl utils associate-iam-oidc-provider --region=us-west-2 --cluster=aquarium --approve

kubectl apply -k "github.com/kubernetes-sigs/aws-ebs-csi-driver/deploy/kubernetes/overlays/stable/?ref=master"

kubectl create configmap plugins --from-file=plugins.yaml=./prow/cluster/components/plugins.yaml
kubectl create configmap config --from-file "./prow/cluster/components/config.yaml"
kubectl create configmap job-config --from-file "./prow/jobs/config.yaml"
kubectl create configmap branding --from-file "./prow/branding"
kubectl create secret generic hmac-token --from-file=hmac=./prow/cluster/components/bot_hmac
kubectl create secret generic oauth-token --from-file=oauth=./prow/cluster/components/bot_oauth
kubectl create secret generic github-oauth-config --from-file=secret="./prow/cluster/components/deck_oauth"
kubectl create secret generic cookie --from-file=secret="./prow/cluster/components/cookie.txt"
kubectl create secret generic kubeconfig --from-file=config="./prow/cluster/components/workload_clusters.yaml"

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
kubectl apply -f "./prow/cluster/components/11-alb_ingress.yaml"
kubectl apply -f "./prow/cluster/components/12-crier.yaml"

# Create the s3 creds for the blog storage buckets
kubectl create secret generic s3-credentials --from-file=service-account.json=./prow/cluster/components/service-account.json --dry-run -o yaml | kubectl replace secret sa-s3-plank -f -

kubectl create secret generic sa-s3-plank --from-file=service-account.json=./prow/cluster/components/service-account.json --dry-run -o yaml | kubectl replace secret sa-s3-plank -f -

# # Create Service Account for Plank Decoration and dropping pod info into bucket logs
eksctl create iamserviceaccount \
                --name s3-deck \
                --namespace default \
                --cluster aquarium \
                --attach-policy-arn arn:aws:iam::aws:policy/AmazonS3FullAccess \
                --approve

# Create the SA for the s3 access for crier and spyglass
eksctl create iamserviceaccount \
                --name s3-crier \
                --namespace default \
                --cluster aquarium \
                --attach-policy-arn arn:aws:iam::aws:policy/AmazonS3FullAccess \
                --approve


eksctl create iamserviceaccount \
                --name s3-prow-controller-manager \
                --namespace default \
                --cluster aquarium \
                --attach-policy-arn arn:aws:iam::aws:policy/AmazonS3FullAccess \
                --approve

#alb ingress
kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/aws-alb-ingress-controller/v1.1.8/docs/examples/rbac-role.yaml

# create the alb ingress controller SA
eksctl create iamserviceaccount \
    --region us-west-2 \
    --name alb-ingress-controller \
    --namespace kube-system \
    --cluster aquarium \
    --attach-policy-arn arn:aws:iam::164382793440:policy/ALBIngressControllerIAMPolicy \
    --override-existing-serviceaccounts \
    --approve

