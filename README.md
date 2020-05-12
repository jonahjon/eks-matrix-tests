


## K8s Feature Tests Test Coverage

EKS Matrix test results of latest prow deck results http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/

badge sytnax is:

	http://PROW-URL/badge.svg?jobs=eks-prod-pdx-eks-1.15

### EKS-Monitoring


| Provider/K8s | v1.15 | v1.14 |  v1.13 |  v1.12 |
| ----------- | -----------| ----------- |----------- |----------- 
| **Open-Source** | [![EKS HPA 1.15 test](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=open-source-readme)](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-hpa-1.15) | [![EKS HPA 1.14 test](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-hpa-1.14)](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-hpa-1.14) | [![EKS HPA 1.13 test](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-hpa-1.13)](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-hpa-1.13)| [![EKS HPA 1.12 test](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-hpa-1.12)](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-hpa-1.12)
| **Base Tests** | [![EKS Tests 1.15 test](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-eks-1.15)](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-eks-1.15) | [![EKS Tests 1.14 test](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-eks-1.14)](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-eks-1.14) | [![EKS Tests 1.13 test](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-eks-1.13)](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-eks-1.13)| [![EKS Tests 1.12 test](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-eks-1.12)](http://a69660e52137f4cbcaefaf44e7c02ebb-1275564336.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-eks-1.12)



### EKS-Security (TODO Show some real tests)


Prow Setup
- Spinup EKS Cluster
- Make sure nodegroup has (Cloudformation priv)..
