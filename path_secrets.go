package kubesecrets

import (
	"context"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

const secretsPrefix = "secret_name"

func pathSecrets(b *secretsReaderBackend) *framework.Path {
	return &framework.Path{
		Pattern: framework.MatchAllRegex(secretsPrefix),

		Fields: map[string]*framework.FieldSchema{
			"secret_name": {
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
		ExistenceCheck:  b.handleExistenceCheck,
	}
}

func (b *secretsReaderBackend) handleRead(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	secretName := data.Get("secret_name").(string)
	namespace := data.Get("namespace").(string)
	b.Logger().Info("In handleRead() secretName: " + secretName + ", namespace: " + namespace)

	if secretName == "" {
		resp := logical.ErrorResponse("Missing secret name")
		return resp, nil
	}

	if namespace == "" {
		resp := logical.ErrorResponse("Missing secret namespace")
		return resp, nil
	}

	fetchedData, err := b.KubeSecretReader.GetSecret(secretName, namespace, b.Logger())
	if err != nil {
		resp := logical.ErrorResponse("Error reading the secret data " + err.Error())
		return resp, nil
	}

	// Generate the response
	resp := &logical.Response{
		Data: fetchedData,
	}

	return resp, nil
}

func (b *secretsReaderBackend) handleExistenceCheck(ctx context.Context, req *logical.Request, data *framework.FieldData) (bool, error) {
	out, err := req.Storage.Get(ctx, req.Path)
	if err != nil {
		return false, errwrap.Wrapf("existence check failed: {{err}}", err)
	}

	return out != nil, nil
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