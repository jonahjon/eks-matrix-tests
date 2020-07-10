FROM 164382793440.dkr.ecr.us-west-2.amazonaws.com/bootstrap-helm

LABEL source=git@github.com:jonahjones/eks-matrix-tests.git

COPY helm.sh /home/prow/go/src/helm.sh

RUN chmod 755 /home/prow/go/src/helm.sh

COPY kubectl.sh /home/prow/go/src/kubectl.sh

RUN chmod 755 /home/prow/go/src/kubectl.sh

ENTRYPOINT ["/bin/bash"]
