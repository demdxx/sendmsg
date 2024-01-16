LIB_TAGS ?=

.PHONY: lint
lint: golint ## Run linters

.PHONY: golint
golint:
	# golint -set_exit_status ./...
	golangci-lint run -v ./...

.PHONY: fmt
fmt: ## Run formatting code
	@echo "Fix formatting"
	@gofmt -w ${GO_FMT_FLAGS} $$(go list -f "{{ .Dir }}" ./...); if [ "$${errors}" != "" ]; then echo "$${errors}"; fi

.PHONY: test
test: ## Run unit tests
	go test -v -tags "${LIB_TAGS}" -race ./...

.PHONY: tidy
tidy: ## Run go mod tidy
	go mod tidy
