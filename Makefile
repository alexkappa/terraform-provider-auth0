PKGS ?= $(shell go list ./...)

build:
	go build

install: build
	go install

test:
	go test $(PKGS) -v

testacc:
	TF_ACC=1 go test $(PKGS) -v
