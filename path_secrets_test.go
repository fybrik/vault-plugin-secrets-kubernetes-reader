package kubesecrets

import (
	"context"
	"testing"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/onsi/gomega"
)

func TestSecretNamespaceMissing(t *testing.T) {
	t.Parallel()
	g := gomega.NewGomegaWithT(t)

	// create new backend
	b, err := newBackend()
	g.Expect(err).To(gomega.BeNil())
	c := &logical.BackendConfig{
		Logger: hclog.New(&hclog.LoggerOptions{}),
	}
	err = b.Setup(context.Background(), c)
	g.Expect(err).To(gomega.BeNil(), "unable to create backend")

	request := &logical.Request{
		Operation: logical.ReadOperation,
		Path:      "dummy-secret",
		Data:      make(map[string]interface{}),
	}

	errMsg := "Missing secret namespace"
	_, err = b.HandleRequest(context.Background(), request)
	g.Expect(err.Error()).Should(gomega.Equal(err.Error()), errMsg)
}

func TestSecretExists(t *testing.T) {
	t.Parallel()
	g := gomega.NewGomegaWithT(t)

	// create new backend
	b, err := newBackend()
	g.Expect(err).To(gomega.BeNil())
	c := &logical.BackendConfig{
		Logger: hclog.New(&hclog.LoggerOptions{}),
	}
	err = b.Setup(context.Background(), c)
	g.Expect(err).To(gomega.BeNil(), "unable to create backend")

	data := make(map[string]interface{})
	data["namespace"] = "default"
	request := &logical.Request{
		Operation: logical.ReadOperation,
		Path:      "dummy-secret",
		Data:      data,
	}

	_, err = b.HandleRequest(context.Background(), request)
	g.Expect(err).To(gomega.BeNil())
}
