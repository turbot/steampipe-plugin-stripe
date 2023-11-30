---
title: "Steampipe Table: stripe_subscription - Query Stripe Subscriptions using SQL"
description: "Allows users to query Stripe Subscriptions, specifically the details of active, past, and upcoming subscriptions."
---

# Table: stripe_subscription - Query Stripe Subscriptions using SQL

Stripe Subscriptions is a service within Stripe that allows businesses to manage recurring billing for their customers. It provides a simplified way to set up and manage subscriptions, including various billing models, prorations, and tax rates. Stripe Subscriptions helps businesses automate their recurring billing, stay informed about subscription changes, and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `stripe_subscription` table provides insights into subscriptions within Stripe. As a financial analyst or a business owner, explore subscription-specific details through this table, including the status, items, and associated customer information. Utilize it to uncover information about subscriptions, such as those with upcoming invoices, the relationships between customers and their subscriptions, and the verification of billing details.

## Examples

### List all subscriptions
Explore all active subscriptions to understand the scope and scale of your service usage. This can aid in assessing revenue streams and identifying potential areas for growth or improvement.

```sql
select
  *
from
  stripe_subscription
```

### Subscriptions currently in the trial period
Explore which subscriptions are presently in their trial period to manage customer engagement and retention strategies effectively. This allows for a timely review of customer behavior and interactions during the trial phase.

```sql
select
  *
from
  stripe_subscription
where
  status = 'trialing'
order by
  created desc
```

### Subscriptions set to cancel at the end of this period
Identify the Stripe subscriptions that are scheduled to cancel at the end of the current billing period. This is useful for understanding your upcoming subscription churn and potentially reaching out to these customers to prevent cancellation.

```sql
select
  *
from
  stripe_subscription
where
  cancel_at_period_end
```

### Subscriptions created in the last 7 days
Discover the segments that have newly subscribed to your service within the past week. This can be useful in understanding recent customer behavior and trends.

```sql
select
  *
from
  stripe_subscription
where
  created > current_timestamp - interval '7 days'
order by
  created
```