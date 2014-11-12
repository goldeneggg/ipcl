GO ?= go
GODEP ?= godep
PGMPKGPATH := .
BINNAME := ipcl
PROFDIR := ./.profile
TESTTARGET := ./lib/parser

all: build

build: test
	$(GO) build -o $(GOBIN)/$(BINNAME) $(PGMPKGPATH)

test: getdeps
	$(GO) test $(TESTTARGET)

# This target is a specification for "go get <your dependency libraries>"
# If you have some dependencies, you should remove comment and edit the following content
getdeps:
	$(GO) get github.com/jessevdk/go-flags

proftest: getdeps
	[ ! -d $(PROFDIR) ] && mkdir $(PROFDIR); $(GO) test -bench . -benchmem -blockprofile $(PROFDIR)/block.out -cover -coverprofile $(PROFDIR)/cover.out -cpuprofile $(PROFDIR)/cpu.out -memprofile $(PROFDIR)/mem.out $(TESTTARGET)

# Following targets using "godep"
depbuild: deptest
	$(GODEP) $(GO) build -o $(GOBIN)/$(BINNAME) $(PGMPKGPATH)

deptest: depsave
	$(GODEP) $(GO) test $(TESTTARGET)

depsave:
	$(GODEP) save ./...
