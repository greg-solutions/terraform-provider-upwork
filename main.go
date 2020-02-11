package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/vadimDidenko/terraform-provider-upwork/upwork"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: upwork.Provider})
}
