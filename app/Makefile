APP := pants
BIN_DIR := $(GOPATH)/bin
LINTER := $(BIN_DIR)/golangci_lint
FRESH := $(BIN_DIR)/fresh
MODOFF := GO111MODULE=off
DB := $(APP).sqlite3

$(LINTER):
	cd $(GOPATH)
	$(MODOFF) go get golang.org/x/lint/golint

.PHONY: lint
lint: $(LINTER)
	golint

$(FRESH):
	cd $(GOPATH)
	$(MODOFF) go get github.com/pilu/fresh

.PHONY: fresh
fresh: $(FRESH)
	fresh

.PHONY: clean
clean:
	rm -r ./tmp $(DB) public/main.js $(APP)
	go clean

.PHONY: view
view:
	cp -r ../view/elm.json ../view/src .
	elm make src/Main.elm --output=public/main.js
	rm -r elm.json src elm-stuff

.PHONY: build
build:
	go build -o $(APP)

.PHONY: app
app: view build
