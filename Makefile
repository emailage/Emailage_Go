GO ?= go
GOFLAGS = CGO_ENABLED=0
LINTER ?= golint

BINDIR := _examples
BINARY := _examples

VERSION := 0.1.0
LDFLAGS = -ldflags "-X main.gitSHA=$(shell git rev-parse HEAD) -X main.version=$(VERSION) -X main.name=$(BINARY)"

OS := $(shell uname)

.PHONY:
$(BINDIR)/$(BINARY): clean
	if [ ! -d $(BINDIR) ]; then mkdir $(BINDIR); fi
ifeq ($(OS),Darwin)
	GOOS=darwin $(GOFLAGS) $(GO) build -v -o $(BINDIR)/$(BINARY) $(LDFLAGS)
endif
ifeq ($(OS),Linux)
	GOOS=linux $(GOFLAGS) $(GO) build -v -o $(BINDIR)/$(BINARY) $(LDFLAGS)
endif

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
docs:
	@godoc -http=:6060 2>/dev/null &
	@printf "To view tarantula docs, point your browser to:\n"
	@printf "\n\thttp://127.0.0.1:6060/pkg/eateam.visualstudio.com/$(BINARY)/$(pkg)\n\n"
	@sleep 1
	@open "http://127.0.0.1:6060/pkg/eateam.visualstudio.com/$(BINARY)/$(pkg)"

.PHONY:
lint:
	$(LINTER) `$(GO) list ./... | grep -v /vendor/`

.PHONY:
cleanbench:
	rm -f ./auth/dbpprof/*.prof
	rm -f ./dbpprof/*.prof

.PHONY:
mkbench:
	if [ ! -d "./auth/dbpprof" ]; then mkdir ./auth/dbpprof; fi
	if [ ! -d "./dbpprof"]; then mkdir ./dbpprof; fi

.PHONY:
bench_run:
	$(GO) test -bench=. -benchmem -benchtime=300ms -run=XXX ./auth -blockprofile=block.prof -cpuprofile=cpu.prof -memprofile=mem.prof -mutexprofile=mutex.prof
	mv ./*.prof ./auth/dbpprof
	$(GO) test -bench=. -benchmem -benchtime=300ms -run=XXX ./ -blockprofile=block.prof -cpuprofile=cpu.prof -memprofile=mem.prof -mutexprofile=mutex.prof
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
