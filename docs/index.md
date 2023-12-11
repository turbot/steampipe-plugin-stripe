---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/stripe.svg"
brand_color: "#635BFF"
display_name: "Stripe"
short_name: "stripe"
description: "Steampipe plugin for querying customers, products, invoices and more from Stripe."
og_description: "Query Stripe with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/stripe-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Stripe + Steampipe

[Stripe](https://stripe.com) provides payment processing software and application programming interfaces (APIs) for e-commerce websites and mobile applications.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  id,
  name,
  description
from
  stripe_product
```

```
+---------------------+-----------------------+------------------------+
| id                  | name                  | description            |
+---------------------+-----------------------+------------------------+
| prod_Gm5Ugng0ZhPIxt | Airstream Deluxe A4   | The Cadillac of Paper. |
| prod_Hx0RdoxVyNJ3aC | 40 pound Letter Stock | <null>                 |
| prod_GnCKT8Km8WsOwK | Premium Copy Paper    | <null>                 |
+---------------------+-----------------------+------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/stripe/tables)**

## Get started

### Install

Download and install the latest Stripe plugin:

```bash
steampipe plugin install stripe
```

### Credentials

| Item        | Description                                                                  |
| :---------- | :--------------------------------------------------------------------------- |
| Credentials | Stripe requires an [API key](https://stripe.com/docs/keys) for all requests. |
| Radius      | Each connection represents a single Stripe account.                          |

### Configuration

Installing the latest stripe plugin will create a config file (`~/.steampipe/config/stripe.spc`) with a single connection named `stripe`:

```hcl
connection "stripe" {
  plugin   = "stripe"
  api_key = "sk_test_giG4MlyrcybGi1YFDEXAMPLE"
}
```

- `api_key` - Your Stripe API key for test or live data.

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-stripe
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
