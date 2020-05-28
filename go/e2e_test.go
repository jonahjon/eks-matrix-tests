package e2e

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/kubernetes/test/e2e/framework"
	"k8s.io/kubernetes/test/e2e/framework/config"

	// ensure that cloud providers are loaded
	_ "k8s.io/kubernetes/test/e2e/framework/providers/aws"

	// test sources
	_ "github.com/jonahjon/eks-matrix-tests/cni"
	// _ "github.com/jonahjon/eks-matrix-tests/iam"
	// _ "github.com/jonahjon/eks-matrix-tests/kms"
	// _ "github.com/jonahjon/eks-matrix-tests/logging"
	_ "github.com/jonahjon/eks-matrix-tests/new"
)

func TestMain(m *testing.M) {
	fmt.Printf("STARTING\n")
	config.CopyFlags(config.Flags, flag.CommandLine)
	framework.RegisterCommonFlags(flag.CommandLine)
	framework.RegisterClusterFlags(flag.CommandLine)
	flag.Parse()
	framework.AfterReadingAllFlags(&framework.TestContext)
	RegisterFailHandler(Fail)
	// Seed to e.g. randomize selection of node
	rand.Seed(time.Now().UnixNano())
	os.Exit(m.Run())
}

func TestE2E(t *testing.T) {
	RunSpecs(t, "EKS Suite")
}
