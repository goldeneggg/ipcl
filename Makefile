GO ?= go
BINNAME := ipcl

all: build

build: test
	$(GO) build -o $(GOBIN)/$(BINNAME)

test:
	$(GO) test ./parser
