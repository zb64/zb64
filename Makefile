GOCMD=go
GOFMT1=gofmt
GOFMT2=goreturns
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GODOC=$(GOCMD) doc
GOGET=$(GOCMD) get
BINARY=zb64

all: fmt build test bench doc
ci: build test bench
doc:
	$(GODOC) -all .
fmt:
	$(GOFMT1) -s -w .
	$(GOFMT2) -l -w .
build:
	$(GOBUILD) -v
buildlinux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -v
buildwin:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -v
buildmac:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -v
test:
	$(GOTEST) -v -race -short -cover -covermode=atomic -coverprofile=coverage.out .
	cat coverage.out >> coverage.txt
testdev:
	$(GOTEST) -race -short -cover -covermode=atomic -count 1 .
bench:
	$(GOTEST) -parallel=4 -run="none" -benchtime="2s" -benchmem -bench=.
run: build
	./$(BINARY) -r "Hello World"
