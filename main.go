package main

import (
	"github.com/andrecloutier/terraform-provider-datadogpagerduty/datadogpagerduty"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: datadogpagerduty.Provider})
}
