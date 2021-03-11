vault login root
vault secrets enable -path=kubernetes-secrets-reader vault-plugin-secrets-kubernetes-reader
vault policy write read-plugin-secrets - <<EOF
        path "kubernetes-secrets-reader/*" {
        capabilities = ["read"]
        }
EOF
