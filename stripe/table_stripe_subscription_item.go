package stripe

import (
	"context"
	//"time"

	"github.com/stripe/stripe-go/v76"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableStripeSubscriptionItem(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "stripe_subscription_item",
		Description: "Subscription Items in Stripe represent the individual products that a customer is subscribed to.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("subscription_id"),
			Hydrate:    listSubscriptionItem,
		},
		Columns: commonColumns([]*plugin.Column{
			// Add columns relevant to SubscriptionItems here
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier for the subscription."},
			{Name: "plan", Type: proto.ColumnType_JSON, Transform: transform.FromField("Plan"), Description: "A plan represents a billing configuration. (Deprecated)"},
			{Name: "price", Type: proto.ColumnType_JSON, Transform: transform.FromField("Price"), Description: "A price represents a unit cost for a product, specifying the amount, currency, and billing frequency."},
			{Name: "subscription_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Subscription"), Description: "The ID of the subscription this item belongs to."},
			{Name: "usage_record_summaries", Type: proto.ColumnType_JSON, Hydrate: listUsageRecordSummaries, Transform: transform.FromValue()},
		}),
	}
}

// listSubscriptionItem lists all subscription items
func listSubscriptionItem(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("stripe_subscription.listSubscriptionItem", "connection_error", err)
		return nil, err
	}

	subscription_id := d.EqualsQuals["subscription_id"].GetStringValue()
	plugin.Logger(ctx).Debug("stripe_subscription.listSubscriptionItem", "subscription_id", subscription_id)

	params := &stripe.SubscriptionItemListParams{
		Subscription: stripe.String(subscription_id),
	}

	i := conn.SubscriptionItems.List(params)

	for i.Next() {
		d.StreamListItem(ctx, i.SubscriptionItem())
	}

	return nil, nil
}

func listUsageRecordSummaries(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	item := h.Item.(*stripe.SubscriptionItem)

	plugin.Logger(ctx).Debug("stripe_subscription.listUsageRecordSummaries", "item", item)

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("stripe_subscription.listSubscriptionItem", "connection_error", err)
		return nil, err
	}

	params := &stripe.SubscriptionItemUsageRecordSummariesParams{
		SubscriptionItem: stripe.String(item.ID),
	}

	var summaries []*stripe.UsageRecordSummary
	u := conn.SubscriptionItems.UsageRecordSummaries(params)
	for u.Next() {
		summary := u.UsageRecordSummary()
		summaries = append(summaries, summary)
	}

	return summaries, nil
}
