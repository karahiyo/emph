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

run:
	gom run emph.go -c "/Users/yu_ke/workspace/emph/t/conf/sample"

test:
	gom test -v

fmt:
	gofmt -w *.go

.PHONY: pkg
