GOCMD=/usr/local/bin/go
GOBUILD=$(GOCMD) build
GOTOOL=$(GOCMD) tool
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

test:
	$(GOTEST) -coverprofile cover.html -v ./...
	$(MAKE) cover

codecov:
	$(GOTEST) -race -coverprofile=coverage.txt -covermode=atomic -v ./...

cover:
	$(GOTOOL) cover -html cover.html
