go build: rpg

all: clean dependencies  build

linux: clean dependencies  build-linux

clean:
	@echo "Cleanning..."
	-rm -f $(BINARY_NAME)
	-rm -f $(BINARY_UNIX)
	-find . -type d -name -exec rm -rf \{} +
	-$(GOCLEAN) -i
	@echo "Done cleanning."


dependencies:
	@echo "Getting dependencies..."
	@echo "Dns Servers : "
	@cat /etc/resolv.conf | grep nameserver
	$(GOMOD) tidy
	$(GOMOD) download
	$(GOMOD) vendor
	@echo "Done getting dependencies."

rpg:
	@echo "Building RPG"
	@echo "Building..."
	$(GOBUILD) ./...
	$(GOBUILD) -o $(BINARY_NAME) -v

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOMOD=GO111MODULE=on $(GOCMD) mod
BINARY_NAME=rpg
BINARY_UNIX=$(BINARY_NAME)
