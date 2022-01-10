# Vault Kubenetes Secrets Plugin

This is a secret engine plugin for [HashiCorp Vault](https://www.vaultproject.io/) which reads Kuberentes secrets.

Requirements:

    make
    golang 1.13 and above
    Vault CLI utility

When Vault is deployed on Kuberentes cluster RBAC should be set to grant Vault with the proper permissions to read the secrets (Please see example/clusterrole.yaml and example/clusterrolebinding.yaml that can be used for such purpose).

## Quick Start

First of all we need to create a secret for testing:

```
$ kubectl create -f example/secret.yaml                    # A sample secret to read

```
Then build the plugin binary and start the Vault dev server:

```
make build
vault server -dev -dev-root-token-id=root -dev-plugin-dir=./vault/plugins
```

Now open a new terminal window and run the following commands:

```
# Open a new terminal window and export Vault dev server http address
$ export VAULT_ADDR='http://127.0.0.1:8200'

# Enable the the plugin
$ vault secrets enable -path=kubernetes-secrets-reader vault-plugin-secrets-kubernetes-reader 

# Read the sample secret:
$ vault read kubernetes-secrets-reader/my-secret namespace=default
Key         Value
---         -----
password    passw0rd
username    admin

```

## Gettings help

```
 vault path-help kubernetes-secrets-reader
```
