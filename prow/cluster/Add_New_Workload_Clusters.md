To add a new workload cluster you must populate a kubeconfig secret that gets mounted by various prow componenets.


This is the basic structure required by the secret kubeconfig

The cluster certificate-authority-data and server endpoint can be found in a kubeconfig file.

To get the user: token you must use gencred.

gencred is found in https://github.com/kubernetes/test-infra/tree/master/gencred


```
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: xxxxxxx
    server: https://xxxxxxxxxx.us-west-2.eks.amazonaws.com
  name: eks-114
contexts:
- context:
    cluster: eks-114
    user: eks-114
  name: eks-114
current-context: default
kind: Config
preferences: {}
users:
- name: eks-114
  user:
    token: xxxxxxxxxxx
```

- export KUBECONFIG=1_14_cluster.config
- export CONTEXT=$(kubectl config current-context)
- git clone https://github.com/kubernetes/test-infra
- cd test-infra/gencred
- bazel build //gencred --context $CONTEXT --name eks-114 --output ../../prow/cluster/components/workload_clusters.yaml --serviceaccount

