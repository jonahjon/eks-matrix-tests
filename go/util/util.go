package util

import (
	"fmt"
	"strings"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubernetes/test/e2e/framework"
)

// added another way to do this without looping but feels equally hacky...

// GetClusterName gets the name of the EKS cluster being tested
// TODO this is hacky, maybe accept a positional arg instead
// func GetClusterNameOrDie() string {
// 	path := framework.TestContext.KubeConfig
// 	fmt.Printf("PATH %v\n", path)
// 	if path == "" {
// 		path = "/Users/jonahjo/.kube/config"
// 	}
// 	c, err := clientcmd.LoadFromFile(path)
// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}
// 	authInfo := c.AuthInfos[c.Contexts[c.CurrentContext].AuthInfo]
// 	for i, v := range authInfo.Exec.Args {
// 		if v == "-i" {
// 			return authInfo.Exec.Args[i+1]
// 			fmt.Printf("Using Clustername %v\n", authInfo.Exec.Args[i+1])
// 		}
// 	}
// 	return ""
// }

func GetClusterNameOrDie() string {
	path := framework.TestContext.KubeConfig
	c, err := clientcmd.LoadFromFile(path)
	if err != nil {
		fmt.Printf("%v", err)
	}
	cc := c.CurrentContext
	clusterurl := cc[strings.LastIndex(cc, "@")+1:]
	clustername := strings.Split(clusterurl, ".")[0]
	fmt.Printf("using EKS clustername: %v", clustername)
	return clustername
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

//     // Specify profile for config and region for requests
//     sess := session.Must(session.NewSessionWithOptions(session.Options{
//          Config: aws.Config{Region: aws.String("us-east-1")},
//          Profile: "profile_name",
//     }))
