#!/bin/bash

export AWS_DEFAULT_REGION=us-west-2

aws ecr create-repository --repository-name alpine-kubectl
aws ecr create-repository --repository-name bootstrap
aws ecr create-repository --repository-name bootstrap-helm
aws ecr create-repository --repository-name eks-matrix/golang
aws ecr create-repository --repository-name grafana/alpine-kubectl