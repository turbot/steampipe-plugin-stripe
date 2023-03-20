package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-stripe/stripe"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: stripe.Plugin})
}
