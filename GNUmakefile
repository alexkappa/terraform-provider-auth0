PKG_NAME = auth0
PKGS ?= $$(go list ./...)
FILES ?= $$(find . -name '*.go' | grep -v vendor)
TESTS ?= ".*"
COVERS ?= "c.out"

default: build

build: fmtcheck
	@go mod edit -replace="gopkg.in/auth0.v5=github.com/Abacus-Insights/auth0@v1.3.1-0.20210512201735-a335ec727e5e"
	@go install

install: build
	@mkdir -p ~/.terraform.d/plugins
	@cp $(GOPATH)/bin/terraform-provider-auth0 ~/.terraform.d/plugins

sweep:
	@echo "WARNING: This will destroy infrastructure. Use only in development accounts."
	@go test ./auth0 -v -sweep="phony" $(SWEEPARGS)

test: fmtcheck
	@go test -i $(PKGS) || exit 1
	@echo $(PKGS) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4 -run ^$(TESTS)$

testacc: fmtcheck
	@TF_ACC=1 go test $(PKGS) -v $(TESTARGS) -timeout 120m -coverprofile=$(COVERS) -run ^$(TESTS)$

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	@gofmt -w $(FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

docgen:
	go run scripts/gendocs.go -resource auth0_<resource>

.PHONY: build test testacc vet fmt fmtcheck errcheck docgen
