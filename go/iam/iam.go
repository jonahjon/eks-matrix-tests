package iam

import (
	"fmt"
	"path"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/test/e2e/framework"
	e2epod "k8s.io/kubernetes/test/e2e/framework/pod"
	imageutils "k8s.io/kubernetes/test/utils/image"

	"github.com/jonahjon/eks-matrix-tests/util"
)

const (
	policyDocumentFormat = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:ListBucket"
      ],
      "Resource": [
        "arn:aws:s3:::%v"
      ]
    }
  ]
}`
	trustDocumentTemplate = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Federated": "arn:aws:iam::${AWS_ACCOUNT_ID}:oidc-provider/${OIDC_PROVIDER}"
      },
      "Action": "sts:AssumeRoleWithWebIdentity",
      "Condition": {
        "StringEquals": {
          "${OIDC_PROVIDER}:sub": "system:serviceaccount:${SERVICE_ACCOUNT_NAMESPACE}:${SERVICE_ACCOUNT_NAME}"
        }
      }
    }
  ]
}`
)

var _ = ginkgo.Describe("[IAMRolesforServiceAccounts]", func() {
	var f *framework.Framework
	f = framework.NewDefaultFramework("iam")

	var (
		ns                       string
		sess                     *session.Session
		cluster                  string
		eksSvc                   *eks.EKS
		region                   string
		iamSvc                   *iam.IAM
		issuer                   string
		openIDConnectProviderArn string
	)

	ginkgo.BeforeEach(func() {
		// https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html
		ginkgo.By(fmt.Sprintf("Enabling IAM roles for service accounts on cluster %v", cluster))

		ns = f.Namespace.Name

		//Get Clustername and Region from current context
		cluster = util.GetClusterNameOrDie()
		region = util.GetAWSRegionOrDie()
		sess = session.Must(session.NewSession(&aws.Config{
			Region: aws.String(region),
		}))

		cluster = util.GetClusterNameOrDie()

		eksSvc = eks.New(sess)
		describeClusterOut, err := eksSvc.DescribeCluster(&eks.DescribeClusterInput{
			Name: aws.String(cluster),
		})
		framework.ExpectNoError(err, "Describing cluster %v", cluster)
		// issuer looks like "https://oidc.eks.us-west-2.amazonaws.com/id/49DAFAB6B0C79B90D4A88E0B853A21C8"
		issuer = *describeClusterOut.Cluster.Identity.Oidc.Issuer

		iamSvc = iam.New(sess)
		createProviderOut, err := iamSvc.CreateOpenIDConnectProvider(&iam.CreateOpenIDConnectProviderInput{
			Url:            aws.String(issuer),
			ClientIDList:   aws.StringSlice([]string{"sts.amazonaws.com"}),
			ThumbprintList: aws.StringSlice([]string{"9e99a48a9960b14926bb7f3b02e22da2b0ab7280"}),
		})
		if err != nil {
			if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == iam.ErrCodeEntityAlreadyExistsException {
				listProvidersOut, err := iamSvc.ListOpenIDConnectProviders(&iam.ListOpenIDConnectProvidersInput{})
				framework.ExpectNoError(err, "Listing OpenID connect providers")
				id := path.Base(issuer)
				for _, v := range listProvidersOut.OpenIDConnectProviderList {
					if strings.Contains(*v.Arn, id) {
						openIDConnectProviderArn = *v.Arn
						break
					}
				}
			} else {
				framework.ExpectNoError(err, "Creating OpenID connect provider for cluster %v", cluster)
			}
		} else {
			openIDConnectProviderArn = *createProviderOut.OpenIDConnectProviderArn
		}
	})

	ginkgo.AfterEach(func() {
		ginkgo.By(fmt.Sprintf("Disabling IAM roles for service accounts on cluster %v", cluster))

		ginkgo.By(fmt.Sprintf("Deleting OpenID connect provider for cluster %v", cluster))
		_, err := iamSvc.DeleteOpenIDConnectProvider(&iam.DeleteOpenIDConnectProviderInput{
			OpenIDConnectProviderArn: aws.String(openIDConnectProviderArn),
		})
		if err != nil {
			framework.Logf("Error deleting OpenID connect provider: %v", err)
		}
	})

	ginkgo.It("should allow an authenticated pod to read an S3 bucket", func() {
		s3Svc := s3.New(sess)
		bucket := "bucket-" + ns
		key := "key-" + ns
		cleanS3Bucket := createS3BucketOrDie(s3Svc, bucket, key)
		defer cleanS3Bucket()

		action := fmt.Sprintf("Getting STS caller identity")
		ginkgo.By(action)
		stsSvc := sts.New(sess)
		getIdentityOut, err := stsSvc.GetCallerIdentity(&sts.GetCallerIdentityInput{})
		framework.ExpectNoError(err, action)
		account := *getIdentityOut.Account

		ginkgo.By("Creating IAM role and policy for our service account")
		policyDocument := fmt.Sprintf(policyDocumentFormat, bucket)
		policyName := "policy-" + ns
		roleName := "role-" + ns
		serviceAccountName := "s3-bucket-reader"
		cleanIAMPolicyAndRole := createIAMPolicyAndRoleOrDie(iamSvc, policyDocument, policyName, roleName, account, issuer, serviceAccountName, ns)
		defer cleanIAMPolicyAndRole()

		action = fmt.Sprintf("Creating our service account %v", serviceAccountName)
		ginkgo.By(action)
		serviceAccount := &v1.ServiceAccount{
			ObjectMeta: metav1.ObjectMeta{
				Name: serviceAccountName,
				Annotations: map[string]string{
					// https://docs.aws.amazon.com/eks/latest/userguide/specify-service-account-role.html
					"eks.amazonaws.com/role-arn": "arn:aws:iam::" + account + ":role/" + roleName,
				},
			},
		}
		serviceAccount, err = f.ClientSet.CoreV1().ServiceAccounts(ns).Create(serviceAccount)
		framework.ExpectNoError(err, action)

		podName := "allowed"
		createS3BucketReaderPodOrDie(f, ns, podName, serviceAccountName, bucket)

		action = fmt.Sprintf("Waiting for pod %v to successfully read S3 bucket %v", podName, bucket)
		ginkgo.By(action)
		err = e2epod.WaitForPodSuccessInNamespace(f.ClientSet, podName, ns)
		logs, _ := e2epod.GetPodLogs(f.ClientSet, ns, podName, podName)
		framework.Logf("Pod %q logs:\n %v", podName, logs)
		framework.ExpectNoError(err, action)
		gomega.Expect(strings.Contains(logs, "AccessDenied")).To(gomega.BeFalse(), "Pod log mustn't contain AccessDenied")
		gomega.Expect(strings.Contains(logs, key)).To(gomega.BeTrue(), "Pod log must contain S3 key")
	})

	ginkgo.It("should forbid an unauthenticated pod from reading an S3 bucket", func() {
		s3Svc := s3.New(sess)
		bucket := "bucket-" + ns
		key := "key-" + ns
		cleanS3Bucket := createS3BucketOrDie(s3Svc, bucket, key)
		defer cleanS3Bucket()

		podName := "forbidden"
		createS3BucketReaderPodOrDie(f, ns, podName, "default", bucket)

		action := fmt.Sprintf("Waiting for pod %v to fail to read S3 bucket %v", podName, bucket)
		ginkgo.By(action)
		err := e2epod.WaitForPodTerminatedInNamespace(f.ClientSet, podName, "", ns)
		logs, _ := e2epod.GetPodLogs(f.ClientSet, ns, podName, podName)
		framework.Logf("Pod %q logs:\n %v", podName, logs)
		framework.ExpectNoError(err, action)
		gomega.Expect(strings.Contains(logs, "AccessDenied")).To(gomega.BeTrue(), "Pod log must contain AccessDenied")
		gomega.Expect(strings.Contains(logs, key)).To(gomega.BeFalse(), "Pod log mustn't contain S3 key")
	})

})

