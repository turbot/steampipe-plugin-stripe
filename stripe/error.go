package stripe

import (
	"context"

	"github.com/stripe/stripe-go/v76"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func isNotFoundError(_ context.Context, _ *plugin.QueryData, _ *plugin.HydrateData, err error) bool {
	if stripeErr, ok := err.(*stripe.Error); ok {
		switch stripeErr.Code {
		case stripe.ErrorCodeMissing, stripe.ErrorCodeResourceMissing:
			return true
		}
	}
	return false
}
