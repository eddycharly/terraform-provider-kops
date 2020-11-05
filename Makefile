PROVIDER_VERSION := "0.0.1"

.PHONY: all
all: clean gen-code fmt build

.PHONY: clean
clean:
	@rm -f terraform-provider-kops
	@rm -rf ./pkg/schemas/*.generated.go
	@rm -rf ./pkg/structures/*.generated.go
	@rm -rf ./docs/*.generated.md

.PHONY: gen-code
gen-code:
	@go run ./hack/gen-structures/main.go
	@go fmt ./pkg/schemas/...
	@go fmt ./pkg/structures/...
	@~/go/bin/goimports -w ./pkg/schemas/
	@~/go/bin/goimports -w ./pkg/structures/

.PHONY: build
build:
	@CGO_ENABLED=0 go build ./cmd/terraform-provider-kops

.PHONY: fmt
fmt:
	@go fmt ./cmd/...
	@go fmt ./pkg/...

.PHONY: install
install: all
	@cp terraform-provider-kops $(HOME)/.terraform.d/plugins/github/eddycharly/kops/${PROVIDER_VERSION}/darwin_amd64/terraform-provider-kops

.PHONY: examples
examples: example-basic example-aws-profile example-bastion

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

.PHONY: example-bastion
example-bastion: install
	@terraform init ./examples/bastion
	@terraform validate ./examples/bastion
	@terraform plan ./examples/bastion

.PHONY: demo-1
demo-1:
	@terraform init ./demo-1
	@terraform validate ./demo-1
	@terraform apply ./demo-1

.PHONY: clean-demo-1
clean-demo-1:
	@terraform init ./demo-1
	@terraform destroy ./demo-1

.PHONY: demo-2
demo-2:
	@terraform init ./demo-2
	@terraform validate ./demo-2
	@terraform apply ./demo-2

.PHONY: clean-demo-2
clean-demo-2:
	@terraform init ./demo-2
	@terraform destroy ./demo-2
