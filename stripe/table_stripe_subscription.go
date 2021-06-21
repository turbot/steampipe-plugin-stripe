package stripe

import (
	"context"

	"github.com/stripe/stripe-go"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableStripeSubscription(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "stripe_subscription",
		Description: "Subscriptions available for purchase or subscription.",
		List: &plugin.ListConfig{
			Hydrate:            listSubscription,
			OptionalKeyColumns: plugin.AnyColumn([]string{"customer", "collection_method", "status"}),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getPlan,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier for the subscription."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The subscription’s full name or business name."},
			// Other columns
			{Name: "amount_off", Type: proto.ColumnType_INT, Transform: transform.FromField("AmountOff"), Description: "Amount (in the currency specified) that will be taken off the subtotal of any invoices for this customer."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Created").Transform(transform.UnixToTimestamp), Description: "Time at which the subscription was created."},
			{Name: "currency", Type: proto.ColumnType_STRING, Description: "If amount_off has been set, the three-letter ISO code for the currency of the amount to take off."},
			{Name: "deleted", Type: proto.ColumnType_BOOL, Description: "True if the customer is marked as deleted."},
			{Name: "duration", Type: proto.ColumnType_STRING, Description: "One of forever, once, and repeating. Describes how long a customer who applies this subscription will get the discount."},
			{Name: "duration_in_months", Type: proto.ColumnType_INT, Transform: transform.FromField("DurationInMonths"), Description: "If duration is repeating, the number of months the subscription applies. Null if subscription duration is forever or once."},
			{Name: "livemode", Type: proto.ColumnType_BOOL, Description: "Has the value true if the subscription exists in live mode or the value false if the subscription exists in test mode."},
			{Name: "max_redemptions", Type: proto.ColumnType_INT, Transform: transform.FromField("MaxRedemptions"), Description: "Maximum number of times this subscription can be redeemed, in total, across all customers, before it is no longer valid."},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Set of key-value pairs that you can attach to an subscription. This can be useful for storing additional information about the subscription in a structured format."},
			{Name: "percent_off", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("PercentOff"), Description: "Percent that will be taken off the subtotal of any invoices for this customer for the duration of the subscription. For example, a subscription with percent_off of 50 will make a $100 invoice $50 instead."},
			{Name: "redeem_by", Type: proto.ColumnType_TIMESTAMP, Description: "Date after which the subscription can no longer be redeemed."},
			{Name: "times_redeemed", Type: proto.ColumnType_INT, Transform: transform.FromField("TimesRedeemed"), Description: "Number of times this subscription has been applied to a customer."},
			{Name: "valid", Type: proto.ColumnType_BOOL, Description: "Taking account of the above properties, whether this subscription can still be applied to a customer."},
		},
	}
}

/*
type Subscription struct {
	ApplicationFeePercent         float64                                `json:"application_fee_percent"`
	BillingCycleAnchor            int64                                  `json:"billing_cycle_anchor"`
	BillingThresholds             *SubscriptionBillingThresholds         `json:"billing_thresholds"`
	CancelAt                      int64                                  `json:"cancel_at"`
	CancelAtPeriodEnd             bool                                   `json:"cancel_at_period_end"`
	CanceledAt                    int64                                  `json:"canceled_at"`
	CollectionMethod              SubscriptionCollectionMethod           `json:"collection_method"`
	Created                       int64                                  `json:"created"`
	CurrentPeriodEnd              int64                                  `json:"current_period_end"`
	CurrentPeriodStart            int64                                  `json:"current_period_start"`
	Customer                      *Customer                              `json:"customer"`
	DaysUntilDue                  int64                                  `json:"days_until_due"`
	DefaultPaymentMethod          *PaymentMethod                         `json:"default_payment_method"`
	DefaultSource                 *PaymentSource                         `json:"default_source"`
	DefaultTaxRates               []*TaxRate                             `json:"default_tax_rates"`
	Discount                      *Discount                              `json:"discount"`
	EndedAt                       int64                                  `json:"ended_at"`
	ID                            string                                 `json:"id"`
	Items                         *SubscriptionItemList                  `json:"items"`
	LatestInvoice                 *Invoice                               `json:"latest_invoice"`
	Livemode                      bool                                   `json:"livemode"`
	Metadata                      map[string]string                      `json:"metadata"`
	NextPendingInvoiceItemInvoice int64                                  `json:"next_pending_invoice_item_invoice"`
	Object                        string                                 `json:"object"`
	OnBehalfOf                    *Account                               `json:"on_behalf_of"`
	PauseCollection               SubscriptionPauseCollection            `json:"pause_collection"`
	PendingInvoiceItemInterval    SubscriptionPendingInvoiceItemInterval `json:"pending_invoice_item_interval"`
	PendingSetupIntent            *SetupIntent                           `json:"pending_setup_intent"`
	PendingUpdate                 *SubscriptionPendingUpdate             `json:"pending_update"`
	Plan                          *Plan                                  `json:"plan"`
	Quantity                      int64                                  `json:"quantity"`
	Schedule                      *SubscriptionSchedule                  `json:"schedule"`
	StartDate                     int64                                  `json:"start_date"`
	Status                        SubscriptionStatus                     `json:"status"`
	TransferData                  *SubscriptionTransferData              `json:"transfer_data"`
	TrialEnd                      int64                                  `json:"trial_end"`
	TrialStart                    int64                                  `json:"trial_start"`

	// This field is deprecated and we recommend that you use TaxRates instead.
	TaxPercent float64 `json:"tax_percent"`
}
*/

func listSubscription(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("stripe_subscription.listSubscription", "connection_error", err)
		return nil, err
	}

	params := &stripe.SubscriptionListParams{
		ListParams: stripe.ListParams{
			Context: ctx,
			Limit:   stripe.Int64(100),
		},
	}
	i := conn.Subscriptions.List(params)
	for i.Next() {
		d.StreamListItem(ctx, i.Subscription())
	}
	if err := i.Err(); err != nil {
		plugin.Logger(ctx).Error("stripe_subscription.listSubscription", "query_error", err, "params", params, "i", i)
		return nil, err
	}

	return nil, nil
}

func getSubscription(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("stripe_subscription.getSubscription", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	id := quals["id"].GetStringValue()
	item, err := conn.Subscriptions.Get(id, &stripe.SubscriptionParams{})
	if err != nil {
		plugin.Logger(ctx).Error("stripe_subscription.getSubscription", "query_error", err, "id", id)
		return nil, err
	}
	return item, nil
}
