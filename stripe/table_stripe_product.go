package stripe

import (
	"context"

	"github.com/stripe/stripe-go"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableStripeProduct(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "stripe_product",
		Description: "Products available for purchase or subscription.",
		List: &plugin.ListConfig{
			Hydrate:    listProduct,
			KeyColumns: plugin.OptionalColumns([]string{"active", "created", "shippable", "url"}),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getProduct,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier for the product."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The product’s full name or business name."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "The product type."},
			{Name: "unit_label", Type: proto.ColumnType_STRING, Description: "A label that represents units of this product in Stripe and on customers’ receipts and invoices. When set, this will be included in associated invoice line item descriptions."},
			// Other columns
			{Name: "active", Type: proto.ColumnType_BOOL, Description: "Whether the product is currently available for purchase."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Created").Transform(transform.UnixToTimestamp), Description: "Time at which the product was created."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "An arbitrary string attached to the product. Often useful for displaying to users."},
			{Name: "images", Type: proto.ColumnType_JSON, Description: "A list of up to 8 URLs of images for this product, meant to be displayable to the customer."},
			{Name: "livemode", Type: proto.ColumnType_BOOL, Description: "Has the value true if the product exists in live mode or the value false if the product exists in test mode."},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Set of key-value pairs that you can attach to an product. This can be useful for storing additional information about the product in a structured format."},
			{Name: "package_dimensions", Type: proto.ColumnType_JSON, Description: "The dimensions of this product for shipping purposes."},
			{Name: "shippable", Type: proto.ColumnType_BOOL, Description: "Whether this product is shipped (i.e., physical goods)."},
			{Name: "statement_descriptor", Type: proto.ColumnType_STRING, Description: "Extra information about a product which will appear on your customer’s credit card statement. In the case that multiple products are billed at once, the first statement descriptor will be used."},
			{Name: "updated", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Updated").Transform(transform.UnixToTimestamp), Description: "Time at which the product was updated."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "A URL of a publicly-accessible webpage for this product."},
		},
	}
}

func listProduct(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("stripe_product.listProduct", "connection_error", err)
		return nil, err
	}

	params := &stripe.ProductListParams{
		ListParams: stripe.ListParams{
			Context: ctx,
			Limit:   stripe.Int64(100),
		},
	}

	// Exact values can leverage optional key quals for optimal caching
	q := d.KeyColumnQuals
	if q["active"] != nil {
		params.Active = stripe.Bool(q["active"].GetBoolValue())
	}
	if q["shippable"] != nil {
		params.Shippable = stripe.Bool(q["shippable"].GetBoolValue())
	}
	if q["url"] != nil {
		// Note: I can't work out how to set and test a URL for a product?
		params.URL = stripe.String(q["url"].GetStringValue())
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

	plugin.Logger(ctx).Warn("stripe_customer.listInvoice", "q", q)
	plugin.Logger(ctx).Warn("stripe_customer.listInvoice", "quals", quals)

	i := conn.Products.List(params)
	for i.Next() {
		d.StreamListItem(ctx, i.Product())
	}
	if err := i.Err(); err != nil {
		plugin.Logger(ctx).Error("stripe_product.listProduct", "query_error", err, "params", params, "i", i)
		return nil, err
	}

	return nil, nil
}

func getProduct(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("stripe_product.getProduct", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	id := quals["id"].GetStringValue()
	item, err := conn.Products.Get(id, &stripe.ProductParams{})
	if err != nil {
		plugin.Logger(ctx).Error("stripe_product.getProduct", "query_error", err, "id", id)
		return nil, err
	}
	return item, nil
}
