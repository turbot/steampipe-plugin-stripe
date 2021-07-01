package stripe

import (
	"context"

	"github.com/stripe/stripe-go"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableStripeCustomer(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "stripe_customer",
		Description: "Customer details.",
		List: &plugin.ListConfig{
			Hydrate:    listCustomer,
			KeyColumns: plugin.OptionalColumns([]string{"created", "email"}),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getCustomer,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier for the customer."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "The customer’s email address."},
			// Other columns
			{Name: "address", Type: proto.ColumnType_JSON, Description: "The customer’s address."},
			{Name: "balance", Type: proto.ColumnType_INT, Transform: transform.FromField("Balance"), Description: "Current balance, if any, being stored on the customer. If negative, the customer has credit to apply to their next invoice. If positive, the customer has an amount owed that will be added to their next invoice. The balance does not refer to any unpaid invoices; it solely takes into account amounts that have yet to be successfully applied to any invoice. This balance is only taken into account as invoices are finalized."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Created").Transform(transform.UnixToTimestamp), Description: "Time at which the object was created."},
			{Name: "currency", Type: proto.ColumnType_STRING, Description: "Three-letter ISO code for the currency the customer can be charged in for recurring billing purposes."},
			{Name: "default_source", Type: proto.ColumnType_JSON, Description: "ID of the default payment source for the customer."},
			{Name: "deleted", Type: proto.ColumnType_BOOL, Description: "True if the customer is marked as deleted."},
			{Name: "delinquent", Type: proto.ColumnType_BOOL, Description: "When the customer’s latest invoice is billed by charging automatically, delinquent is true if the invoice’s latest charge failed. When the customer’s latest invoice is billed by sending an invoice, delinquent is true if the invoice isn’t paid by its due date. If an invoice is marked uncollectible by dunning, delinquent doesn’t get reset to false."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "An arbitrary string attached to the object. Often useful for displaying to users."},
			{Name: "discount", Type: proto.ColumnType_JSON, Description: "Describes the current discount active on the customer, if there is one."},
			{Name: "invoice_prefix", Type: proto.ColumnType_STRING, Description: "The prefix for the customer used to generate unique invoice numbers."},
			{Name: "invoice_settings", Type: proto.ColumnType_JSON, Description: "The customer’s default invoice settings."},
			{Name: "livemode", Type: proto.ColumnType_BOOL, Description: "Has the value true if the object exists in live mode or the value false if the object exists in test mode."},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The customer’s full name or business name."},
			{Name: "next_invoice_sequence", Type: proto.ColumnType_INT, Description: "The suffix of the customer’s next invoice number, e.g., 0001."},
			{Name: "phone", Type: proto.ColumnType_STRING, Description: "The customer’s phone number."},
			{Name: "preferred_locales", Type: proto.ColumnType_JSON, Description: "The customer’s preferred locales (languages), ordered by preference."},
			{Name: "shipping", Type: proto.ColumnType_JSON, Description: "Mailing and shipping address for the customer. Appears on invoices emailed to this customer."},
			//{Name: "sources", Type: proto.ColumnType_JSON, Transform: transform.FromField("Sources.Data"), Description: "The customer’s payment sources, if any."},
			//{Name: "subscriptions", Type: proto.ColumnType_JSON, Transform: transform.FromField("Subscriptions.Data"), Description: "The customer’s current subscriptions, if any."},
			{Name: "tax_exempt", Type: proto.ColumnType_STRING, Description: "Describes the customer’s tax exemption status. One of none, exempt, or reverse."},
			{Name: "tax_ids", Type: proto.ColumnType_JSON, Description: "The customer’s tax IDs."},
		},
	}
}

func listCustomer(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("stripe_customer.listCustomer", "connection_error", err)
		return nil, err
	}
	params := &stripe.CustomerListParams{
		ListParams: stripe.ListParams{
			Context: ctx,
			Limit:   stripe.Int64(100),
		},
	}

	// Exact values can leverage optional key quals for optimal caching
	q := d.KeyColumnQuals
	if q["email"] != nil {
		params.Email = stripe.String(q["email"].GetStringValue())
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

	i := conn.Customers.List(params)
	for i.Next() {
		d.StreamListItem(ctx, i.Customer())
	}
	if err := i.Err(); err != nil {
		plugin.Logger(ctx).Error("stripe_customer.listCustomer", "query_error", err, "params", params, "i", i)
		return nil, err
	}
	return nil, nil
}

func getCustomer(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("stripe_customer.getCustomer", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	id := quals["id"].GetStringValue()
	item, err := conn.Customers.Get(id, &stripe.CustomerParams{})
	if err != nil {
		plugin.Logger(ctx).Error("stripe_customer.getCustomer", "query_error", err, "id", id)
		return nil, err
	}
	return item, nil
}
