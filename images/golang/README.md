# Buildpack Golang Docker Image

## Overview

This folder contains the Buildpack Golang image that is based on the Bootstrap image. Use it to build Golang components.

The image consists of:

- golang 1.14.2
- dep 0.5.4

or 

- golang 1.13.11
- dep 0.5.4

if you use tag ```go1.13```

## Installation

To build the Docker image, run this command:

```bash
docker build -t eks-matrix/golang .
```