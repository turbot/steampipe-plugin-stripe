package stripe

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableStripeAccount(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "stripe_account",
		Description: "This is an object representing a Stripe account.",
		List: &plugin.ListConfig{
			Hydrate: listAccount,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier for the account."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "An email address associated with the account. You can treat this as metadata: it is not used for authentication or messaging account holders."},
			// Other columns
			{Name: "business_profile", Type: proto.ColumnType_JSON, Description: "Business information about the account."},
			{Name: "business_type", Type: proto.ColumnType_STRING, Description: "The business type."},
			{Name: "capabilities", Type: proto.ColumnType_JSON, Description: "A hash containing the set of capabilities that was requested for this account and their associated states. Keys are names of capabilities. You can see the full list here. Values may be active, inactive, or pending."},
			{Name: "charges_enabled", Type: proto.ColumnType_BOOL, Description: "Whether the account can create live charges."},
			{Name: "company", Type: proto.ColumnType_JSON, Description: "Information about the company or business. This field is available for any business_type."},
			{Name: "country", Type: proto.ColumnType_STRING, Description: "The account’s country."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Created").Transform(transform.UnixToTimestamp), Description: "Time at which the account was created."},
			{Name: "default_currency", Type: proto.ColumnType_STRING, Description: "Three-letter ISO currency code representing the default currency for the account."},
			{Name: "deleted", Type: proto.ColumnType_BOOL, Description: "True if the customer is marked as deleted."},
			{Name: "details_submitted", Type: proto.ColumnType_BOOL, Description: "Whether account details have been submitted. Standard accounts cannot receive payouts before this is true."},
			{Name: "external_accounts", Type: proto.ColumnType_JSON, Description: "External accounts (bank accounts and debit cards) currently attached to this account."},
			{Name: "individual", Type: proto.ColumnType_JSON, Description: "Information about the person represented by the account. This field is null unless business_type is set to individual."},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Set of key-value pairs that you can attach to an account. This can be useful for storing additional information about the account in a structured format."},
			{Name: "payouts_enabled", Type: proto.ColumnType_BOOL, Description: "Whether Stripe can send payouts to this account."},
			{Name: "requirements", Type: proto.ColumnType_JSON, Description: "Information about the requirements for the account, including what information needs to be collected, and by when."},
			{Name: "settings", Type: proto.ColumnType_JSON, Description: "Options for customizing how the account functions within Stripe."},
			{Name: "tos_acceptance", Type: proto.ColumnType_JSON, Transform: transform.FromField("TOSAcceptance"), Description: "Details on the acceptance of the Stripe Services Agreement."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "The Stripe account type. Can be standard, express, or custom."},
		},
	}
}

func listAccount(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("stripe_account.listAccount", "connection_error", err)
		return nil, err
	}
	item, err := conn.Account.Get()
	if err != nil {
		plugin.Logger(ctx).Error("stripe_customer.listAccount", "query_error", err)
		return nil, err
	}
	d.StreamListItem(ctx, item)
	return nil, nil
}
