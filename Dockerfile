FROM alpine:latest

WORKDIR /
COPY ./vault/plugins/vault-plugin-secrets-kubernetes-reader .
COPY ./startup_plugin.sh .

