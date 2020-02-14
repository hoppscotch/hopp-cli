PREFIX?=/usr/local
BINDIR?=$(_INSTDIR)/bin
VERSION?=$(shell git tag | grep ^v | sort -V | tail -n 1)
GOFLAGS?=-ldflags '-X main.VERSION=${VERSION}'

pwcli: cli.go go.mod go.sum
	@echo
	@echo Building pwcli. This may take a minute or two.
	@echo
	go build $(GOFLAGS) -o $@
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
	@echo Installing pwcli...
	@echo
	install -m755 pwcli $(BINDIR)
	@echo
	@echo ...Done\!

.PHONY: uninstall
uninstall:
	@echo
	@echo Uninstalling pwcli...
	@echo
	rm -f $(BINDIR)/pwcli
	@echo
	@echo ...Done\!
