package stripe

import (
	"context"

	"github.com/stripe/stripe-go/v76"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableStripeCharge(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "stripe_charge",
		Description: "Retrieves comprehensive details of historical Stripe charges or a specific charge by ID.",
		List: &plugin.ListConfig{
			Hydrate: listCharges,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "created", Operators: []string{">", ">=", "=", "<", "<="}, Require: plugin.Optional},
				{Name: "payment_intent", Require: plugin.Optional, Operators: []string{"="}},
				{Name: "transfer_group", Require: plugin.Optional, Operators: []string{"="}},
				{Name: "customer", Require: plugin.Optional, Operators: []string{"="}},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getCharge,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: commonColumns([]*plugin.Column{
			// Basic fields
			{
				Name:        "id",
				Description: "Unique identifier for the charge.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "amount",
				Description: "Amount charged in cents.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "amount_refunded",
				Description: "Amount refunded in cents.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "authorization_code",
				Description: "Authorization code for the charge.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "balance_transaction",
				Description: "Balance transaction related to the charge.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "captured",
				Description: "Indicates whether the charge was captured.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "created",
				Description: "Timestamp when the charge was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Created").Transform(transform.UnixToTimestamp),
			},
			{
				Name:        "currency",
				Description: "Currency of the charge.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "customer",
				Description: "Customer related to the charge.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "Description of the charge.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "disputed",
				Description: "Indicates whether the charge is disputed.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "failure_code",
				Description: "Failure code if the charge failed.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "failure_message",
				Description: "Failure message if the charge failed.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "livemode",
				Description: "Indicates whether the charge was created in live mode.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "paid",
				Description: "Indicates whether the charge was paid.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "payment_intent",
				Description: "Payment intent related to the charge.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "payment_method",
				Description: "Payment method used for the charge.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "receipt_email",
				Description: "Receipt email for the charge.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "receipt_number",
				Description: "Receipt number for the charge.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "receipt_url",
				Description: "URL to the receipt of the charge.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "refunded",
				Description: "Indicates whether the charge was refunded.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "status",
				Description: "Status of the charge (e.g., succeeded, pending, failed).",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "transfer_group",
				Description: "Transfer group related to the charge.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON columns for complex data
			{
				Name:        "application",
				Description: "Application that initiated the charge.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "application_fee",
				Description: "Application fee related to the charge.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "billing_details",
				Description: "Billing details associated with the charge.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "destination",
				Description: "Destination account receiving the funds.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "dispute",
				Description: "Details about the dispute if the charge is disputed.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "fraud_details",
				Description: "Fraud details for the charge.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "invoice",
				Description: "Invoice associated with the charge.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "level3",
				Description: "Level 3 data for the charge.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "metadata",
				Description: "Metadata associated with the charge.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "outcome",
				Description: "Details about the outcome of the charge.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "payment_method_details",
				Description: "Details about the payment method used for the charge.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "refunds",
				Description: "List of refunds applied to the charge.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "review",
				Description: "Review associated with the charge.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "shipping",
				Description: "Shipping details for the charge.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "source",
				Description: "Payment source for the charge.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "source_transfer",
				Description: "Transfer related to the charge source.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "transfer",
				Description: "Transfer related to the charge.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "transfer_data",
				Description: "Transfer data for the charge.",
				Type:        proto.ColumnType_JSON,
			},
		}),
	}
}

func listCharges(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("stripe_charge.listCharges", "connection_error", err)
		return nil, err
	}
	params := &stripe.ChargeListParams{
		ListParams: stripe.ListParams{
			Context: ctx,
			Limit:   stripe.Int64(100),
		},
	}

	q := d.EqualsQuals
	if q["customer"] != nil {
		params.Customer = stripe.String(q["customer"].GetStringValue())
	}
	if q["payment_intent"] != nil {
		params.PaymentIntent = stripe.String(q["payment_intent"].GetStringValue())
	}
	if q["transfer_group"] != nil {
		params.TransferGroup = stripe.String(q["transfer_group"].GetStringValue())
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

	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if *limit < *params.ListParams.Limit {
			params.ListParams.Limit = limit
		}
	}

	var count int64
	i := conn.Charges.List(params)
	for i.Next() {
		d.StreamListItem(ctx, i.Charge())
		count++
		if limit != nil {
			if count >= *limit {
				break
			}
		}
	}
	if err := i.Err(); err != nil {
		plugin.Logger(ctx).Error("stripe_charge.listCharges", "query_error", err, "params", params, "i", i)
		return nil, err
	}

	return nil, nil
}

func getCharge(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("stripe_charge.getCharge", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	id := quals["id"].GetStringValue()
	item, err := conn.Charges.Get(id, &stripe.ChargeParams{})
	if err != nil {
		plugin.Logger(ctx).Error("stripe_charge.getCharge", "query_error", err, "id", id)
		return nil, err
	}
	return item, nil
}
