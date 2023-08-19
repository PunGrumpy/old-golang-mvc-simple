.PHONY: build test clean
.SILENT: build test clean

BIN=$(CURDIR)/bin
GO=$(shell which go)
NAME="simple-mvc"

build:
	$(GO) build -o $(BIN)/$(NAME) cmd/main.go

test:
	$(GO) test -race -cover ./...

clean:
	rm -rf $(BIN)
