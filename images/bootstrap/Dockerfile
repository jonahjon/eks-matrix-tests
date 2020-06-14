FROM debian:stretch

WORKDIR /workspace
RUN mkdir -p /workspace
ENV WORKSPACE=/workspace \
    TERM=xterm

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    ca-certificates \
    curl \
    git \
    jq \
    dnsutils \
    openssh-client \
    python \
    wget \
    apt-transport-https \
    gnupg2 \
    shellcheck \
    software-properties-common \
    lsb-release \
    gettext \
    awscli \
    groff \
    unzip \
    && apt-get clean

RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" \
    && unzip awscliv2.zip \
    && ./aws/install \
    && rm awscliv2.zip \
    && rm -R aws

ENV KUBECTL_VERSION="v1.18.2"

RUN	wget https://storage.googleapis.com/kubernetes-release/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl -O /usr/local/bin/kubectl &&\
	chmod +x /usr/local/bin/kubectl

# Docker-in-docker
RUN curl -fsSL https://download.docker.com/linux/$(. /etc/os-release; echo "$ID")/gpg \
    | apt-key add - && \
    add-apt-repository \
    "deb [arch=amd64] https://download.docker.com/linux/$(. /etc/os-release; echo "$ID") \
    $(lsb_release -cs) stable"

ENV DOCKER_VERSION=18.06.1*

ARG EKS114_KUBE_CONFIG
ENV EKS114_KUBE_CONFIG=$EKS114_KUBE_CONFIG

ARG EKS115_KUBE_CONFIG
ENV EKS115_KUBE_CONFIG=$EKS115_KUBE_CONFIG

ARG EKS116_KUBE_CONFIG
ENV EKS116_KUBE_CONFIG=$EKS116_KUBE_CONFIG

ARG PROW_KUBE_CONFIG
ENV PROW_KUBE_CONFIG=$PROW_KUBE_CONFIG

COPY tmp/aws-iam-authenticator /bin/aws-iam-authenticator

RUN apt-get update && \
    apt-get install -y --no-install-recommends docker-ce=${DOCKER_VERSION}