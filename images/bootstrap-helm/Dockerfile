FROM 164382793440.dkr.ecr.us-west-2.amazonaws.com/bootstrap

ENV HELM_VERSION="v3.2.1"

RUN wget -q https://get.helm.sh/helm-${HELM_VERSION}-linux-amd64.tar.gz -O - | tar -xzO linux-amd64/helm > /usr/local/bin/helm \
    && chmod +x /usr/local/bin/helm