package kubesecrets

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	kclient "sigs.k8s.io/controller-runtime/pkg/client"
	kconfig "sigs.k8s.io/controller-runtime/pkg/client/config"
)

// backend wraps the backend framework
type secretsReaderBackend struct {
	*framework.Backend

	KubeSecretReader KubernetesSecretsReader
}

var _ logical.Factory = Factory

// Factory configures and returns the plugin backends
func Factory(ctx context.Context, conf *logical.BackendConfig) (logical.Backend, error) {
	b, err := newBackend()
	if err != nil {
		return nil, err
	}

	if conf == nil {
		return nil, fmt.Errorf("configuration passed into backend is nil")
	}

	if err := b.Setup(ctx, conf); err != nil {
		return nil, err
	}

	return b, nil
}

func newBackend() (*secretsReaderBackend, error) {
	scheme := runtime.NewScheme()
	err := clientgoscheme.AddToScheme(scheme)
	if err != nil {
		return nil, err
	}

	// TODO: support configuration where Vault installed out of cluster
	client, err := kclient.New(kconfig.GetConfigOrDie(), kclient.Options{Scheme: scheme})
	if err != nil {
		return nil, err
	}

	b := &secretsReaderBackend{
		KubeSecretReader: KubernetesSecretsReader{
			client: client,
		},
	}

	b.Backend = &framework.Backend{
		Help: strings.TrimSpace(backendHelp),
		// TypeLogical indicates that the backend (plugin) is a secret provider.
		BackendType: logical.TypeLogical,
		// Define the path for which this backend will respond.
		Paths: []*framework.Path{
			pathSecrets(b),
		},
	}

	return b, nil
}
