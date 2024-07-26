BUF_VERSION=v1.34.0
SQL_BOILER_VERSION=v4.16.2
GO_LDFLAGS := -s -w -X github.com/averak/hbaas/app/core/build_info.serverVersion=$(shell git describe --tags --always)

.PHONY: install-tools
install-tools:
	go install github.com/bufbuild/buf/cmd/buf@${BUF_VERSION}
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/fdaines/arch-go@latest
	go install github.com/google/wire/cmd/wire@latest

BREAKING_CHANGE_BASE_BRANCH?=develop
.PHONY: lint
lint:
	 golangci-lint run --issues-exit-code=1 ./...
	 arch-go
	buf lint
	buf breaking --against '.git#branch=$(BREAKING_CHANGE_BASE_BRANCH)'

.PHONY: codegen
codegen:
	find . -type f \( -name 'wire_gen.go' \) -delete
	wire ./...
	find . -type f \( -name '*.connect.go' -or -name '*.pb.go' \) -delete
	buf generate

.PHONY: run-api-server
run-api-server:
	go run -ldflags="$(GO_LDFLAGS)" ./entrypoint/api_server
