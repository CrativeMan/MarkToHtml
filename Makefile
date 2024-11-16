GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
SRCFOLDER = src
BINARYFOLDER = bin
BINARY_NAME=main

all:	build

build: 
	mkdir -p $(BINARYFOLDER)
	cd $(SRCFOLDER) && $(GOBUILD) -o ../$(BINARYFOLDER)/$(BINARY_NAME)

clean: 
	$(GOCLEAN)
	rm -rf $(BINARYFOLDER)

run:
	$(GOBUILD) -o $(BINARY_NAME)
	./$(BINARY_NAME)

mod:
	$(GOMOD) tidy