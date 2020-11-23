PROVIDER_VERSION := "0.0.1"
OS := $(shell echo `uname` | tr '[:upper:]' '[:lower:]')

.PHONY: all
all: clean gen fmt build

.PHONY: clean
clean:
	@rm -f terraform-provider-kops
	@rm -rf ./pkg/schemas/config
	@rm -rf ./pkg/schemas/datasources
	@rm -rf ./pkg/schemas/kops
	@rm -rf ./pkg/schemas/kube
	@rm -rf ./pkg/schemas/resources
	@rm -rf ./docs/data-sources/*.md
	@rm -rf ./docs/provider-config/*.md
	@rm -rf ./docs/resources/*.md

.PHONY: gen-tf-code
gen-tf-code:
	@go run ./hack/gen-tf-code/main.go
	@go fmt ./pkg/schemas/...
	@~/go/bin/goimports -w ./pkg/schemas

.PHONY: gen
gen: gen-tf-code

.PHONY: build
build:
	@CGO_ENABLED=0 go build -ldflags="-s -w" ./cmd/terraform-provider-kops

.PHONY: fmt
fmt:
	@go fmt ./cmd/...
	@go fmt ./pkg/...

.PHONY: install
install: all
	@mkdir -p ${HOME}/.terraform.d/plugins/github/eddycharly/kops/${PROVIDER_VERSION}/${OS}_amd64
	@cp terraform-provider-kops $(HOME)/.terraform.d/plugins/github/eddycharly/kops/${PROVIDER_VERSION}/${OS}_amd64/terraform-provider-kops

.PHONY: examples
examples: example-basic example-aws-profile example-aws-assume-role example-bastion

.PHONY: example-basic
example-basic: install
	@terraform init ./examples/basic
	@terraform validate ./examples/basic
	@terraform plan ./examples/basic

.PHONY: example-aws-profile
example-aws-profile: install
	@terraform init ./examples/aws-profile
	@terraform validate ./examples/aws-profile
	@terraform plan ./examples/aws-profile

.PHONY: example-aws-assume-role
example-aws-assume-role: install
	@terraform init ./examples/aws-assume-role
	@terraform validate ./examples/aws-assume-role

.PHONY: example-bastion
example-bastion: install
	@terraform init ./examples/bastion
	@terraform validate ./examples/bastion
	@terraform plan ./examples/bastion
