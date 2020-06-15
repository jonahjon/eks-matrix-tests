FROM 164382793440.dkr.ecr.us-west-2.amazonaws.com/golang:go1.13

LABEL source=git@github.com:jonahjones/eks-matrix-tests.git

COPY images/update-jobs/update-jobs.sh usr/local/bin/update-jobs.sh

RUN ls -la usr/local/bin

RUN chmod 755 usr/local/bin/update-jobs.sh

ENTRYPOINT ["/bin/bash"]
