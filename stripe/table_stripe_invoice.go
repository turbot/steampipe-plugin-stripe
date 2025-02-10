package stripe

import (
	"context"

	"github.com/stripe/stripe-go/v76"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableStripeInvoice(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "stripe_invoice",
		Description: "Invoices available for purchase or subscription.",
		List: &plugin.ListConfig{
			Hydrate: listInvoice,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "collection_method", Require: plugin.Optional},
				{Name: "created", Operators: []string{">", ">=", "=", "<", "<="}, Require: plugin.Optional},
				{Name: "due_date", Operators: []string{">", ">=", "=", "<", "<="}, Require: plugin.Optional},
				{Name: "subscription_id", Require: plugin.Optional},
				{Name: "status", Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getInvoice,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier for the invoice."},
			{Name: "number", Type: proto.ColumnType_STRING, Description: "A unique, identifying string that appears on emails sent to the customer for this invoice. This starts with the customer’s unique invoice_prefix if it is specified."},
			{Name: "amount_due", Type: proto.ColumnType_INT, Transform: transform.FromField("AmountDue"), Description: "Final amount due at this time for this invoice. If the invoice’s total is smaller than the minimum charge amount, for example, or if there is account credit that can be applied to the invoice, the amount_due may be 0. If there is a positive starting_balance for the invoice (the customer owes money), the amount_due will also take that into account. The charge that gets generated for the invoice will be for the amount specified in amount_due."},
			{Name: "amount_paid", Type: proto.ColumnType_INT, Transform: transform.FromField("AmountPaid"), Description: "The amount, in cents, that was paid."},
			{Name: "amount_remaining", Type: proto.ColumnType_INT, Transform: transform.FromField("AmountRemaining"), Description: "The amount remaining, in cents, that is due."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Created").Transform(transform.UnixToTimestamp), Description: "Time at which the invoice was created."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The status of the invoice, one of draft, open, paid, uncollectible, or void."},
			// Other columns
			{Name: "account_country", Type: proto.ColumnType_STRING, Description: "The country of the business associated with this invoice, most often the business creating the invoice."},
			{Name: "account_name", Type: proto.ColumnType_STRING, Description: "The public name of the business associated with this invoice, most often the business creating the invoice."},
			{Name: "application_fee_amount", Type: proto.ColumnType_INT, Transform: transform.FromField("ApplicationFeeAmount"), Description: "The fee in cents that will be applied to the invoice and transferred to the application owner’s Stripe account when the invoice is paid."},
			{Name: "attempt_count", Type: proto.ColumnType_INT, Transform: transform.FromField("AttemptCount"), Description: "Number of payment attempts made for this invoice, from the perspective of the payment retry schedule. Any payment attempt counts as the first attempt, and subsequently only automatic retries increment the attempt count. In other words, manual payment attempts after the first attempt do not affect the retry schedule."},
			{Name: "attempted", Type: proto.ColumnType_BOOL, Description: "Whether an attempt has been made to pay the invoice. An invoice is not attempted until 1 hour after the invoice.created webhook, for example, so you might not want to display that invoice as unpaid to your users."},
			{Name: "auto_advance", Type: proto.ColumnType_BOOL, Description: "Controls whether Stripe will perform automatic collection of the invoice. When false, the invoice’s state will not automatically advance without an explicit action."},
			{Name: "billing_reason", Type: proto.ColumnType_STRING, Description: "Indicates the reason why the invoice was created. subscription_cycle indicates an invoice created by a subscription advancing into a new period. subscription_create indicates an invoice created due to creating a subscription. subscription_update indicates an invoice created due to updating a subscription. subscription is set for all old invoices to indicate either a change to a subscription or a period advancement. manual is set for all invoices unrelated to a subscription (for example: created via the invoice editor). The upcoming value is reserved for simulated invoices per the upcoming invoice endpoint. subscription_threshold indicates an invoice created due to a billing threshold being reached."},
			{Name: "charge", Type: proto.ColumnType_JSON, Description: "ID of the latest charge generated for this invoice, if any."},
			{Name: "collection_method", Type: proto.ColumnType_STRING, Description: "Either charge_automatically, or send_invoice. When charging automatically, Stripe will attempt to pay this invoice using the default source attached to the customer. When sending an invoice, Stripe will email this invoice to the customer with payment instructions."},
			{Name: "currency", Type: proto.ColumnType_STRING, Description: "Three-letter ISO currency code, in lowercase. Must be a supported currency."},
			{Name: "custom_fields", Type: proto.ColumnType_JSON, Description: "Custom fields displayed on the invoice."},
			{Name: "customer", Type: proto.ColumnType_JSON, Description: "The ID of the customer who will be billed."},
			{Name: "customer_address", Type: proto.ColumnType_JSON, Description: "The customer’s address. Until the invoice is finalized, this field will equal customer.address. Once the invoice is finalized, this field will no longer be updated."},
			{Name: "customer_email", Type: proto.ColumnType_STRING, Description: "The customer’s email. Until the invoice is finalized, this field will equal customer.email. Once the invoice is finalized, this field will no longer be updated."},
			{Name: "customer_name", Type: proto.ColumnType_STRING, Description: "The customer’s name. Until the invoice is finalized, this field will equal customer.name. Once the invoice is finalized, this field will no longer be updated."},
			{Name: "customer_phone", Type: proto.ColumnType_STRING, Description: "The customer’s phone number. Until the invoice is finalized, this field will equal customer.phone. Once the invoice is finalized, this field will no longer be updated."},
			{Name: "customer_shipping", Type: proto.ColumnType_JSON, Description: "The customer’s shipping information. Until the invoice is finalized, this field will equal customer.shipping. Once the invoice is finalized, this field will no longer be updated."},
			{Name: "customer_tax_exempt", Type: proto.ColumnType_STRING, Description: "The customer’s tax exempt status. Until the invoice is finalized, this field will equal customer.tax_exempt. Once the invoice is finalized, this field will no longer be updated."},
			{Name: "customer_tax_ids", Type: proto.ColumnType_JSON, Description: "The customer’s tax IDs. Until the invoice is finalized, this field will contain the same tax IDs as customer.tax_ids. Once the invoice is finalized, this field will no longer be updated."},
			{Name: "default_payment_method", Type: proto.ColumnType_STRING, Description: "ID of the default payment method for the invoice. It must belong to the customer associated with the invoice. If not set, defaults to the subscription’s default payment method, if any, or to the default payment method in the customer’s invoice settings."},
			{Name: "default_source", Type: proto.ColumnType_STRING, Description: "ID of the default payment source for the invoice. It must belong to the customer associated with the invoice and be in a chargeable state. If not set, defaults to the subscription’s default source, if any, or to the customer’s default source."},
			{Name: "default_tax_rates", Type: proto.ColumnType_JSON, Description: "The tax rates applied to this invoice, if any."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "An arbitrary string attached to the object. Often useful for displaying to users. Referenced as ‘memo’ in the Dashboard."},
			{Name: "discount", Type: proto.ColumnType_JSON, Description: "Describes the current discount applied to this invoice, if there is one. Not populated if there are multiple discounts."},
			{Name: "due_date", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("DueDate").Transform(transform.UnixToTimestamp), Description: "The date on which payment for this invoice is due. This value will be null for invoices where collection_method=charge_automatically."},
			{Name: "ending_balance", Type: proto.ColumnType_INT, Transform: transform.FromField("EndingBalance"), Description: "Ending customer balance after the invoice is finalized. Invoices are finalized approximately an hour after successful webhook delivery or when payment collection is attempted for the invoice. If the invoice has not been finalized yet, this will be null."},
			{Name: "footer", Type: proto.ColumnType_STRING, Description: "Footer displayed on the invoice."},
			{Name: "hosted_invoice_url", Type: proto.ColumnType_STRING, Description: "The URL for the hosted invoice page, which allows customers to view and pay an invoice. If the invoice has not been finalized yet, this will be null."},
			{Name: "invoice_pdf", Type: proto.ColumnType_STRING, Description: "The link to download the PDF for the invoice. If the invoice has not been finalized yet, this will be null."},
			{Name: "lines", Type: proto.ColumnType_JSON, Description: "The individual line items that make up the invoice. lines is sorted as follows: invoice items in reverse chronological order, followed by the subscription, if any."},
			{Name: "livemode", Type: proto.ColumnType_BOOL, Description: "Has the value true if the invoice exists in live mode or the value false if the invoice exists in test mode."},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Set of key-value pairs that you can attach to an invoice. This can be useful for storing additional information about the invoice in a structured format."},
			{Name: "next_payment_attempt", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("NextPaymentAttempt").Transform(transform.UnixToTimestamp), Description: "The time at which payment will next be attempted. This value will be null for invoices where collection_method=send_invoice."},
			{Name: "paid", Type: proto.ColumnType_BOOL, Description: "Whether payment was successfully collected for this invoice. An invoice can be paid (most commonly) with a charge or with credit from the customer’s account balance."},
			{Name: "payment_intent", Type: proto.ColumnType_JSON, Description: "The PaymentIntent associated with this invoice. The PaymentIntent is generated when the invoice is finalized, and can then be used to pay the invoice. Note that voiding an invoice will cancel the PaymentIntent."},
			{Name: "period_end", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("PeriodEnd").Transform(transform.UnixToTimestamp), Description: "End of the usage period during which invoice items were added to this invoice."},
			{Name: "period_start", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("PeriodStart").Transform(transform.UnixToTimestamp), Description: "Start of the usage period during which invoice items were added to this invoice."},
			{Name: "post_payment_credit_notes_amount", Type: proto.ColumnType_INT, Transform: transform.FromField("PostPaymentCreditNotesAmount"), Description: "Total amount of all post-payment credit notes issued for this invoice."},
			{Name: "pre_payment_credit_notes_amount", Type: proto.ColumnType_INT, Transform: transform.FromField("PrePaymentCreditNotesAmount"), Description: "Total amount of all pre-payment credit notes issued for this invoice."},
			{Name: "receipt_number", Type: proto.ColumnType_STRING, Description: "This is the transaction number that appears on email receipts sent for this invoice."},
			{Name: "starting_balance", Type: proto.ColumnType_INT, Transform: transform.FromField("StartingBalance"), Description: "Starting customer balance before the invoice is finalized. If the invoice has not been finalized yet, this will be the current customer balance."},
			{Name: "statement_descriptor", Type: proto.ColumnType_STRING, Description: "Extra information about an invoice for the customer’s credit card statement."},
			{Name: "status_transitions", Type: proto.ColumnType_JSON, Description: "The timestamps at which the invoice status was updated."},
			//{Name: "subscription", Type: proto.ColumnType_JSON, Description: "The subscription that this invoice was prepared for, if any."},
			{Name: "subscription_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Subscription.ID"), Description: "ID of the subscription that this invoice was prepared for, if any."},
			{Name: "subscription_proration_date", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("SubscriptionProrationDate").Transform(transform.UnixToTimestamp), Description: "Only set for upcoming invoices that preview prorations. The time used to calculate prorations."},
			{Name: "subtotal", Type: proto.ColumnType_INT, Transform: transform.FromField("Subtotal"), Description: "Total of all subscriptions, invoice items, and prorations on the invoice before any invoice level discount or tax is applied. Item discounts are already incorporated"},
			{Name: "tax", Type: proto.ColumnType_INT, Transform: transform.FromField("Tax"), Description: "The amount of tax on this invoice. This is the sum of all the tax amounts on this invoice."},
			{Name: "threshold_reason", Type: proto.ColumnType_JSON, Description: "If billing_reason is set to subscription_threshold this returns more information on which threshold rules triggered the invoice."},
			{Name: "total", Type: proto.ColumnType_INT, Transform: transform.FromField("Total"), Description: "Total after discounts and taxes."},
			{Name: "total_tax_amounts", Type: proto.ColumnType_JSON, Description: "The aggregate amounts calculated per tax rate for all line items."},
			{Name: "transfer_data", Type: proto.ColumnType_JSON, Description: "The account (if any) the payment will be attributed to for tax reporting, and where funds from the payment will be transferred to for the invoice."},
			{Name: "webhooks_delivered_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("WebhooksDeliveredAt").Transform(transform.UnixToTimestamp), Description: "Invoices are automatically paid or sent 1 hour after webhooks are delivered, or until all webhook delivery attempts have been exhausted. This field tracks the time when webhooks for this invoice were successfully delivered. If the invoice had no webhooks to deliver, this will be set while the invoice is being created."},
		}),
	}
}

