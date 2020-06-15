FROM 164382793440.dkr.ecr.us-west-2.amazonaws.com/golang:go1.13

LABEL source=git@github.com:jonahjones/eks-matrix-tests.git

COPY update-jobs.sh /workspace/update-jobs.sh

RUN chmod 755 /workspace/update-jobs.sh

ENTRYPOINT ["/bin/bash"]
