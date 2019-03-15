package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-datadog/datadogpagerduty"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: datadogpagerduty.Provider})
}
