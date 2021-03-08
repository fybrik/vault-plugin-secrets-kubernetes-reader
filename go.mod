module github.com/hashicorp/vault-plugin-kubernetes-secrets

go 1.13

require (
	github.com/hashicorp/go-hclog v0.15.0
	github.com/hashicorp/vault/api v1.0.4
	github.com/hashicorp/vault/sdk v0.1.13
	k8s.io/api v0.20.4
	k8s.io/apimachinery v0.20.4
	k8s.io/client-go v0.20.2
	sigs.k8s.io/controller-runtime v0.8.2
)
