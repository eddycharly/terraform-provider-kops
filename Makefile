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
	@rm -rf ./pkg/schemas/utils
	@rm -rf ./docs/data-sources/*.md
	@rm -rf ./docs/provider-config/*.md
	@rm -rf ./docs/resources/*.md

.PHONY: gen-tf-code
gen-tf-code:
	@go run ./hack/gen-tf-code/main.go ./hack/gen-tf-code/docs.go
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
examples: example-basic example-aws-profile example-aws-assume-role example-bastion example-klog

.PHONY: example-basic
example-basic: install
	@terraform init 		${TF_ARGS} ./examples/basic
	@terraform validate ${TF_ARGS} ./examples/basic
	@terraform plan 		${TF_ARGS} ./examples/basic

.PHONY: example-aws-profile
example-aws-profile: install
	@terraform init 		${TF_ARGS} ./examples/aws-profile
	@terraform validate ${TF_ARGS} ./examples/aws-profile
	@terraform plan 		${TF_ARGS} ./examples/aws-profile

.PHONY: example-aws-assume-role
example-aws-assume-role: install
	@terraform init 		${TF_ARGS} ./examples/aws-assume-role
	@terraform validate ${TF_ARGS} ./examples/aws-assume-role

.PHONY: example-bastion
example-bastion: install
	@terraform init 		${TF_ARGS} ./examples/bastion
	@terraform validate ${TF_ARGS} ./examples/bastion
	@terraform plan 		${TF_ARGS} ./examples/bastion

.PHONY: example-klog
example-klog: install
	@terraform init 		${TF_ARGS} ./examples/klog
	@terraform validate ${TF_ARGS} ./examples/klog
	@terraform plan			${TF_ARGS} ./examples/klog
