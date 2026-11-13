package stripe

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type stripeConfig struct {
	APIKey *string `hcl:"api_key"`
}

func ConfigInstance() interface{} {
	return &stripeConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) stripeConfig {
	if connection == nil || connection.GetConfig() == nil {
		return stripeConfig{}
	}
	config, _ := connection.GetConfig().(stripeConfig)
	return config
}
