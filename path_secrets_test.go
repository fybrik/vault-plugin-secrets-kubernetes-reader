package kubesecrets

import (
	"context"
	"testing"

	"github.com/hashicorp/vault/sdk/logical"

	"github.com/hashicorp/go-hclog"
)

func getTestBackend(t *testing.T) logical.Backend {
	b, _ := newBackend()

	c := &logical.BackendConfig{
		Logger: hclog.New(&hclog.LoggerOptions{}),
	}
	err := b.Setup(context.Background(), c)
	if err != nil {
		t.Fatalf("unable to create backend: %v", err)
	}
	return b
}

func TestSecretNamespaceMissing(t *testing.T) {

}
