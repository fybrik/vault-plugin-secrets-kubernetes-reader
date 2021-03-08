package kubesecrets

import (
	"context"

	log "github.com/hashicorp/go-hclog"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kclient "sigs.k8s.io/controller-runtime/pkg/client"
)

type KubernetesSecretsReader struct {
	client kclient.Client
}

// GetSecret returns the content of kubernetes secret.
func (s *KubernetesSecretsReader) GetSecret(ctx context.Context, secretName string,
	namespace string, log log.Logger) (map[string]interface{}, error) {
	// Read the secret
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      secretName,
		},
	}
	objectKey := kclient.ObjectKeyFromObject(secret)

	// Read the secret.
	err := s.client.Get(ctx, objectKey, secret)
	if err != nil {
		log.Error("Error in GetSecret: " + err.Error())
		return nil, err
	}

	data := make(map[string]interface{})
	for key, value := range secret.Data {
		data[key] = value
	}

	return data, nil
}
