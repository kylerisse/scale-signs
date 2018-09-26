BINARY := signs
VERSION := 0.1
BUILDOS := "darwin" # builder OS [darwin or linux]
BIN_DIR := $(GOPATH)/bin
GOMETALINTER := $(BIN_DIR)/gometalinter

.PHONY: info
info:
	# please use make test or make build

# TODO eval broken in linux, build uses os specific parameter for tar [-s darwin / --transform linux]
.PHONY: build
build: test clean releases resources
	# make: build
	@if [[ $(BUILDOS) =~ .*darwin.* ]]; then\
		tar -cvzf $(BINARY)-v$(VERSION).tar.gz -s /^out/$(BINARY)/ out/;\
	fi
	@if [[ $(BUILDOS) =~ .*linux.* ]]; then\
		tar -cvzf $(BINARY)-v$(VERSION).tar.gz --transform s/^out/$(BINARY)/ out/;\
	fi

.PHONY: test
test: lint
	# make: test
	go test --race -v ./...

.PHONY: lint
lint: $(GOMETALINTER)
	# make: lint
	gometalinter ./...

$(GOMETALINTER):
	go get -v -u github.com/alecthomas/gometalinter
	gometalinter --install

.PHONY: releases
releases: test
	# make: release
	mkdir -p out
	GOOS=linux GOARCH=amd64 go build -o out/$(BINARY)-v$(VERSION)-linux-amd64
	GOOS=darwin GOARCH=amd64 go build -o out/$(BINARY)-v$(VERSION)-darwin-amd64

.PHONY: resources
resources:
	# make: resources
	mkdir -p out
	cp -r client out/client
	cp -r images out/images
	cp -r LICENSE.md out/

.PHONY: clean
clean:
	# make: clean
	rm -rf out

.PHONY: mrproper
mrproper:
	# make: mrproper
	rm -rf out
	rm -f $(BINARY)-v*.tar.gz
