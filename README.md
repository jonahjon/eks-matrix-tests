![Push Bootstrap ECR image with cluster access](https://github.com/jonahjon/eks-matrix-tests/workflows/Push%20Bootstrap%20ECR%20image%20with%20cluster%20access/badge.svg?branch=master)

## Aquarium K8s Test Coverage

This repository contains tools and configuration files for the testing 3rd party software on various versions of Amazon EKS.

The [architecture diagram](static/architecture.png) provides an overview of how the different services interact, and test on EKS.

## Onboarding Steps

For onboarding 3rd party software onto Aquarium there a minimum of two steps needed to be completed. 

- [Add or update job definition](/prow/jobs/README.md#adding-or-updating-jobs)
- [Add or update tests](/images/README.md#adding-or-updating-tests)


Now that you've created both a job definition, and job test. 

Aquarium will listen to part of the job spec called "run_if_changed" to listen to file changes on PR's and kickoff the testing process when it sees changes.

```    
run_if_changed: '^images/grafana/'
```

In this example any files within ```images/grafana```  that get changed on a PR will kickoff a testing job defined by yaml file ```prow/jobs/grafana/grafana.yaml```.

## Results and Logs

Aquarium test results, and test logs can be found on the deck server http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/

The [architecture diagram](static/architecture.png) provides an overview of how the different services interact, and test on EKS.

![](static/prow_navigation.png)


## Re-Testing PR Workflow




## Build Badges

basic badge sytnax for deck jobs is:

	http://PROW-URL/badge.svg?jobs=eks-open-source-readme


# Test Results


### Grafana

| Install Method | v1.16 | v1.15 | v1.14 | 
| ----------- | ----------- | ----------- | -----------
| **Helm** | [![Helm 1.16](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=grafana-helm-1-1.16)](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=grafana-helm-1-1.16) | [![Helm 1.15](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=grafana-helm-1-1.15)](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=grafana-helm-1-1.15) | [![Helm 1.14](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=grafana-helm-1-1.14)](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=grafana-helm-1-1.14)
| **Kubectl** | [![Kubectl 1.16](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=open-source-product2-1.16)](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=open-source-product2-1.16) | [![Kubectl 1.15](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=open-source-product2-1.15)](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=open-source-product2-1.15) | [![Kubectl 1.14](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=open-source-product2-1.14)](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=open-source-product2-1.14)


### Prometheus

| Install Method | v1.16 | v1.15 | v1.14 | 
| ----------- | ----------- | ----------- | -----------
| **Helm** | [![EKS HPA 1.16 test](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=partner-product1-1.16)](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=partner-product1-1.16) | [![EKS HPA 1.15 test](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=partner-product1-1.15)](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=partner-product1-1.15) | [![EKS HPA 1.14 test](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=partner-product1-1.14)](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=partner-product1-1.14)
| **Kubectl** | [![EKS Tests 1.16 test](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=partner-product2-1.16)](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=partner-product2-1.16) | [![EKS Tests 1.15 test](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=partner-product2-1.15)](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=partner-product2-1.15) | [![EKS Tests 1.14 test](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=partner-product2-1.14)](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=partner-product2-1.14)


