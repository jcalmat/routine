PACKAGE		= routine

GO			= go
GOLINT		= golangci-lint
GODOC		= godoc
GOFMT		= gofmt

V			= 0
Q			= $(if $(filter 1,$V),,@)
M			= $(shell printf "\033[0;35m▶\033[0m")

.PHONY: all
all: check ## Run check

.PHONY: vendor
vendor: ## Get dependencies without updating .lock file
	$(info $(M) running mod vendor…) @
	$Q $(GO) mod vendor

.PHONY: clean
clean: ## Remove vendor folder
	$(info $(M) cleaning vendor…) 
	$Q rm -rf vendor

# Lint
.PHONY: lint
lint: ## Run linter check on project
	$(info $(M) running $(GOLINT)…)
	$Q $(GOLINT) run

# Test
.PHONY: test
test: ## Run unit tests
	$(info $(M) running go test…) @
	$Q $(GO) test -race -v -coverprofile .coverage.txt ./...

.PHONY: cover
cover: test ## Generate global code coverage report
	$Q $(GO) tool cover -func .coverage.txt
	@rm .coverage.txt

# Check
.PHONY: check
check: vendor lint cover ## Run complete check

.PHONY: fmt
fmt: ## Format code
	$(info $(M) running $(GOFMT)…) @
	$Q $(GOFMT) ./...