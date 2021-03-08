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
