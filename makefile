# Go parameters
GOCMD=go
GOGENERATE=$(GOCMD) generate
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=git-pull
BINARY_UNIX=$(BINARY_NAME)-linux-amd64
DOCKER_SLUG=soupstore/$(BINARY_NAME)

build:
	$(GOBUILD) -o bin/$(BINARY_NAME) -v
clean:
	$(GOCLEAN)
	rm -rdf bin
run:
	$(GOBUILD) -o bin/$(BINARY_NAME) -v
	bin/$(BINARY_NAME)

# Cross compilation
build-docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_UNIX) -v
	docker build -t $(DOCKER_SLUG):dev .