func listInvoice(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("stripe_invoice.listInvoice", "connection_error", err)
		return nil, err
	}

	params := &stripe.InvoiceListParams{
		ListParams: stripe.ListParams{
			Context: ctx,
			Limit:   stripe.Int64(100),
			Expand:  stripe.StringSlice([]string{"data.default_payment_method", "data.default_source", "data.subscription"}),
		},
	}

	equalQuals := d.EqualsQuals
	if equalQuals["status"] != nil {
		params.Status = stripe.String(equalQuals["status"].GetStringValue())
	}
	if equalQuals["collection_method"] != nil {
		params.CollectionMethod = stripe.String(equalQuals["collection_method"].GetStringValue())
	}
	if equalQuals["subscription_id"] != nil {
		params.Subscription = stripe.String(equalQuals["subscription_id"].GetStringValue())
	}

	// Comparison values
	quals := d.Quals

	if quals["created"] != nil {
		for _, q := range quals["created"].Quals {
			tsSecs := q.Value.GetTimestampValue().GetSeconds()
			switch q.Operator {
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
				params.Created = stripe.Int64(tsSecs)
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

	if quals["due_date"] != nil {
		for _, q := range quals["due_date"].Quals {
			tsSecs := q.Value.GetTimestampValue().GetSeconds()
			switch q.Operator {
			case ">":
				if params.DueDateRange == nil {
					params.DueDateRange = &stripe.RangeQueryParams{}
				}
				params.DueDateRange.GreaterThan = tsSecs
			case ">=":
				if params.DueDateRange == nil {
					params.DueDateRange = &stripe.RangeQueryParams{}
				}
				params.DueDateRange.GreaterThanOrEqual = tsSecs
			case "=":
				params.DueDate = stripe.Int64(tsSecs)
			case "<=":
				if params.DueDateRange == nil {
					params.DueDateRange = &stripe.RangeQueryParams{}
				}
				params.DueDateRange.LesserThanOrEqual = tsSecs
			case "<":
				if params.DueDateRange == nil {
					params.DueDateRange = &stripe.RangeQueryParams{}
				}
				params.DueDateRange.LesserThan = tsSecs
			}
		}
	}

	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if *limit < *params.ListParams.Limit {
			params.ListParams.Limit = limit
		}
	}

	var count int64
	i := conn.Invoices.List(params)
	for i.Next() {
		d.StreamListItem(ctx, i.Invoice())
		count++
		if limit != nil {
			if count >= *limit {
				break
			}
		}
	}
	if err := i.Err(); err != nil {
		plugin.Logger(ctx).Error("stripe_invoice.listInvoice", "query_error", err, "params", params, "i", i)
		return nil, err
	}

	return nil, nil
}

func getInvoice(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("stripe_invoice.getInvoice", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	id := quals["id"].GetStringValue()
	item, err := conn.Invoices.Get(id, &stripe.InvoiceParams{})
	if err != nil {
		plugin.Logger(ctx).Error("stripe_invoice.getInvoice", "query_error", err, "id", id)
		return nil, err
	}
	return item, nil
}
