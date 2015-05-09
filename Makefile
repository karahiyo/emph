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

package:
	gom exec go build -o $(CWD)/package/emph

run:
	gom run emph.go

test:
	gom test -v

fmt:
	gofmt -w *.go

