package main

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

var (
	version = "dev"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: provider.NewProvider(version)})
}
