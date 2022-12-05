package ripeatlas

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/schema"
)

type ripeatlasConfig struct {
	Key *string `cty:"key"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"key": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &ripeatlasConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) ripeatlasConfig {
	if connection == nil || connection.Config == nil {
		return ripeatlasConfig{}
	}
	config, _ := connection.Config.(ripeatlasConfig)
	return config
}
