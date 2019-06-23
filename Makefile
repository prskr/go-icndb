VERSION = $(shell git describe --dirty --tags --always)
REPO = github.com/baez90/go-icndb
BUILD_PATH = $(REPO)/cmd/icndb-server
PKGS = $(shell go list ./... | grep -v /vendor/)
TEST_PKGS = $(shell find . -type f -name "*_test.go" -printf '%h\n' | sort -u)
GOARGS = GOOS=linux GOARCH=amd64
GO_BUILD_ARGS = -ldflags="-w -s"
BINARY_NAME = icndb
DIR = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
DEBUG_PORT = 2345

GORELEASER_VERSION = 0.110.0

export CGO_ENABLED:=0

.PHONY: all
all: format compile

.PHONY: clean-all
clean-all: clean clean-vendor

.PHONY: rebuild
rebuild: clean format compile

.PHONY: format
format:
	@go fmt $(PKGS)

.PHONY: revive
revive: ensure-revive
	@revive --config $(DIR)assets/lint/config.toml -exclude $(DIR)vendor/... -formatter friendly $(DIR)...

.PHONY: clean
clean: ensure-packr2
	@rm -f debug $(BINARY_NAME)
	@rm -rf dist
	@packr2 clean

.PHONY: clean-vendor
clean-vendor:
	rm -rf vendor/

.PHONY: test
test:
	@go test -coverprofile=./cov-raw.out -v $(TEST_PKGS)
	@cat ./cov-raw.out | grep -v "generated" > ./cov.out

.PHONY: cli-cover-report
cli-cover-report:
	@go tool cover -func=cov.out

.PHONY: html-cover-report
html-cover-report:
	@go tool cover -html=cov.out -o .coverage.html

.PHONY: deps
deps:
	@go build -v ./...

.PHONY: compile
compile: deps ensure-packr2
	@$(GOARGS) packr2 build $(GO_BUILD_ARGS) -o $(DIR)/$(BINARY_NAME) $(BUILD_PATH)

.PHONY: docker
docker:
	@docker build --build-arg VCS_REF=$(VERSION) \
       			  --build-arg BUILD_DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ") \
     			  -t baez90/go-icndb:latest .

.PHONY: podman
podman:
	@podman build --build-arg VCS_REF=$(VERSION) \
       			  --build-arg BUILD_DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ") \
     			  -t baez90/go-icndb:latest .

.PHONY: run
run:
	@go run $(BUILD_PATH) \
		--host=127.0.0.1 \
	 	--port=8000 \
	 	--scheme=http

.PHONY: generate
generate: ensure-swagger
	@swagger generate server --target $(DIR) --spec $(DIR)assets/api/swagger.yml

.PHONY: debug
debug: ensure-delve
	@dlv debug \
		--headless \
		--listen=127.10.10.2:$(DEBUG_PORT) \
		--api-version=2 $(BUILD_PATH) \
		--build-flags="-tags debug" \
		-- --host=127.0.0.1 --port=8000 --scheme=http

.PHONY: watch
watch: ensure-reflex
	@reflex -r '\.go$$' -s -- sh -c 'make debug'

.PHONY: watch-test
watch-test: ensure-reflex
	@reflex -r '_test\.go$$' -s -- sh -c 'make test'

.PHONY: cloc
cloc:
	@cloc --vcs=git --exclude-dir=.idea,.vscode,.theia,public,docs, .

.PHONY: serve-godoc
serve-godoc: ensure-godoc
	@godoc -http=:6060

.PHONY: serve-docs
serve-docs: ensure-reflex docs
	@reflex -r '\.md$$' -s -- sh -c 'mdbook serve -d $(DIR)/public -n 127.0.0.1 $(DIR)/docs'

.PHONY: docs
docs:
	@mdbook build -d $(DIR)/public $(DIR)/docs`

.PHONY: test-release
test-release: ensure-goreleaser ensure-packr2
	@goreleaser --snapshot --skip-publish --rm-dist

.PHONY: ensure-swagger
ensure-swagger:
ifeq (, $(shell which swagger))
	$(shell go get -u github.com/go-swagger/go-swagger/cmd/swagger)
endif

.PHONY: ensure-revive
ensure-revive:
ifeq (, $(shell which revive))
	$(shell go get -u github.com/mgechev/revive)
endif

.PHONY: ensure-delve
ensure-delve:
ifeq (, $(shell which dlv))
	$(shell go get -u github.com/go-delve/delve/cmd/dlv)
endif

.PHONY: ensure-reflex
ensure-reflex:
ifeq (, $(shell which reflex))
	$(shell go get -u github.com/cespare/reflex)
endif

.PHONY: ensure-godoc
ensure-godoc:
ifeq (, $(shell which godoc))
	$(shell go get -u golang.org/x/tools/cmd/godoc)
endif

.PHONY: ensure-packr2
ensure-packr2:
ifeq (, $(shell which packr2))
	$(shell go get -u github.com/gobuffalo/packr/v2/packr2)
endif

.PHONY: ensure-goreleaser
ensure-goreleaser:
ifeq (, $(shell which goreleaser))
	$(shell curl -sL https://github.com/goreleaser/goreleaser/releases/download/v$(GORELEASER_VERSION)/goreleaser_Linux_x86_64.tar.gz | tar -xvz --exclude "*.md" -C $$GOPATH/bin)
endif