func createS3BucketOrDie(s3Svc *s3.S3, bucket, key string) (clean func()) {
	action := fmt.Sprintf("Creating S3 bucket %v", bucket)
	ginkgo.By(action)
	_, err := s3Svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	framework.ExpectNoError(err, action)

	action = fmt.Sprintf("Putting key %v in S3 bucket %v", key, bucket)
	ginkgo.By(action)
	_, err = s3Svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	framework.ExpectNoError(err, action)

	return func() {
		ginkgo.By(fmt.Sprintf("Deleting key %v in S3 bucket %v", key, bucket))
		_, err := s3Svc.DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})
		if err != nil {
			framework.Logf("Error deleting S3 key: %v", err)
		}

		ginkgo.By(fmt.Sprintf("Deleting S3 bucket %v", bucket))
		_, err = s3Svc.DeleteBucket(&s3.DeleteBucketInput{
			Bucket: aws.String(bucket),
		})
		if err != nil {
			framework.Logf("Error deleting S3 bucket: %v", err)
		}
	}
}

func createS3BucketReaderPodOrDie(f *framework.Framework, ns, podName, serviceAccountName, bucket string) *v1.Pod {
	action := fmt.Sprintf("Creating pod %v authenticated as our service account %v", podName, serviceAccountName)
	ginkgo.By(action)
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: podName,
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:    podName,
					Image:   imageutils.GetE2EImage(imageutils.Agnhost),
					Command: strings.Fields("bash -c"),
					Args:    []string{"apk update && apk add python3 && pip3 install awscli && aws s3 ls " + bucket},
				},
			},
			ServiceAccountName: serviceAccountName,
			RestartPolicy:      v1.RestartPolicyNever,
		},
	}
	pod, err := f.ClientSet.CoreV1().Pods(ns).Create(pod)
	framework.ExpectNoError(err, action)

	return pod
}

