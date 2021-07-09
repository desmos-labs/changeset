lint:
	golangci-lint run --out-format=tab --timeout=10m

lint-fix:
	golangci-lint run --fix --out-format=tab --issues-exit-code=0 --timeout=10m
.PHONY: lint lint-fix

format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs goimports -w -local github.com/desmos-labs/changeset
.PHONY: format

###############################################################################
###                                  Build                                  ###
###############################################################################

BUILD_TARGETS := build install

build: BUILD_ARGS=-o $(BUILDDIR)/

build-linux: go.sum
	GOOS=linux GOARCH=amd64 LEDGER_ENABLED=false $(MAKE) build

build-arm32:go.sum
	GOOS=linux GOARCH=arm GOARM=7 LEDGER_ENABLED=false $(MAKE) build

build-arm64: go.sum
	GOOS=linux GOARCH=arm64 LEDGER_ENABLED=false $(MAKE) build

$(BUILD_TARGETS): go.sum $(BUILDDIR)/
	go $@ -mod=readonly $(BUILD_FLAGS) $(BUILD_ARGS) ./...

$(BUILDDIR)/:
	mkdir -p $(BUILDDIR)/