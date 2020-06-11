#!/bin/bash

export AWS_DEFAULT_REGION=us-west-2

aws ecr create-repository --repository-name alpine-kubectl || true
aws ecr create-repository --repository-name bootstrap || true
aws ecr create-repository --repository-name bootstrap-helm || true
aws ecr create-repository --repository-name golang || true
aws ecr create-repository --repository-name prow/update-jobs || true
aws ecr create-repository --repository-name grafana/alpine-kubectl || true
aws ecr create-repository --repository-name github-runner || true


aws ecr create-repository --repository-name voting-app/vote
aws ecr create-repository --repository-name voting-app/worker
aws ecr create-repository --repository-name voting-app/results



