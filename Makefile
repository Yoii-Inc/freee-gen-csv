# メタ情報
NAME := freee-get-sales
VERSION := $(gobump show -r)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.revision=$(REVISION)'

export GO111MODULE=on

.PHONY: deps
deps:
	go mod tidy

# 必要なツール類をセットアップする
.PHONY: devel-deps
devel-deps: deps
	go install \
		golang.org/x/lint/golint@latest
	go install \
		github.com/x-motemen/gobump/cmd/gobump@latest
	go install \
		github.com/Songmu/make2help/cmd/make2help@latest

# テストを実行する
## Run tests
.PHONY: test
test: deps
	go test ./...

## Lint
.PHONY: lint
lint: devel-deps
	go vet ./...
	golint --set_exit_status./...

## build binaries ex. make bin/myproj
bin/%: cmd/%/main.go deps
	go build -ldflags "$(LDFLAGS)" -o $@ $<

.PHONY: build
build: bin/gen-sales-csv

## Show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)

