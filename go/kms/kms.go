package kms

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/jonahjon/eks-matrix-tests/util"
	"github.com/onsi/ginkgo"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/kubernetes/test/e2e/framework"
)

/**
	Overview
		- Test the KMS Encryption of Secrets with KMS master key

	Prerequisite
		- EKS Cluster 1.15 with KMS encryption turned on

	Test: Create kubernetes secrets, and configure those in the pod specs
	- using EnvFrom
	- using ValueFrom

	Expected Output
		- Kubernetes Secrets de-crypted into plaintext using KMS key provides the same value
 **/

var _ = ginkgo.Describe("[KMSsecrets]", func() {
	var f *framework.Framework
	f = framework.NewDefaultFramework("kms")
	var (
		ns      string
		sess    *session.Session
		region  string
		cluster string
		eksSvc  *eks.EKS
	)

	ginkgo.BeforeEach(func() {
		ginkgo.By(fmt.Sprintf("Checking KMS secret decryption on cluster: %v", cluster))
		ns = f.Namespace.Name

		//Get Clustername and Region from current context
		cluster = util.GetClusterNameOrDie()
		region = util.GetAWSRegionOrDie()
		sess = session.Must(session.NewSession(&aws.Config{
			Region: aws.String(region),
		}))

		eksSvc = eks.New(sess)
		describeClusterOut, err := eksSvc.DescribeCluster(&eks.DescribeClusterInput{
			Name: aws.String(cluster),
		})
		framework.ExpectNoError(err, "Describing cluster %v in namespace: %v", cluster, ns)
		// Verify that kms is configured and the keyarn configured matches an active KMS key in the account
		clusteroutput := *describeClusterOut.Cluster
		for _, v := range clusteroutput.EncryptionConfig {
			keyarn := listKmsKeysOrDie(f, sess, *v.Provider.KeyArn)
			if keyarn == "" {
				ginkgo.Skip(fmt.Sprintf("Skipping KMS test as no cluster KMS encryption configured for cluster: %v\n", cluster))
			}
			ginkgo.By(fmt.Sprintf("Found Active Encryption Key in Use: %v\n", keyarn))
		}
	})

	ginkgo.AfterEach(func() {
		ginkgo.By(fmt.Sprintf("Deleting Secrets in cluster from KMS secret test: %v", cluster))
	})

	// Create the secret, then the pod referencing it via `SecretKeyRef`, and verify the output is un-encrypted
	ginkgo.It("It should read an encrypted secret value from pod using `SecretKeyRef`", func() {
		name := "kms-test-" + string(uuid.NewUUID())
		secret := secretForTest(f.Namespace.Name, name)
		ginkgo.By(fmt.Sprintf("Creating Dummy Secret: \n"))
		var err error
		if secret, err = f.ClientSet.CoreV1().Secrets(f.Namespace.Name).Create(secret); err != nil {
			framework.ExpectNoError(err, "unable to create kubernetes secret %s: %v", secret.Name)
		}
		secretPod := newBusyBoxSecretKeyRef("secrets", name, "env")
		f.TestContainerOutput("SecretKeyRef", secretPod, 0, []string{
			"SECRET_DATA=value-1",
		})
	})

	// Create the secret, then the pod referencing it via `EnvFrom`, and verify the outputs are un-encrypted
	ginkgo.It("It should read all encrypted secret values from pod using `EnvFrom`", func() {
		cluster = util.GetClusterNameOrDie()
		ginkgo.By(fmt.Sprintf("Using Clustername %v\n", cluster))
		name := "kms-test-" + string(uuid.NewUUID())
		secret := secretForTest(f.Namespace.Name, name)
		ginkgo.By(fmt.Sprintf("Creating Dummy Secret: \n"))
		var err error
		if secret, err = f.ClientSet.CoreV1().Secrets(f.Namespace.Name).Create(secret); err != nil {
			framework.ExpectNoError(err, "unable to create kubernetes secret %s: %v", secret.Name)
		}
		secretPod := newBusyBoxEnvFrom("secrets", name, "env")
		f.TestContainerOutput("EnvFrom", secretPod, 0, []string{
			"data_1=value-1", "data_2=value-2", "data_3=value-3",
			"aws_data_1=value-1", "aws_data_2=value-2", "aws_data_3=value-3",
		})
	})
})

func listKmsKeysOrDie(f *framework.Framework, s *session.Session, key string) string {
	action := fmt.Sprintf("Listing KMS keys in account region to check if using Encryption Key...")
	ginkgo.By(action)
	kmsSvc := kms.New(s)
	allkeys, err := kmsSvc.ListKeys(&kms.ListKeysInput{})
	framework.ExpectNoError(err, "Listing KMS keys")
	for _, v := range allkeys.Keys {
		if strings.Contains(key, *v.KeyArn) {
			return *v.KeyArn
		}
	}
	return ""
}

func secretForTest(namespace, name string) *v1.Secret {
	return &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
		Data: map[string][]byte{
			"data_1": []byte("value-1\n"),
			"data_2": []byte("value-2\n"),
			"data_3": []byte("value-3\n"),
		},
	}
}

func newBusyBoxSecretKeyRef(name, keyname, command string) *v1.Pod {
	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: name + "-",
		},
		Spec: v1.PodSpec{
			RestartPolicy: v1.RestartPolicyNever,
			Containers: []v1.Container{
				{
					Name:    name,
					Image:   framework.BusyBoxImage,
					Command: []string{"/bin/sh"},
					Args:    []string{"-c", command},
					Env: []v1.EnvVar{
						{
							Name: "SECRET_DATA",
							ValueFrom: &v1.EnvVarSource{
								SecretKeyRef: &v1.SecretKeySelector{
									LocalObjectReference: v1.LocalObjectReference{
										Name: keyname,
									},
									Key: "data_1",
								},
							},
						},
					},
				},
			},
		},
	}
}

func newBusyBoxEnvFrom(name, keyname, command string) *v1.Pod {
	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: name + "-",
		},
		Spec: v1.PodSpec{
			RestartPolicy: v1.RestartPolicyNever,
			Containers: []v1.Container{
				{
					Name:    name,
					Image:   framework.BusyBoxImage,
					Command: []string{"/bin/sh"},
					Args:    []string{"-c", command},
					EnvFrom: []v1.EnvFromSource{
						{
							SecretRef: &v1.SecretEnvSource{LocalObjectReference: v1.LocalObjectReference{Name: keyname}},
						},
						{
							Prefix:    "aws_",
							SecretRef: &v1.SecretEnvSource{LocalObjectReference: v1.LocalObjectReference{Name: keyname}},
						},
					},
				},
			},
		},
	}
}
