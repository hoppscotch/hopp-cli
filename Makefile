BIN:=hopp.bin
PREFIX?=/usr/local
BINDIR?=$(PREFIX)/bin
VERSION?=$(shell git tag | grep ^v | sort -V | tail -n 1)
STATIC := ./templates/index.html ./templates/template.md:/template.md

deps:
	go get -u github.com/knadh/stuffbin/...

build: cli.go go.mod go.sum
	@echo
	@echo Building hopp-cli. This may take a minute or two.
	@echo
	go build -o ${BIN} -ldflags="-s -w -X 'main.buildVersion=${VERSION}'" *.go
	stuffbin -a stuff -in ${BIN} -out ${BIN} ${STATIC}
	@echo
	@echo ...Done\!

.PHONY: clean
clean:
	@echo
	@echo Cleaning...
	@echo
	go clean
	@echo
	@echo ...Done\!

.PHONY: update
update:
	@echo
	@echo Updating from upstream repository...
	@echo
	git pull --rebase origin master
	@echo
	@echo ...Done\!

.PHONY: install
install:
	@echo
	@echo Installing hopp-cli...
	@echo
	install -m755 hopp-cli $(BINDIR)
	@echo
	@echo ...Done\!

.PHONY: uninstall
uninstall:
	@echo
	@echo Uninstalling hopp-cli...
	@echo
	rm -f $(BINDIR)/hopp-cli
	@echo
	@echo ...Done\!
.PHONY: pack-releases
pack-releases:
	$(foreach var,$(RELEASE_BUILDS),stuffbin -a stuff -in ${var} -out ${var} ${STATIC};)