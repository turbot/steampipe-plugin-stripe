package stripe

import (
	"context"

	"github.com/stripe/stripe-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableStripeCoupon(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "stripe_coupon",
		Description: "Coupons available for purchase or subscription.",
		List: &plugin.ListConfig{
			Hydrate: listCoupon,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "created", Operators: []string{">", ">=", "=", "<", "<="}, Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getCoupon,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier for the coupon."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The coupon’s full name or business name."},
			// Other columns
			{Name: "amount_off", Type: proto.ColumnType_INT, Transform: transform.FromField("AmountOff"), Description: "Amount (in the currency specified) that will be taken off the subtotal of any invoices for this customer."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Created").Transform(transform.UnixToTimestamp), Description: "Time at which the coupon was created."},
			{Name: "currency", Type: proto.ColumnType_STRING, Description: "If amount_off has been set, the three-letter ISO code for the currency of the amount to take off."},
			{Name: "deleted", Type: proto.ColumnType_BOOL, Description: "True if the customer is marked as deleted."},
			{Name: "duration", Type: proto.ColumnType_STRING, Description: "One of forever, once, and repeating. Describes how long a customer who applies this coupon will get the discount."},
			{Name: "duration_in_months", Type: proto.ColumnType_INT, Transform: transform.FromField("DurationInMonths"), Description: "If duration is repeating, the number of months the coupon applies. Null if coupon duration is forever or once."},
			{Name: "livemode", Type: proto.ColumnType_BOOL, Description: "Has the value true if the coupon exists in live mode or the value false if the coupon exists in test mode."},
			{Name: "max_redemptions", Type: proto.ColumnType_INT, Transform: transform.FromField("MaxRedemptions"), Description: "Maximum number of times this coupon can be redeemed, in total, across all customers, before it is no longer valid."},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Set of key-value pairs that you can attach to an coupon. This can be useful for storing additional information about the coupon in a structured format."},
			{Name: "percent_off", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("PercentOff"), Description: "Percent that will be taken off the subtotal of any invoices for this customer for the duration of the coupon. For example, a coupon with percent_off of 50 will make a $100 invoice $50 instead."},
			{Name: "redeem_by", Type: proto.ColumnType_TIMESTAMP, Description: "Date after which the coupon can no longer be redeemed."},
			{Name: "times_redeemed", Type: proto.ColumnType_INT, Transform: transform.FromField("TimesRedeemed"), Description: "Number of times this coupon has been applied to a customer."},
			{Name: "valid", Type: proto.ColumnType_BOOL, Description: "Taking account of the above properties, whether this coupon can still be applied to a customer."},
		}),
	}
}

func listCoupon(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("stripe_coupon.listCoupon", "connection_error", err)
		return nil, err
	}

	params := &stripe.CouponListParams{
		ListParams: stripe.ListParams{
			Context: ctx,
			Limit:   stripe.Int64(100),
		},
	}

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
	i := conn.Coupons.List(params)
	for i.Next() {
		d.StreamListItem(ctx, i.Coupon())
		count++
		if limit != nil {
			if count >= *limit {
				break
			}
		}
	}
	if err := i.Err(); err != nil {
		plugin.Logger(ctx).Error("stripe_coupon.listCoupon", "query_error", err, "params", params, "i", i)
		return nil, err
	}

	return nil, nil
}

func getCoupon(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("stripe_coupon.getCoupon", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	id := quals["id"].GetStringValue()
	item, err := conn.Coupons.Get(id, &stripe.CouponParams{})
	if err != nil {
		plugin.Logger(ctx).Error("stripe_coupon.getCoupon", "query_error", err, "id", id)
		return nil, err
	}
	return item, nil
}
