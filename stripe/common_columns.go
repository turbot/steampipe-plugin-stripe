package stripe

import (
	"context"

	"github.com/stripe/stripe-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "account_id",
			Description: "The Stripe account ID.",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getAccountId,
			Transform:   transform.FromValue(),
		},
	}, c...)
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getAccountMemoized = plugin.HydrateFunc(getAccountUncached).Memoize(memoize.WithCacheKeyFunction(getAccountCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getAccount(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getAccountMemoized(ctx, d, h)
}

func getAccountId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	acc, err := getAccountMemoized(ctx, d, h)
	if err != nil {
		return nil, err
	}
	return acc.(*stripe.Account).ID, nil
}

// Build a cache key for the call to getAccountIdCacheKey.
func getAccountCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getAccountId"
	return key, nil
}

func getAccountUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("stripe_account.listAccount", "connection_error", err)
		return nil, err
	}

	item, err := conn.Account.Get()
	if err != nil {
		plugin.Logger(ctx).Error("stripe_customer.listAccount", "query_error", err)
		return nil, err
	}

	return item, nil
}
