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