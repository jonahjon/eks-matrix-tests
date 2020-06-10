# Running tests

Assuming your KUBECONFIG is ~/.kube/config and your current context is an EKS cluster in us-west-2:
```
env AWS_REGION=us-west-2 KUBECONFIG=~/.kube/config go test -timeout=0 -v -ginkgo.v ./e2e_test.go
```

# Writing tests

Create a package and write a test like the example test in `./cni`.

```
…
var _ = ginkgo.Describe("[CNI]", func() {
	var f *framework.Framework
	f = framework.NewDefaultFramework("cni")

	ginkgo.It("should enable pod-pod communication", func() {
…
```

Import the package in `e2e_test.go`.

```
…
	_ "go.amzn.com/eks-dataplane-tests/cni"
…
```

## Ginkgo

Upstream kubernetes uses the Ginkgo testing framework to organize e2e tests and
perform common actions before/after each test like create/delete a unique
namespace. Every `It` is a test "spec."

See also https://onsi.github.io/ginkgo/.

## The e2e framework

The [e2e
framework](https://github.com/kubernetes/kubernetes/tree/master/test/e2e/framework)
is a package used by upstream kubernetes. It contains many useful & reusable
functions like
[`WaitForPodNameRunningInNamespace`](https://github.com/kubernetes/kubernetes/blob/release-1.17/test/e2e/framework/pod/wait.go#L323).
You just need to call `NewDefaultFramework`, then the framework's before/after
actions will be performed and you will have access to a kubernetes client via
`f.Clientset`.

Note that the e2e framework is not yet a stable API and it's actively being
refactored to be more consumable. For example, in kubernetes 1.17 there was a
big cleanup of the e2e framework that moved/removed many functions:
https://github.com/kubernetes/kubernetes/issues/84380. But generally it's worth
the trouble to treat it as a library as opposed to rewriting or copying/pasting
functions.

See also
https://kubernetes.io/blog/2019/03/22/kubernetes-end-to-end-testing-for-everyone/.

## Dependency management

TODO: there is a working go.mod+go.sum checked in already but sometimes the framework/client-go/etc. do need to be updated
