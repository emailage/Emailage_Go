GO ?= go
GOFLAGS = CGO_ENABLED=0
LINTER ?= golint

BINDIR := _examples
BINARY := _examples

VERSION := 0.1.9
LDFLAGS = -ldflags "-X main.gitSHA=$(shell git rev-parse HEAD) -X main.version=$(VERSION) -X main.name=$(BINARY)"

OS := $(shell go env GOHOSTOS)

.PHONY:
$(BINDIR)/$(BINARY): clean
	if [ ! -d $(BINDIR) ]; then mkdir $(BINDIR); fi
	GOOS=$(OS) $(GOFLAGS) $(GO) build -v -o $(BINDIR)/$(BINARY) $(LDFLAGS)

.PHONY:
test:
	$(GO) test -v -cover ./...

.PHONY:
clean:
	$(GO) clean
	rm -f $(BINDIR)/$(BINARY)
	rm -f ./auth/dbpprof/*.prof
	rm -f ./dbpprof/*.prof

.PHONY:
lint:
	$(LINTER)

.PHONY:
cleanbench:
	rm -f ./auth/dbpprof/*.prof
	rm -f ./dbpprof/*.prof

.PHONY:
mkbench:
	if [ ! -d "./auth/dbpprof" ]; then mkdir ./auth/dbpprof; fi
	if [ ! -d "./dbpprof" ]; then mkdir ./dbpprof; fi

.PHONY:
bench_run:
	$(GO) test -bench=. -benchmem -benchtime=1s -run=XXX ./auth -blockprofile=block.prof -cpuprofile=cpu.prof -memprofile=mem.prof -mutexprofile=mutex.prof
	mv ./*.prof ./auth/dbpprof
	$(GO) test -bench=. -benchmem -benchtime=2s -run=XXX ./ -blockprofile=block.prof -cpuprofile=cpu.prof -memprofile=mem.prof -mutexprofile=mutex.prof
	mv ./*.prof ./dbpprof

.PHONY:
bench: | mkbench cleanbench bench_run

# use the new google UI for pprof
# requires installation via:
#
# go get github.com/google/pprof
#
.PHONY:
auth_cpu_prof:
	pprof -http ":8888" auth/dbpprof/cpu.prof

.PHONY:
emailage_cpu_prof:
	pprof -http ":8888" dbpprof/cpu.prof
