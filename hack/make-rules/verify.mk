GO_VERSION:=1.13
CODE_MAINT += go-version
.PHONY: go-version
go-version:
	@(go version | grep -q 'go$(GO_VERSION)\(\.[0-9]*\)\? ') || \
	echo 'WARNING: bad go version to fix run: eval "$$(gimme $(GO_VERSION))"'

CODE_MAINT += fmt
.PHONY: fmt
fmt:
	go fmt ./...

CODE_MAINT += vet
.PHONY: vet
vet:
	go vet ./...

CODE_MAINT += fix
.PHONY: fix
fix:
	go fix ./...

CODE_MAINT += tidy
.PHONY: tidy
tidy:
	go mod tidy

.PHONY: verify
verify: $(CODE_MAINT)
