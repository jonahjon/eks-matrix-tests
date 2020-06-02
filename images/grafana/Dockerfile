FROM 164382793440.dkr.ecr.us-west-2.amazonaws.com/bootstrap-helm

LABEL source=git@github.com:jonahjones/eks-matrix-tests.git

COPY runner.sh usr/local/bin/runner.sh

RUN chmod 755 usr/local/bin/runner.sh

ENTRYPOINT ["/bin/bash"]