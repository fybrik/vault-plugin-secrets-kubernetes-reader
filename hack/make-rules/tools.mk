INSTALL_TOOLS += $(TOOLBIN)/kind
$(TOOLBIN)/kind:
	GOBIN=$(ABSTOOLBIN) go install sigs.k8s.io/kind@v0.11.1

INSTALL_TOOLS += $(TOOLBIN)/yq
.PHONY: $(TOOLBIN)/yq
$(TOOLBIN)/yq:
	cd $(TOOLS_DIR); ./install_yq.sh

INSTALL_TOOLS += $(TOOLBIN)/kubectl
.PHONY: $(TOOLBIN)/kubectl
$(TOOLBIN)/kubectl: $(TOOLBIN)/yq
	cd $(TOOLS_DIR); ./install_kubectl.sh

.PHONY: install-tools
install-tools: $(INSTALL_TOOLS)
	go mod tidy
	ls -l $(TOOLS_DIR)/bin

.PHONY: uninstall-tools
uninstall-tools:
	find $(TOOLBIN) -mindepth 1 -delete

