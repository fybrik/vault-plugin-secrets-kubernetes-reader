module github.com/hashicorp/vault-plugin-kubernetes-secrets

go 1.13

require (
	github.com/hashicorp/errwrap v1.0.0
	github.com/hashicorp/go-hclog v0.9.2
	github.com/hashicorp/vault/api v1.0.2
	github.com/hashicorp/vault/sdk v0.1.11
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
	k8s.io/api v0.20.4
	k8s.io/apimachinery v0.20.4
	k8s.io/client-go v0.20.2
	sigs.k8s.io/controller-runtime v0.8.2
)
