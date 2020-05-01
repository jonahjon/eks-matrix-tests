package logging

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/jonahjon/eks-matrix-tests/util"
	"github.com/onsi/ginkgo"
	"k8s.io/kubernetes/test/e2e/framework"
)

/**
	Overview
		- Tests the Logging Capability for service

	Prerequisite
		- Shoot Cluster with  Condition "APIServerAvailable" equals true

	Test: Scale down API Server deployment of the Shoot in the Seed

	Expected Output
		- Shoot Condition "APIServerAvailable" becomes unhealthy
 **/

const (
	// logsCount uint64 = 10000

	initializationTimeout = 15 * time.Minute
	// k
	ibanaAvailableTimeout = 10 * time.Second
	// getLogsFromElasticsearchTimeout = 5 * time.Minute

	// fluentBitName = "fluent-bit"
	// fluentdName   = "fluentd-es"
	// logger        = "logger"
)

var _ = ginkgo.Describe("[Logging]", func() {
	var f *framework.Framework
	//Framework creates namespace based of this name "logging", needs to be lowercase
	f = framework.NewDefaultFramework("logging")
	var (
		ns      string
		sess    *session.Session
		region  string
		cluster string
		eksSvc  *eks.EKS
	)
	util.CBeforeSuite(func(ctx context.Context) {

	}, initializationTimeout)

	util.CBeforeEach(func(ctx context.Context) {
		ns = f.Namespace.Name
		//Get Clustername and Region from current context
		cluster = util.GetClusterNameOrDie()
		region = util.GetAWSRegionOrDie()
		sess = session.Must(session.NewSession(&aws.Config{
			Region: aws.String(region),
		}))
		eksSvc = eks.New(sess)
		ginkgo.By(fmt.Sprintf("session: %v \n", eksSvc))
	}, initializationTimeout)

	ginkgo.AfterEach(func() {
		ginkgo.By(fmt.Sprintf("Cleaning up new test stuff"))
	})

	// Create the secret, then the pod referencing it via `SecretKeyRef`, and verify the output is un-encrypted
	ginkgo.It("Logging test purpose and expected`", func() {
		ginkgo.By(fmt.Sprintf("Cluster: %v \n", cluster))
		loggingPassword, err := util.GetLoggingPassword(ctx, f, ns)
		framework.ExpectNoError(err)
		ginkgo.By(fmt.Sprintf("Namespace: %v \n", ns))
	})

	// Create the secret, then the pod referencing it via `EnvFrom`, and verify the outputs are un-encrypted
	ginkgo.It("Logging test purpose and expected2`", func() {
		ginkgo.By(fmt.Sprintf("Cluster: %v \n", cluster))
		ginkgo.By(fmt.Sprintf("Namespace: %v \n", ns))
	})
})
