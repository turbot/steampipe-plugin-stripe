package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-stripe/stripe"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: stripe.Plugin})
}
