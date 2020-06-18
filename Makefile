OS_NAME := $(shell uname -s | tr A-Z a-z)
OS_ARCH := $(shell uname -m)
REVISION ?= $(shell git rev-parse --short HEAD)

GOOS :=
GOARCH :=
ifeq ($(OS_NAME),darwin)
	GOOS = darwin
else
	GOOS = linux
endif
ifeq ($(OS_ARCH),x86_64)
	GOARCH = amd64
else
	GOARCH = 386
endif

# Go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

all: test

clean:
	$(GOCLEAN) ./...

test: clean
	$(GOTEST) -v ./...
