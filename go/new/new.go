package new

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/jonahjon/eks-matrix-tests/util"
	"github.com/onsi/ginkgo"
	"k8s.io/kubernetes/test/e2e/framework"
)

//The goals of this test is to verify when configured with KMS envoloped secrets
//that we are able to reference secrets within pods, that have been encrypted/decrypted
//https://aws.amazon.com/blogs/containers/using-eks-encryption-provider-support-for-defense-in-depth/
var _ = ginkgo.Describe("[NEWtest]", func() {
	var f *framework.Framework
	f = framework.NewDefaultFramework("new")
	var (
		ns      string
		sess    *session.Session
		region  string
		cluster string
		eksSvc  *eks.EKS
	)
	ginkgo.BeforeEach(func() {
		ns = f.Namespace.Name

		//Get Clustername and Region from current context
		cluster = util.GetClusterNameOrDie()
		region = util.GetAWSRegionOrDie()
		sess = session.Must(session.NewSession(&aws.Config{
			Region: aws.String(region),
		}))

		eksSvc = eks.New(sess)
		ginkgo.By(fmt.Sprintf("session: %v \n", eksSvc))
	})
	ginkgo.AfterEach(func() {
		ginkgo.By(fmt.Sprintf("Cleaning up new test stuff"))
	})

	// Create the secret, then the pod referencing it via `SecretKeyRef`, and verify the output is un-encrypted
	ginkgo.It("New test purpose and expected`", func() {
		ginkgo.By(fmt.Sprintf("Cluster: %v \n", cluster))
		ginkgo.By(fmt.Sprintf("Namespace: %v \n", ns))
	})

	// Create the secret, then the pod referencing it via `EnvFrom`, and verify the outputs are un-encrypted
	ginkgo.It("New test purpose and expected2`", func() {
		ginkgo.By(fmt.Sprintf("Cluster: %v \n", cluster))
		ginkgo.By(fmt.Sprintf("Namespace: %v \n", ns))
	})
})
