FROM 164382793440.dkr.ecr.us-west-2.amazonaws.com/bootstrap-helm

LABEL source=git@github.com:jonahjones/eks-matrix-tests.git

COPY helm.sh usr/local/bin/helm.sh

RUN chmod 755 usr/local/bin/helm.sh

COPY kubectl.sh usr/local/bin/kubectl.sh

RUN chmod 755 usr/local/bin/kubectl.sh

COPY templates templates

ENTRYPOINT ["/bin/bash"]
