BIN_DIR := $(GOPATH)/bin
LINTER := $(BIN_DIR)/golangci_lint
FRESH := $(BIN_DIR)/fresh

$(LINTER):
	cd $(GOPATH)
	go get golang.org/x/lint/golint

.PHONY: lint
lint: $(LINTER)
	golint

$(FRESH):
	cd $(GOPATH)
	go get github.com/pilu/fresh

.PHONY: fresh
fresh: $(FRESH)
	fresh
