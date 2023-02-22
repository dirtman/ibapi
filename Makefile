GOBUILD=go build
BINDIR=/bin
CFGDIR=/etc/opt/ibapi

all: bin/ibapi bin/ibapi_cgo

install: bin/ibapi bin/ibapi_cgo \
	$(BINDIR)/ibapi $(CFGDIR)/ibapi.conf ibapi.conf

bin/ibapi: *.go go.mod go.sum
	$(GOBUILD) -o bin/ibapi *.go
	strip bin/ibapi

bin/ibapi_cgo: *.go go.mod go.sum
	CGO_ENABLED=0 $(GOBUILD) -o bin/ibapi_cgo *.go
	strip bin/ibapi_cgo

$(BINDIR)/ibapi: bin/ibapi
	sudo cp -p bin/ibapi /bin

$(CFGDIR)/ibapi.conf: ibapi.conf
	sudo mkdir -p $(CFGDIR)
	sudo cp -p ibapi.conf $(CFGDIR)

