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

// GetSecret returns base64 decoded content of a kubernetes secret.
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
		// Kubernetes secrets are stored as unencrypted base64-encoded strings by default.
		// return the decoded values.

		data[key] = string(value)
		if err != nil {
			log.Error("Error in decoding secret: " + err.Error())
			return nil, err
		}
	}

	return data, nil
}
