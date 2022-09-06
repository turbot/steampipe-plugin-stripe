package stripe

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-stripe",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError,
		},
		TableMap: map[string]*plugin.Table{
			"stripe_account":      tableStripeAccount(ctx),
			"stripe_coupon":       tableStripeCoupon(ctx),
			"stripe_customer":     tableStripeCustomer(ctx),
			"stripe_invoice":      tableStripeInvoice(ctx),
			"stripe_plan":         tableStripePlan(ctx),
			"stripe_product":      tableStripeProduct(ctx),
			"stripe_subscription": tableStripeSubscription(ctx),
		},
	}
	return p
}
