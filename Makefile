PROVIDER_VERSION          := "0.0.1"
OS                        := $(shell echo `uname` | tr '[:upper:]' '[:lower:]')
TOOLS_DIR                 := $(PWD)/.tools
GOIMPORTS                 := $(TOOLS_DIR)/goimports
GOIMPORTS_VERSION         := latest
GOARCH                    ?= amd64

.PHONY: all
all: clean gen fmt build verify-gen vet test

.PHONY: clean
clean:
	@rm -f terraform-provider-kops
	@rm -rf ./pkg/schemas/config
	@rm -rf ./pkg/schemas/datasources
	@rm -rf ./pkg/schemas/kops
	@rm -rf ./pkg/schemas/kube
	@rm -rf ./pkg/schemas/resources
	@rm -rf ./pkg/schemas/utils
	@rm -rf ./docs/data-sources/*.md
	@rm -rf ./docs/provider-config/*.md
	@rm -rf ./docs/resources/*.md

$(GOIMPORTS):
	@echo Install goimports... >&2
	@GOBIN=$(TOOLS_DIR) go install golang.org/x/tools/cmd/goimports@$(GOIMPORTS_VERSION)

.PHONY: gen-tf-code
gen-tf-code: clean $(GOIMPORTS)
	@go run ./hack/gen-tf-code/...
	@go fmt ./pkg/schemas/...
	@$(GOIMPORTS) -w ./pkg/schemas

.PHONY: gen
gen: gen-tf-code

.PHONY: build
build: gen
	@CGO_ENABLED=0 go build -ldflags="-s -w -X 'github.com/eddycharly/terraform-provider-kops/pkg/version.BuildVersion=v${PROVIDER_VERSION}'" ./cmd/terraform-provider-kops

.PHONY: fmt
fmt: build
	@go fmt ./cmd/...
	@go fmt ./pkg/...

.PHONY: verify-gen
verify-gen: fmt
	@git --no-pager diff .
	@echo 'If this test fails, it is because the git diff is non-empty after running "make gen".' >&2
	@echo 'To correct this, locally run "make gen", commit the changes, and re-run tests.' >&2
	@git diff --quiet --exit-code .

.PHONY: test
test: fmt
	@go test ./...

.PHONY: vet
vet: fmt
	@go vet ./...

.PHONY: install
install: clean gen fmt build
	@mkdir -p ${HOME}/.terraform.d/plugins/github/eddycharly/kops/${PROVIDER_VERSION}/${OS}_${GOARCH}
	@cp terraform-provider-kops $(HOME)/.terraform.d/plugins/github/eddycharly/kops/${PROVIDER_VERSION}/${OS}_${GOARCH}/terraform-provider-kops

# EXAMPLES FOR TERRAFORM >= 0.15

.PHONY: examples
examples: example-basic example-aws-profile example-aws-assume-role example-bastion example-klog

.PHONY: example-basic
example-basic: install
	@terraform -chdir=./examples/basic init
	@terraform -chdir=./examples/basic validate
	@terraform -chdir=./examples/basic plan

.PHONY: example-aws-profile
example-aws-profile: install
	@terraform -chdir=./examples/aws-profile init
	@terraform -chdir=./examples/aws-profile validate
	@terraform -chdir=./examples/aws-profile plan

.PHONY: example-aws-assume-role
example-aws-assume-role: install
	@terraform -chdir=./examples/aws-assume-role init
	@terraform -chdir=./examples/aws-assume-role validate

.PHONY: example-bastion
example-bastion: install
	@terraform -chdir=./examples/bastion init
	@terraform -chdir=./examples/bastion validate
	@terraform -chdir=./examples/bastion plan

.PHONY: example-klog
example-klog: install
	@terraform -chdir=./examples/klog init
	@terraform -chdir=./examples/klog validate
	@terraform -chdir=./examples/klog plan

# INTEGRATION TESTS

.PHONY: integration
integration: integration-basic integration-external-policies

.PHONY: integration-reset
integration-reset:
	@rm -rf ./store
	@rm -f 	./terraform.tfstate

.PHONY: integration-basic
integration-basic: integration-reset
	@terraform -chdir=./tests/basic init
	@terraform -chdir=./tests/basic validate
	@terraform -chdir=./tests/basic plan
	@terraform -chdir=./tests/basic apply -auto-approve

.PHONY: integration-external-policies
integration-external-policies: integration-reset
	@terraform -chdir=./tests/external-policies init
	@terraform -chdir=./tests/external-policies validate
	@terraform -chdir=./tests/external-policies plan
	@terraform -chdir=./tests/external-policies apply -auto-approve
