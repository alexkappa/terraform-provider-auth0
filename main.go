package main

import (
	"github.com/terraform-providers/terraform-provider-auth0/auth0"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return auth0.Provider()
		},
	})
}
