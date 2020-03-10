GOCMD=$(GOROOT)/bin/go
GOBUILD=$(GOCMD) build
GOTOOL=$(GOCMD) tool
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

test:
	APP_ENV=test $(GOTEST) -coverprofile cover.html -v ./...
	$(MAKE) cover

codecov:
	$(GOTEST) -race -coverprofile=coverage.txt -covermode=atomic -v ./...

cover:
	$(GOTOOL) cover -html cover.html
