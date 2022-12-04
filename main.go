package main

import (
	"github.com/francois2metz/steampipe-plugin-ripeatlas/ripeatlas"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: ripeatlas.Plugin})
}
