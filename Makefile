PROVIDER_VERSION := "0.0.1"
OS := $(shell echo `uname` | tr '[:upper:]' '[:lower:]')
ARCH := $(shell arch)

.PHONY: all
all: clean gen fmt build vet test

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
gen-tf-code: clean
	@go run ./hack/gen-tf-code/...
	@go fmt ./pkg/schemas/...
	@~/go/bin/goimports -w ./pkg/schemas

.PHONY: gen
gen: gen-tf-code

.PHONY: build
build: gen
	@CGO_ENABLED=0 go build -ldflags="-s -w" ./cmd/terraform-provider-kops

.PHONY: fmt
fmt: build
	@go fmt ./cmd/...
	@go fmt ./pkg/...

.PHONY: test
test: fmt
	@go test ./...

.PHONY: vet
vet: fmt
	@go vet ./...

.PHONY: install
install: all
	@mkdir -p ${HOME}/.terraform.d/plugins/github/eddycharly/kops/${PROVIDER_VERSION}/${OS}_${ARCH}
	@cp terraform-provider-kops $(HOME)/.terraform.d/plugins/github/eddycharly/kops/${PROVIDER_VERSION}/${OS}_${ARCH}/terraform-provider-kops

# EXAMPLES FOR TERRAFORM < 0.15

.PHONY: examples
examples: example-basic example-aws-profile example-aws-assume-role example-bastion example-klog

.PHONY: example-basic
example-basic: install
	@terraform init 		./examples/basic
	@terraform validate ./examples/basic
	@terraform plan 		./examples/basic

.PHONY: example-aws-profile
example-aws-profile: install
	@terraform init 		./examples/aws-profile
	@terraform validate ./examples/aws-profile
	@terraform plan 		./examples/aws-profile

.PHONY: example-aws-assume-role
example-aws-assume-role: install
	@terraform init 		./examples/aws-assume-role
	@terraform validate ./examples/aws-assume-role

.PHONY: example-bastion
example-bastion: install
	@terraform init 		./examples/bastion
	@terraform validate ./examples/bastion
	@terraform plan 		./examples/bastion

.PHONY: example-klog
example-klog: install
	@terraform init 		./examples/klog
	@terraform validate ./examples/klog
	@terraform plan 		./examples/klog

# EXAMPLES FOR TERRAFORM >= 0.15

.PHONY: examples-15
examples-15: example-15-basic example-15-aws-profile example-15-aws-assume-role example-15-bastion example-15-klog

.PHONY: example-15-basic
example-15-basic: install
	@terraform -chdir=./examples/basic init
	@terraform -chdir=./examples/basic validate
	@terraform -chdir=./examples/basic plan

.PHONY: example-15-aws-profile
example-15-aws-profile: install
	@terraform -chdir=./examples/aws-profile init
	@terraform -chdir=./examples/aws-profile validate
	@terraform -chdir=./examples/aws-profile plan

.PHONY: example-15-aws-assume-role
example-15-aws-assume-role: install
	@terraform -chdir=./examples/aws-assume-role init
	@terraform -chdir=./examples/aws-assume-role validate

.PHONY: example-15-bastion
example-15-bastion: install
	@terraform -chdir=./examples/bastion init
	@terraform -chdir=./examples/bastion validate
	@terraform -chdir=./examples/bastion plan

.PHONY: example-15-klog
example-15-klog: install
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
	@terraform init 									./tests/basic
	@terraform validate 							./tests/basic
	@terraform plan 									./tests/basic
	@terraform apply  -auto-approve 	./tests/basic

.PHONY: integration-external-policies
integration-external-policies: integration-reset
	@terraform init 									./tests/external-policies
	@terraform validate 							./tests/external-policies
	@terraform plan 									./tests/external-policies
	@terraform apply  -auto-approve 	./tests/external-policies
