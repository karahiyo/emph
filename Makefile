CWD=$(shell pwd)
GOROOT:=
GOPATH:=$(shell pwd)
PATH:=/bin:/usr/bin:/usr/local/bin:$(CWD)/bin

env:
	echo $(GOPATH)
	go env

install:
	go get github.com/mattn/gom
	gom install

pkg:
	gom exec go build -o $(CWD)/bin/emph

test:
	gom test -v

integration_test:
	find $(CWD)/integration_test/log -type f | \
		xargs tail -f | \
		$(CWD)/bin/emph -c $(CWD)/integration_test/conf/sample

fmt:
	gofmt -w *.go

.PHONY: pkg integration_test
