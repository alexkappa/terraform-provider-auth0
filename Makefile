PKGS ?= $(shell go list ./...)
TESTS ?= ".*"
COVERS ?= "c.out"
PKG_OS ?= darwin linux
PKG_ARCH ?= amd64
BASE_PATH ?= $(shell pwd)
BUILD_PATH := $(BASE_PATH)/build
PROVIDER := terraform-provider-auth0

clean:
	@rm -rf $(BUILD_PATH)

build:
	@go build $(PKGS)

install: build
	@go install

test:
	@go test $(PKGS)

testacc:
	@TF_ACC=1 go test $(PKGS) -v -coverprofile=$(COVERS) -run ^$(TESTS)$

packages:
	@for os in $(PKG_OS); do \
		for arch in $(PKG_ARCH); do \
			mkdir -p $(BUILD_PATH)/$(PROVIDER)_$${os}_$${arch} && \
				cd $(BASE_PATH) && \
				GOOS=$${os} GOARCH=$${arch} CGO_ENABLED=0 go build -o $(BUILD_PATH)/$(PROVIDER)_$${os}_$${arch}/$(PROVIDER) . && \
				cd $(BUILD_PATH) && \
				tar -cvzf $(BUILD_PATH)/$(PROVIDER)_$${os}_$${arch}.tar.gz $(PROVIDER)_$${os}_$${arch}/; \
		done; \
	done;
