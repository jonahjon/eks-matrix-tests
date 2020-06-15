FROM alpine:3.11

LABEL source=git@github.com:jonahjones/eks-matrix-tests.git

ENV HELM_3_VERSION="v3.2.1"
ENV KUBECTL_VERSION="v1.17.3"

RUN apk update &&\
    apk add --no-cache openssl coreutils curl bash jq grep &&\
	wget https://storage.googleapis.com/kubernetes-release/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl -O /usr/local/bin/kubectl &&\
	chmod +x /usr/local/bin/kubectl &&\
	wget https://get.helm.sh/helm-${HELM_3_VERSION}-linux-amd64.tar.gz -O - | tar -xzO linux-amd64/helm > /usr/local/bin/helm &&\
	chmod +x /usr/local/bin/helm
