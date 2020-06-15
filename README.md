## Aquarium K8s Test Coverage

This repository contains tools and configuration files for the testing 3rd party software on various versions of Amazon EKS. 


Aquarium is built upon [Prow](https://github.com/kubernetes/test-infra/tree/master/prow) which is a Kubernetes based CI/CD system. Jobs can be triggered by various types of events and report their status to many different services. In addition to job execution, Prow provides GitHub automation in the form of policy enforcement, chat-ops via `/foo` style commands.


Current Supported Platforms:
- EKS K8s version 1.14
- EKS K8s version 1.15
- EKS K8s version 1.16

The [architecture diagram](static/architecture.png) provides an overview of how the different services interact, and test on EKS clusters.


## Getting Started Testing Your Software 

For onboarding 3rd party software onto Aquarium there a minimum of two steps needed to be completed. Fork this repoisotory, and create a PR to get included into this system.

1. Add in a job definition and OWNERS file To read more about about job configuration check out [Prow Job FAQ](/prow/jobs/README.md#adding-or-updating-jobs)

## Job Examples

A presubmit job named "pull-community-verify" that will run against all PRs to
kubernetes/community's master branch. It will run `make verify` in a checkout
of kubernetes/community at the PR's HEAD. It will report back to the PR via a
status context named `pull-kubernetes-community`. Its logs and results are going
to end up in GCS under `kubernetes-jenkins/pr-logs/pull/community`. Historical
results will display in testgrid on the `sig-contribex-community` dashboard
under the `pull-verify` tab

```yaml
presubmits:
  kubernetes/community:
  - name: pull-community-verify  # convention: (job type)-(repo name)-(suite name)
    annotations:
      testgrid-dashboards: sig-contribex-community
      testgrid-tab-name: pull-verify
    branches:
    - master
    decorate: true
    always_run: true
    spec:
      containers:
      - image: golang:1.12.5
        command:
        - /bin/bash
        args:
        - -c
        # Add GOPATH/bin back to PATH to workaround #9469
        - "export PATH=$GOPATH/bin:$PATH && make verify"
```







2. Add in a job definition and OWNERS file
Add in your test image and OWNERS file



- [Create a folder and OWNER file](/prow/jobs/README.md#adding-or-updating-jobs)
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


