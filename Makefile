PROVIDER_VERSION := "0.0.1"

.PHONY: all
all: clean gen-code fmt build

.PHONY: clean
clean:
	@rm -f terraform-provider-kops
	@rm -rf ./pkg/structures/*.generated.go

.PHONY: gen-code
gen-code:
	@go run ./hack/gen-structures/main.go
	@go fmt ./pkg/structures/...
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
