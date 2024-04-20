package ripeatlas

import (
	"context"
	"errors"
	"os"

	"github.com/keltia/ripe-atlas"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*atlas.Client, error) {
	// get ripeatlas client from cache
	cacheKey := "ripeatlas"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*atlas.Client), nil
	}

	key := os.Getenv("RIPEATLAS_KEY")

	ripeatlasConfig := GetConfig(d.Connection)
	if ripeatlasConfig.Key != nil {
		key = *ripeatlasConfig.Key
	}

	if key == "" {
		return nil, errors.New("'key' must be set in the connection configuration. Edit your connection configuration file or set the RIPEATLAS_KEY environment variable and then restart Steampipe")
	}

	config := atlas.Config{
		APIKey: key,
	}
	client, err := atlas.NewClient(config)

	if err != nil {
		return nil, err
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}
