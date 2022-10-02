package kubesecrets

import (
	"context"
	"errors"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

const secretsPrefix = "secret_name"

// pathSecrets returns the path configuration for reading kubernetes secrets.
func pathSecrets(b *secretsReaderBackend) *framework.Path {
	return &framework.Path{
		Pattern: framework.MatchAllRegex(secretsPrefix),

		Fields: map[string]*framework.FieldSchema{
			secretsPrefix: {
				Type:        framework.TypeString,
				Description: "Specifies the name of the kubernetes secret.",
				Query:       true,
				Required:    true,
			},
			"namespace": {
				Type:        framework.TypeString,
				Description: "Specifies the name of the kubernetes secret namespace.",
				Query:       true,
				Required:    true,
			},
		},
		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.ReadOperation: b.handleRead,
		},
		HelpDescription: pathInvalidHelp,
	}
}

// handleRead handles a read request: it extracts the secret name and namespace
// and returns the secret content if no error occured.
func (b *secretsReaderBackend) handleRead(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	secretName := data.Get(secretsPrefix).(string)
	namespace := data.Get("namespace").(string)
	b.Logger().Info("In handleRead() secretName: " + secretName + ", namespace: " + namespace)

	if secretName == "" {
		resp := logical.ErrorResponse("missing secret name")
		return resp, errors.New("missing secret name")
	}

	if namespace == "" {
		resp := logical.ErrorResponse("missing secret namespace")
		return resp, errors.New("missing sceret namespace")
	}

	fetchedData, err := b.KubeSecretReader.GetSecret(ctx, secretName, namespace, b.Logger())
	if err != nil {
		resp := logical.ErrorResponse("Error reading the secret data: " + err.Error())
		return resp, err
	}

	// Generate the response
	resp := &logical.Response{
		Data: fetchedData,
	}

	return resp, nil
}

var backendHelp string = `
This backend reads kubernetes secrets.`

var pathInvalidHelp string = backendHelp + `

## PATHS

The following paths are supported by this backend. To view help for
any of the paths below, use the help command with any route matching
the path pattern. Note that depending on the policy of your auth token,
you may or may not be able to access certain paths.

{{range .Paths}}{{indent 4 .Path}}
{{indent 8 .Help}}

{{end}}
`
