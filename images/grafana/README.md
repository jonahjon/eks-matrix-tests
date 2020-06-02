# Kubectl Docker Image

## Overview

This folder contains a minimal Docker image used to control the cluster.

This image consists of:

- alpine linux 3.8
- openssl
- curl
- base64
- kubectl (1.18)
- helm (3.2.1)
- aws-iam-authenticator

## Installation

To build the Docker image, run this command:

```bash
docker build -t grafana/alpine-kubectl .
```
test