// createIAMPolicyAndRoleOrDie does almost exactly what is documented here:
// https://docs.aws.amazon.com/eks/latest/userguide/create-service-account-iam-policy-and-role.html
func createIAMPolicyAndRoleOrDie(iamSvc *iam.IAM, policyDocument, policyName, roleName, account, issuer, serviceAccountName, serviceAccountNamespace string) (clean func()) {
	action := fmt.Sprintf("Creating IAM policy %v", policyName)
	ginkgo.By(action)
	framework.Logf("PolicyDocument:\n %v", policyDocument)
	createPolicyOut, err := iamSvc.CreatePolicy(&iam.CreatePolicyInput{
		PolicyDocument: aws.String(policyDocument),
		PolicyName:     aws.String(policyName),
	})
	framework.ExpectNoError(err, action)
	policyArn := *createPolicyOut.Policy.Arn

	action = fmt.Sprintf("Creating IAM role %v", roleName)
	ginkgo.By(action)
	trustDocument := strings.ReplaceAll(trustDocumentTemplate, "${AWS_ACCOUNT_ID}", account)
	trustDocument = strings.ReplaceAll(trustDocument, "${OIDC_PROVIDER}", strings.ReplaceAll(issuer, "https://", ""))
	trustDocument = strings.ReplaceAll(trustDocument, "${SERVICE_ACCOUNT_NAMESPACE}", serviceAccountNamespace)
	trustDocument = strings.ReplaceAll(trustDocument, "${SERVICE_ACCOUNT_NAME}", serviceAccountName)
	framework.Logf("AssumeRolePolicyDocument:\n %v", trustDocument)
	_, err = iamSvc.CreateRole(&iam.CreateRoleInput{
		AssumeRolePolicyDocument: aws.String(trustDocument),
		RoleName:                 aws.String(roleName),
	})
	framework.ExpectNoError(err, action)

	action = fmt.Sprintf("Attaching IAM policy %v to IAM role %v", policyArn, roleName)
	ginkgo.By(action)
	_, err = iamSvc.AttachRolePolicy(&iam.AttachRolePolicyInput{
		PolicyArn: aws.String(policyArn),
		RoleName:  aws.String(roleName),
	})
	framework.ExpectNoError(err, action)

	return func() {
		ginkgo.By(fmt.Sprintf("Detaching IAM policy %v from IAM role %v", policyArn, roleName))
		_, err := iamSvc.DetachRolePolicy(&iam.DetachRolePolicyInput{
			PolicyArn: aws.String(policyArn),
			RoleName:  aws.String(roleName),
		})
		if err != nil {
			framework.Logf("Error detaching IAM policy: %v", err)
		}

		ginkgo.By(fmt.Sprintf("Deleting IAM role %v", roleName))
		_, err = iamSvc.DeleteRole(&iam.DeleteRoleInput{
			RoleName: aws.String(roleName),
		})
		if err != nil {
			framework.Logf("Error deleting IAM role: %v", err)
		}

		ginkgo.By(fmt.Sprintf("Deleting IAM policy %v", policyArn))
		_, err = iamSvc.DeletePolicy(&iam.DeletePolicyInput{
			PolicyArn: aws.String(policyArn),
		})
		if err != nil {
			framework.Logf("Error deleting IAM policy: %v", err)
		}
	}
}
