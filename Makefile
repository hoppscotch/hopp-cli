BIN:=hopp.bin
BINSTUFFED:=hopp-cli
PREFIX?=/usr/local
BINDIR?=$(PREFIX)/bin
VERSION?=$(shell git tag | grep ^v | sort -V | tail -n 1)
STATIC := ./templates/index.html ./templates/template.md:/template.md

all: deps build

.PHONY: deps
deps:
	@printf '%s\n' 'Checking binary dependencies...'
	@if [ ! -f "$(shell go env GOPATH)/bin/stuffbin" ]; then printf '%s\n' 'Installing stuffbin to GOPATH' && go install github.com/knadh/stuffbin/stuffbin@latest; fi

.PHONY: build
build:
	@printf '%s\n' 'Building hopp-cli. This may take a minute or two.'
	go build -o ${BIN} -ldflags="-s -w -X 'main.buildVersion=${VERSION}'"
	$(shell go env GOPATH)/bin/stuffbin -a stuff -in ${BIN} -out ${BINSTUFFED} ${STATIC}
	rm ${BIN}

.PHONY: clean
clean:
	@printf '%s\n' 'Cleaning...'
	@go clean
	@if [ -f 'hopp-cli' ]; then rm hopp-cli; fi
	@if [ -f 'hopp.bin' ]; then rm hopp.bin; fi

.PHONY: install
install:
	@printf '%s\n' 'Installing hopp-cli...'
	install -m755 hopp-cli $(BINDIR)

.PHONY: uninstall
uninstall:
	@printf '%s\n' 'Uninstalling hopp-cli...'
	rm -f $(BINDIR)/hopp-cli

.PHONY: pack-releases
pack-releases:
	$(foreach var,$(RELEASE_BUILDS),stuffbin -a stuff -in ${var} -out ./${var} ${STATIC};)