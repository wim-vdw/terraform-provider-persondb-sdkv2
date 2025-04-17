package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/wim-vdw/terraform-provider-myprovider/myprovider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: myprovider.Provider,
	})
}
