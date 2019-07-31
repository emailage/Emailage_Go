GO ?= go
GOFLAGS = CGO_ENABLED=0
LINTER ?= golint

BINDIR := bin
BINARY := tarantula
APIPPROFDIR := apipprof
DBPPROFDIR := dbpprof

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
	rm -f $(APIPPROFDIR)/*.prof
	rm -f $(DBPPROFDIR)/*.prof

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
bench:
	$(GO) test -bench=By100 -benchmem -benchtime=300ms -run=XXX ./database -blockprofile=block.prof -cpuprofile=cpu.prof -memprofile=mem.prof -mutexprofile=mutex.prof
	mv ./*.prof ./dbpprof

	$(GO) test -bench=Translate -benchmem -benchtime=300ms -run=XXX ./api -blockprofile=block.prof -cpuprofile=cpu.prof -memprofile=mem.prof -mutexprofile=mutex.prof
	mv ./*.prof ./apipprof

# use the new google UI for pprof
# requires installation via:
#
# go get github.com/google/pprof
#
.PHONY:
db_cpu_prof:
	pprof -http ":8888" dbpprof/cpu.prof

.PHONY:
api_cpu_prof:
	pprof -http ":8888" apipprof/cpu.prof
