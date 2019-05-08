.PHONY: build

# binary to build
BINARIES ?= hum

# common variables
SOURCES :=	$(shell find . -type f -name "*.go")
GOPATH :=	$(shell if [ -n "${GOPATH}" ]; then echo ${PWD}:${GOPATH}; else pwd; fi)
GOENV ?=	GO15VENDOREXPERIMENT=1 GOPATH=$(GOPATH)
GO ?=		$(GOENV) go
GODEP ?=	$(GOENV) dep
USER ?=		$(shell whoami)

all:	build

build:	$(BINARIES)

$(BINARIES):	$(SOURCES)
	cd src/$@ && $(GODEP) ensure
	$(GO) build -o ./bin/$@ ./src/$@