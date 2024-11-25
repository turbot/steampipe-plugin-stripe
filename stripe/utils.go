package stripe

import (
	"context"
	"errors"
	"os"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/client"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(_ context.Context, d *plugin.QueryData) (*client.API, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "stripe"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*client.API), nil
	}

	// Define our app information
	stripe.SetAppInfo(&stripe.AppInfo{
		Name: "Steampipe",
		URL:  "https://hub.steampipe.io/plugins/turbot/stripe",
	})

	// Default to using env vars
	apiKey := os.Getenv("STRIPE_API_KEY")

	// But prefer the config
	stripeConfig := GetConfig(d.Connection)
	if stripeConfig.APIKey != nil {
		apiKey = *stripeConfig.APIKey
	}

	if apiKey == "" {
		// Credentials not set
		return nil, errors.New("api_key must be configured")
	}

	maxRetries := int64(10)
	config := &stripe.BackendConfig{
		MaxNetworkRetries: &maxRetries,
	}

	conn := &client.API{}
	conn.Init(apiKey, &stripe.Backends{
		API:     stripe.GetBackendWithConfig(stripe.APIBackend, config),
		Uploads: stripe.GetBackendWithConfig(stripe.UploadsBackend, config),
	})

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}
