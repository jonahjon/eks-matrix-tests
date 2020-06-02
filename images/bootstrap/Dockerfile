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
    && apt-get clean

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

COPY data/1_14_cluster.config 1_14_cluster.config
COPY data/1_15_cluster.config 1_15_cluster.config
COPY data/1_16_cluster.config 1_16_cluster.config

COPY tmp/aws-iam-authenticator /bin/aws-iam-authenticator

RUN apt-get update && \
    apt-get install -y --no-install-recommends docker-ce=${DOCKER_VERSION}