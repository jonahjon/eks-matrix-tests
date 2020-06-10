package util

import (
	"fmt"
	"strings"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubernetes/test/e2e/framework"
)

// GetClusterName gets the name of the EKS cluster being tested
// TODO this is hacky, maybe accept a positional arg instead
func GetClusterNameOrDie() string {
	c, err := clientcmd.LoadFromFile(framework.TestContext.KubeConfig)
	framework.ExpectNoError(err, "failed to load kubeconfig from file")

	authInfo := c.AuthInfos[c.Contexts[c.CurrentContext].AuthInfo]

	for i, v := range authInfo.Exec.Args {
		// aws-iam-authenticator token
		if v == "-i" {
			return authInfo.Exec.Args[i+1]
		}
		// aws eks get-token
		if v == "--cluster-name" {
			return authInfo.Exec.Args[i+1]
		}
	}
	framework.Fail("failed to get EKS cluster name")
	return ""
}

func GetAWSRegionOrDie() string {
	path := framework.TestContext.KubeConfig
	c, err := clientcmd.LoadFromFile(path)
	if err != nil {
		fmt.Printf("%v", err)
	}
	cc := c.CurrentContext
	clusterurl := cc[strings.LastIndex(cc, "@")+1:]
	region := strings.Split(clusterurl, ".")[1]
	fmt.Printf("using AWS region: %v", region)
	return region
}
