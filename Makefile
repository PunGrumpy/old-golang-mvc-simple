.PHONY: test build clean
.SILENT: test build clean

BIN=$(CURDIR)/build
GO=$(shell which go)
NAME="simple-mvc"

build:
	$(GO) build -o $(BIN)/$(NAME) ./cmd/main.go

test:
	$(GO) test -race -cover -coverprofile=coverage.out -covermode=atomic ./...

clean:
	rm -rf $(BIN)
