package stripe

import (
	"context"

	"github.com/stripe/stripe-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableStripePlan(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "stripe_plan",
		Description: "Plans define the base price, currency, and billing cycle for recurring purchases of products.",
		List: &plugin.ListConfig{
			Hydrate: listPlan,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "active", Operators: []string{"=", "<>"}, Require: plugin.Optional},
				{Name: "created", Operators: []string{">", ">=", "=", "<", "<="}, Require: plugin.Optional},
				{Name: "product_id", Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getPlan,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier for the plan."},
			{Name: "nickname", Type: proto.ColumnType_STRING, Description: "A brief description of the plan, hidden from customers."},
			// Other columns
			{Name: "active", Type: proto.ColumnType_BOOL, Description: "Whether the plan is currently available for purchase."},
			{Name: "aggregate_usage", Type: proto.ColumnType_STRING, Description: "Specifies a usage aggregation strategy for plans of usage_type=metered. Allowed values are sum for summing up all usage during a period, last_during_period for using the last usage record reported within a period, last_ever for using the last usage record ever (across period bounds) or max which uses the usage record with the maximum reported usage during a period. Defaults to sum."},
			{Name: "amount", Type: proto.ColumnType_INT, Transform: transform.FromField("Amount"), Description: "The unit amount in cents to be charged, represented as a whole integer if possible. Only set if billing_scheme=per_unit."},
			{Name: "amount_decimal", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("AmountDecimal"), Description: "The unit amount in cents to be charged, represented as a decimal string with at most 12 decimal places. Only set if billing_scheme=per_unit."},
			{Name: "billing_scheme", Type: proto.ColumnType_STRING, Description: "Describes how to compute the price per period. Either per_unit or tiered."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Created").Transform(transform.UnixToTimestamp), Description: "Time at which the plan was created."},
			{Name: "currency", Type: proto.ColumnType_STRING, Description: "Three-letter ISO currency code, in lowercase. Must be a supported currency."},
			{Name: "deleted", Type: proto.ColumnType_BOOL, Description: "True if the plan is marked as deleted."},
			{Name: "interval", Type: proto.ColumnType_STRING, Description: "The frequency at which a subscription is billed. One of day, week, month or year."},
			{Name: "interval_count", Type: proto.ColumnType_INT, Transform: transform.FromField("IntervalCount"), Description: "The number of intervals (specified in the interval attribute) between subscription billings. For example, interval=month and interval_count=3 bills every 3 months."},
			{Name: "livemode", Type: proto.ColumnType_BOOL, Description: "Has the value true if the plan exists in live mode or the value false if the plan exists in test mode."},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Set of key-value pairs that you can attach to an plan. This can be useful for storing additional information about the plan in a structured format."},
			{Name: "product_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Product.ID"), Description: "ID of the product whose pricing this plan determines."},
			{Name: "tiers", Type: proto.ColumnType_JSON, Description: "Each element represents a pricing tier. This parameter requires billing_scheme to be set to tiered."},
			{Name: "tiers_mode", Type: proto.ColumnType_STRING, Description: "Defines if the tiering price should be graduated or volume based. In volume-based tiering, the maximum quantity within a period determines the per unit price. In graduated tiering, pricing can change as the quantity grows."},
			{Name: "transform_usage", Type: proto.ColumnType_JSON, Description: "Apply a transformation to the reported usage or set quantity before computing the amount billed."},
			{Name: "trial_period_days", Type: proto.ColumnType_INT, Transform: transform.FromField("TrialPeriodDays"), Description: "Default number of trial days when subscribing a customer to this plan using trial_from_plan=true."},
			{Name: "usage_type", Type: proto.ColumnType_STRING, Description: "Configures how the quantity per period should be determined. Can be either metered or licensed. licensed automatically bills the quantity set when adding it to a subscription. metered aggregates the total usage based on usage records. Defaults to licensed."},
		}),
	}
}

func listPlan(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("stripe_plan.listPlan", "connection_error", err)
		return nil, err
	}
	params := &stripe.PlanListParams{
		ListParams: stripe.ListParams{
			Context: ctx,
			Limit:   stripe.Int64(100),
		},
	}

	equalQuals := d.EqualsQuals
	if equalQuals["product_id"] != nil {
		params.Product = stripe.String(equalQuals["product_id"].GetStringValue())
	}

	// Comparison values
	quals := d.Quals

	if quals["active"] != nil {
		for _, q := range quals["active"].Quals {
			switch q.Operator {
			case "=":
				params.Active = stripe.Bool(q.Value.GetBoolValue())
			case "<>":
				params.Active = stripe.Bool(!q.Value.GetBoolValue())
			}
		}
	}

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
	i := conn.Plans.List(params)
	for i.Next() {
		d.StreamListItem(ctx, i.Plan())
		count++
		if limit != nil {
			if count >= *limit {
				break
			}
		}
	}
	if err := i.Err(); err != nil {
		plugin.Logger(ctx).Error("stripe_plan.listPlan", "query_error", err, "params", params, "i", i)
		return nil, err
	}

	return nil, nil
}

func getPlan(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("stripe_plan.getPlan", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	id := quals["id"].GetStringValue()
	item, err := conn.Plans.Get(id, &stripe.PlanParams{})
	if err != nil {
		plugin.Logger(ctx).Error("stripe_plan.getPlan", "query_error", err, "id", id)
		return nil, err
	}
	return item, nil
}
