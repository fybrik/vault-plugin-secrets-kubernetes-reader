package kubesecrets

import (
	"context"
	"fmt"
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
	b := getTestBackend(t)

	request := &logical.Request{
		Operation: logical.ReadOperation,
		Path:      fmt.Sprintf("%s/", secretsPrefix),
		Data:      make(map[string]interface{}),
	}

	fmt.Printf("REVIT")
	errMsg := "Missing secret namespace"
	resp, _ := b.HandleRequest(context.Background(), request)
	if resp.Error().Error() != errMsg {
		t.Errorf("Error must be '%s', get '%s'", errMsg, resp.Error())
	}
}
