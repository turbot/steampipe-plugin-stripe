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
			OptionalKeyColumns: plugin.AnyColumn([]string{"customer_id", "collection_method", "status"}),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSubscription,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier for the subscription."},
			{Name: "customer_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Customer.ID"), Description: "ID of the customer who owns the subscription."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Possible values are incomplete, incomplete_expired, trialing, active, past_due, canceled, or unpaid."},
			// Other columns
			{Name: "application_fee_percent", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("ApplicationFeePercent"), Description: "A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice subtotal that will be transferred to the application owner’s Stripe account."},
			{Name: "billing_cycle_anchor", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("BillingCycleAnchor").Transform(transform.UnixToTimestamp), Description: "Determines the date of the first full invoice, and, for plans with month or year intervals, the day of the month for subsequent invoices."},
			{Name: "billing_thresholds", Type: proto.ColumnType_JSON, Description: "Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period."},
			{Name: "cancel_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CancelAt").Transform(transform.UnixToTimestamp), Description: "A date in the future at which the subscription will automatically get canceled."},
			{Name: "cancel_at_period_end", Type: proto.ColumnType_BOOL, Description: "If the subscription has been canceled with the at_period_end flag set to true, cancel_at_period_end on the subscription will be true. You can use this attribute to determine whether a subscription that has a status of active is scheduled to be canceled at the end of the current period."},
			{Name: "canceled_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CanceledAt").Transform(transform.UnixToTimestamp), Description: "If the subscription has been canceled, the date of that cancellation. If the subscription was canceled with cancel_at_period_end, canceled_at will reflect the time of the most recent update request, not the end of the subscription period when the subscription is automatically moved to a canceled state."},
			{Name: "collection_method", Type: proto.ColumnType_STRING, Description: "Either charge_automatically, or send_invoice. When charging automatically, Stripe will attempt to pay this subscription at the end of the cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Created").Transform(transform.UnixToTimestamp), Description: "Time at which the subscription was created."},
			{Name: "current_period_end", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CurrentPeriodEnd").Transform(transform.UnixToTimestamp), Description: "End of the current period that the subscription has been invoiced for. At the end of this period, a new invoice will be created."},
			{Name: "current_period_start", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CurrentPeriodStart").Transform(transform.UnixToTimestamp), Description: "Start of the current period that the subscription has been invoiced for."},
			//{Name: "customer", Type: proto.ColumnType_JSON, Description: "Customer who owns the subscription."},
			{Name: "days_until_due", Type: proto.ColumnType_INT, Transform: transform.FromField("DaysUntilDue"), Description: "Number of days a customer has to pay invoices generated by this subscription. This value will be null for subscriptions where collection_method=charge_automatically."},
			{Name: "default_payment_method", Type: proto.ColumnType_JSON, Description: "ID of the default payment method for the subscription. It must belong to the customer associated with the subscription. This takes precedence over default_source. If neither are set, invoices will use the customer’s invoice_settings.default_payment_method or default_source."},
			{Name: "default_source", Type: proto.ColumnType_JSON, Description: "ID of the default payment source for the subscription. It must belong to the customer associated with the subscription and be in a chargeable state. If default_payment_method is also set, default_payment_method will take precedence. If neither are set, invoices will use the customer’s invoice_settings.default_payment_method or default_source."},
			{Name: "default_tax_rates", Type: proto.ColumnType_JSON, Description: "The tax rates that will apply to any subscription item that does not have tax_rates set. Invoices created will have their default_tax_rates populated from the subscription."},
			{Name: "discount", Type: proto.ColumnType_JSON, Description: "Describes the current discount applied to this subscription, if there is one. When billing, a discount applied to a subscription overrides a discount applied on a customer-wide basis."},
			{Name: "ended_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("EndedAt").Transform(transform.UnixToTimestamp), Description: "If the subscription has ended, the date the subscription ended."},
			{Name: "items", Type: proto.ColumnType_JSON, Transform: transform.FromField("Items.Data"), Description: "List of subscription items, each with an attached price."},
			//{Name: "latest_invoice", Type: proto.ColumnType_JSON, Description: "The most recent invoice this subscription has generated."},
			{Name: "latest_invoice_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("LatestInvoice.ID"), Description: "ID of the most recent invoice this subscription has generated."},
			{Name: "livemode", Type: proto.ColumnType_BOOL, Description: "Has the value true if the subscription exists in live mode or the value false if the subscription exists in test mode."},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Set of key-value pairs that you can attach to an subscription. This can be useful for storing additional information about the subscription in a structured format."},
			{Name: "next_pending_invoice_item_invoice", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("EndedAt").Transform(transform.UnixToTimestamp), Description: "Specifies the approximate timestamp on which any pending invoice items will be billed according to the schedule provided at pending_invoice_item_interval."},
			{Name: "pause_collection", Type: proto.ColumnType_JSON, Description: "If specified, payment collection for this subscription will be paused."},
			{Name: "pending_invoice_item_interval", Type: proto.ColumnType_JSON, Description: "Specifies an interval for how often to bill for any pending invoice items. It is analogous to calling Create an invoice for the given subscription at the specified interval."},
			{Name: "pending_setup_intent", Type: proto.ColumnType_STRING, Description: "You can use this SetupIntent to collect user authentication when creating a subscription without immediate payment or updating a subscription’s payment method, allowing you to optimize for off-session payments. Learn more in the SCA Migration Guide."},
			{Name: "pending_update", Type: proto.ColumnType_JSON, Description: "If specified, pending updates that will be applied to the subscription once the latest_invoice has been paid."},
			//{Name: "plan", Type: proto.ColumnType_JSON, Description: ""},
			//{Name: "quantity", Type: proto.ColumnType_INT, Transform: transform.FromField("Quantity"), Description: ""},
			{Name: "schedule", Type: proto.ColumnType_JSON, Description: "The schedule attached to the subscription."},
			{Name: "start_date", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("StartDate").Transform(transform.UnixToTimestamp), Description: "Date when the subscription was first created. The date might differ from the created date due to backdating."},
			{Name: "transfer_data", Type: proto.ColumnType_JSON, Description: "The account (if any) the subscription’s payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription’s invoices."},
			{Name: "trial_end", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("TrialEnd").Transform(transform.UnixToTimestamp), Description: "If the subscription has a trial, the end of that trial."},
			{Name: "trial_start", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("TrialStart").Transform(transform.UnixToTimestamp), Description: "If the subscription has a trial, the start of that trial."},
		},
	}
}

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

	// Exact values can leverage optional key quals for optimal caching
	q := d.OptionalKeyColumnQuals
	if q["status"] != nil {
		params.Status = q["status"].GetStringValue()
	}
	if q["collection_method"] != nil {
		params.CollectionMethod = stripe.String(q["collection_method"].GetStringValue())
	}
	if q["customer_id"] != nil {
		params.Customer = q["customer_id"].GetStringValue()
	}

	// Comparison values
	quals := d.QueryContext.RawQuals

	if quals["created"] != nil {
		for _, q := range quals["created"].Quals {
			op := q.GetStringValue()
			tsSecs := q.Value.GetTimestampValue().GetSeconds()
			switch op {
			case ">":
				if params.CreatedRange == nil {
					params.CreatedRange = &stripe.RangeQueryParams{}
				}
				params.CreatedRange.GreaterThan = tsSecs
			case ">=":
				if params.CreatedRange == nil {
					params.CreatedRange = &stripe.RangeQueryParams{}
				}
				params.CreatedRange.GreaterThanOrEqual = tsSecs
			case "=":
				params.Created = tsSecs
			case "<=":
				if params.CreatedRange == nil {
					params.CreatedRange = &stripe.RangeQueryParams{}
				}
				params.CreatedRange.LesserThanOrEqual = tsSecs
			case "<":
				if params.CreatedRange == nil {
					params.CreatedRange = &stripe.RangeQueryParams{}
				}
				params.CreatedRange.LesserThan = tsSecs
			}
		}
	}

	if quals["current_period_start"] != nil {
		for _, q := range quals["current_period_start"].Quals {
			op := q.GetStringValue()
			tsSecs := q.Value.GetTimestampValue().GetSeconds()
			switch op {
			case ">":
				if params.CurrentPeriodStartRange == nil {
					params.CurrentPeriodStartRange = &stripe.RangeQueryParams{}
				}
				params.CurrentPeriodStartRange.GreaterThan = tsSecs
			case ">=":
				if params.CurrentPeriodStartRange == nil {
					params.CurrentPeriodStartRange = &stripe.RangeQueryParams{}
				}
				params.CurrentPeriodStartRange.GreaterThanOrEqual = tsSecs
			case "=":
				params.CurrentPeriodStart = stripe.Int64(tsSecs)
			case "<=":
				if params.CurrentPeriodStartRange == nil {
					params.CurrentPeriodStartRange = &stripe.RangeQueryParams{}
				}
				params.CurrentPeriodStartRange.LesserThanOrEqual = tsSecs
			case "<":
				if params.CurrentPeriodStartRange == nil {
					params.CurrentPeriodStartRange = &stripe.RangeQueryParams{}
				}
				params.CurrentPeriodStartRange.LesserThan = tsSecs
			}
		}
	}

	if quals["current_period_end"] != nil {
		for _, q := range quals["current_period_end"].Quals {
			op := q.GetStringValue()
			tsSecs := q.Value.GetTimestampValue().GetSeconds()
			switch op {
			case ">":
				if params.CurrentPeriodEndRange == nil {
					params.CurrentPeriodEndRange = &stripe.RangeQueryParams{}
				}
				params.CurrentPeriodEndRange.GreaterThan = tsSecs
			case ">=":
				if params.CurrentPeriodEndRange == nil {
					params.CurrentPeriodEndRange = &stripe.RangeQueryParams{}
				}
				params.CurrentPeriodEndRange.GreaterThanOrEqual = tsSecs
			case "=":
				params.CurrentPeriodEnd = stripe.Int64(tsSecs)
			case "<=":
				if params.CurrentPeriodEndRange == nil {
					params.CurrentPeriodEndRange = &stripe.RangeQueryParams{}
				}
				params.CurrentPeriodEndRange.LesserThanOrEqual = tsSecs
			case "<":
				if params.CurrentPeriodEndRange == nil {
					params.CurrentPeriodEndRange = &stripe.RangeQueryParams{}
				}
				params.CurrentPeriodEndRange.LesserThan = tsSecs
			}
		}
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