GO ?= go
GODEP ?= godep
BINNAME := ipcl
PROFDIR := ./.profile
TESTTARGET := ./parser

all: build

build: test
	$(GO) build -o $(GOBIN)/$(BINNAME)

test: getdeps
	$(GO) test $(TESTTARGET)

getdeps:
	$(GO) get github.com/jessevdk/go-flags

proftest:
	[ ! -d $(PROFDIR) ] && mkdir $(PROFDIR); $(GO) test -bench . -benchmem -blockprofile $(PROFDIR)/block.out -cover -coverprofile $(PROFDIR)/cover.out -cpuprofile $(PROFDIR)/cpu.out -memprofile $(PROFDIR)/mem.out $(TESTTARGET)

depbuild: deptest
	$(GODEP) $(GO) build -o $(GOBIN)/$(BINNAME)

deptest: depsave
	$(GODEP) $(GO) test $(TESTTARGET)

depsave:
	$(GODEP) save ./...
