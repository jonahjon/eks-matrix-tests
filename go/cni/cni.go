package cni

import (
	"github.com/onsi/ginkgo"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/test/e2e/framework"
	e2enode "k8s.io/kubernetes/test/e2e/framework/node"
	e2epod "k8s.io/kubernetes/test/e2e/framework/pod"
)

var _ = ginkgo.Describe("[CNI]", func() {
	var f *framework.Framework
	f = framework.NewDefaultFramework("cni")

	ginkgo.It("should enable pod-pod communication", func() {
		serverPod := newBusyBoxPod("server", "sleep 60")
		serverPod, err := f.ClientSet.CoreV1().Pods(f.Namespace.Name).Create(serverPod)
		framework.ExpectNoError(err, "creating pod")
		err = f.WaitForPodRunning(serverPod.Name)
		framework.ExpectNoError(err, "waiting for pod running")
		serverPod, err = f.ClientSet.CoreV1().Pods(f.Namespace.Name).Get(serverPod.Name, metav1.GetOptions{})
		framework.ExpectNoError(err, "getting pod")

		clientPod := newBusyBoxPod("client", "ping -c 3 -w 2 -w 30 "+serverPod.Status.PodIP)
		clientPod, err = f.ClientSet.CoreV1().Pods(f.Namespace.Name).Create(clientPod)
		framework.ExpectNoError(err, "creating pod")
		err = e2epod.WaitForPodSuccessInNamespace(f.ClientSet, clientPod.Name, f.Namespace.Name)
		framework.ExpectNoError(err, "waiting for pod success")

		err = f.ClientSet.CoreV1().Pods(f.Namespace.Name).Delete(serverPod.Name, &metav1.DeleteOptions{})
		framework.ExpectNoError(err, "deleting pod")
	})

	ginkgo.It("should enable pod-node communication", func() {
		node, err := e2enode.GetRandomReadySchedulableNode(f.ClientSet)
		framework.ExpectNoError(err, "getting random ready schedulable node")
		internalIP, err := e2enode.GetInternalIP(node)
		framework.ExpectNoError(err, "getting node internal IP")

		clientPod := newBusyBoxPod("client", "ping -c 3 -w 2 -w 30 "+internalIP)
		clientPod, err = f.ClientSet.CoreV1().Pods(f.Namespace.Name).Create(clientPod)
		framework.ExpectNoError(err, "creating pod")
		err = e2epod.WaitForPodSuccessInNamespace(f.ClientSet, clientPod.Name, f.Namespace.Name)
		framework.ExpectNoError(err, "waiting for pod success")
	})

})

func newBusyBoxPod(name, command string) *v1.Pod {
	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: name + "-",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:    name,
					Image:   framework.BusyBoxImage,
					Command: []string{"/bin/sh"},
					Args:    []string{"-c", command},
				},
			},
			RestartPolicy: v1.RestartPolicyNever,
		},
	}
}
