GOBUILD=go build
BINDIR=/usr/bin
CFGDIR=/etc/opt/ibapi

default: bin/ibapi bin/ibapi_cgo

all: install bin/ibapi_cgo

install: $(BINDIR)/ibapi $(CFGDIR)/ibapi.conf

bin/ibapi: *.go go.mod go.sum
	CGO_ENABLED=0 $(GOBUILD) -o bin/ibapi *.go
	strip bin/ibapi

bin/ibapi_cgo: *.go go.mod go.sum
	CGO_ENABLED=1 $(GOBUILD) -o bin/ibapi_cgo *.go
	strip bin/ibapi_cgo

$(BINDIR)/ibapi: bin/ibapi
	sudo cp -p bin/ibapi /usr/bin

$(CFGDIR)/ibapi.conf: ibapi.conf
	sudo mkdir -p $(CFGDIR)
	sudo cp -p ibapi.conf $(CFGDIR)


