package stripe

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type stripeConfig struct {
	APIKey *string `hcl:"api_key"`
}

func ConfigInstance() interface{} {
	return &stripeConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) stripeConfig {
	if connection == nil || connection.Config == nil {
		return stripeConfig{}
	}
	config, _ := connection.Config.(stripeConfig)
	return config
}
