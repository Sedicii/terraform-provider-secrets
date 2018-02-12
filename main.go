package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/sedicii/terraform-provider-secrets/secrets"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: secrets.Provider})
}
