GOCMD=/usr/local/bin/go
GOBUILD=$(GOCMD) build
GOTOOL=$(GOCMD) tool
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

test:
	$(GOTEST) -coverprofile cover.html -v ./...
	$(MAKE) cover

cover:
	$(GOTOOL) cover -html cover.html
