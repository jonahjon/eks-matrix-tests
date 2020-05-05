


## K8s Feature Tests Test Coverage

EKS Matrix test results of latest prow deck results http://a359ac6a366bc4f4abf4c3581be5110d-99053184.us-west-2.elb.amazonaws.com/

badge sytnax is:

	http://PROW-URL/badge.svg?jobs=eks-prod-pdx-eks-1.15

### EKS-Monitoring

| Provider/K8s | v1.15 | v1.14 |  v1.13 |  v1.12 |
| ----------- | -----------| ----------- |----------- |----------- 
| **HPA** | [![EKS HPA 1.15 test](http://a359ac6a366bc4f4abf4c3581be5110d-99053184.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-hpa-1.15)](http://a359ac6a366bc4f4abf4c3581be5110d-99053184.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-hpa-1.15) | [![EKS HPA 1.14 test](http://a359ac6a366bc4f4abf4c3581be5110d-99053184.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-hpa-1.14)](http://a359ac6a366bc4f4abf4c3581be5110d-99053184.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-hpa-1.14) | [![EKS HPA 1.13 test](http://a359ac6a366bc4f4abf4c3581be5110d-99053184.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-hpa-1.13)](http://a359ac6a366bc4f4abf4c3581be5110d-99053184.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-hpa-1.13)| [![EKS HPA 1.12 test](http://a359ac6a366bc4f4abf4c3581be5110d-99053184.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-hpa-1.12)](http://a359ac6a366bc4f4abf4c3581be5110d-99053184.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-hpa-1.12)
| **Base Tests** | [![EKS Tests 1.15 test](http://a359ac6a366bc4f4abf4c3581be5110d-99053184.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-eks-1.15)](http://a359ac6a366bc4f4abf4c3581be5110d-99053184.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-eks-1.15) | [![EKS Tests 1.14 test](http://a359ac6a366bc4f4abf4c3581be5110d-99053184.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-eks-1.14)](http://a359ac6a366bc4f4abf4c3581be5110d-99053184.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-eks-1.14) | [![EKS Tests 1.13 test](http://a359ac6a366bc4f4abf4c3581be5110d-99053184.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-eks-1.13)](http://a359ac6a366bc4f4abf4c3581be5110d-99053184.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-eks-1.13)| [![EKS Tests 1.12 test](http://a359ac6a366bc4f4abf4c3581be5110d-99053184.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-eks-1.12)](http://a359ac6a366bc4f4abf4c3581be5110d-99053184.us-west-2.elb.amazonaws.com/badge.svg?jobs=eks-prod-pdx-eks-1.12)



### EKS-Security (TODO Show some real tests)


Prow Setup
- Spinup EKS Cluster
- Make sure nodegroup has (Cloudformation priv)..
--
