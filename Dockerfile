FROM golang:1.13.8-alpine as builder

WORKDIR /workspace
COPY . . 

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o vault/plugins/vault-plugin-secrets-kubernetes-reader cmd/vault-plugin-secrets-kubernetes-reader/main.go

EXPOSE 8080

# Use distroless as minimal base image to package the datauserserver binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /

COPY --from=builder /workspace/vault/plugins/vault-plugin-secrets-kubernetes-reader .
USER nonroot:nonroot

ENTRYPOINT ["/vault-plugin-secrets-kubernetes-reader"]

