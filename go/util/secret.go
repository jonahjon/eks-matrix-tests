package util

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/kubernetes/test/e2e/framework"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	kubeconfig                = "kubeconfig"
	loggingIngressCredentials = "logging-ingress-credentials"
	password                  = "password"
	token                     = "token"
	KubecfgSecretName         = "kubecfg"
)

// func getAdminToken(ctx context.Context, f *framework.Framework, namespace, secretName, objectKey string) (string, error) {
// 	return GetObjectFromSecret(ctx, f, namespace, secretName, objectKey)
// }

func getLoggingPassword(ctx context.Context, f *framework.Framework, namespace, secretName, objectKey string) (string, error) {
	return GetObjectFromSecret(ctx, f, namespace, secretName, objectKey)
}

// GetObjectFromSecret returns object from secret
func GetObjectFromSecret(ctx context.Context, k8sClient *framework.Framework, namespace, secretName, objectKey string) (string, error) {
	secret := &corev1.Secret{}
	err := k8sClient.Client().Get(ctx, client.ObjectKey{Namespace: namespace, Name: secretName}, secret)
	if err != nil {
		return "", err
	}
	if _, ok := secret.Data[objectKey]; ok {
		return string(secret.Data[objectKey]), nil
	}
	return "", fmt.Errorf("secret %s/%s did not contain object key %q", namespace, secretName, objectKey)
}
