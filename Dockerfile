FROM alpine:latest

WORKDIR /
COPY vault/plugins/vault-plugin-secrets-kubernetes-reader .

USER 10001

