---
title: "Steampipe Table: stripe_subscription_item - Query Stripe Subscriptions using SQL"
description: "Query the items associaated with a Stripe Subscription."
---

# Table: stripe_subscription_item - Query Stripe Subscription Items using SQL

Stripe Subscriptions is a service within Stripe that allows businesses to manage recurring billing for their customers. Subscription Items provide details including price, quantity, and billing thresholds.

## Table Usage Guide

The `stripe_subscription_item` table in Steampipe provides detailed information about subscription records in your Stripe account. This table allows you, as a financial analyst or developer, to query subscription-specific details, including the status, customer, start and end dates, billing cycle, pricing plans, discounts, and metadata. You can leverage this table to analyze active or canceled subscriptions, monitor revenue streams, identify subscriptions with specific plans or discounts, and more. The schema outlines the attributes of a Stripe subscription, such as the subscription ID, customer information, pricing details, discount information, and metadata, enabling a comprehensive view of your subscription data.

**Important Notes**
- You must specify a `subscription_id` in a where or join clause in order to use this table.

## Examples

### List details for a subscription
Explore all the stripe subscriptions.

```sql+postgres
select
  *
from
  stripe_subscription_item
where
  subscription_id = 'sub_1Oo64zCWwOK68BLnfPDrQWIX'
```

```sql+sqlite
select
  *
from
  stripe_subscription_item
where
  subscription_id = 'sub_1Oo64zCWwOK68BLnfPDrQWIX'
```

### Retrieve subscription items with applied discounts
Identify subscription items from a specific subscription that have discounts applied.

```sql+postgres
select
  subscription_id,
  plan,
  price -> 'discounts' as discounts
from
  stripe_subscription_item
where
  subscription_id = 'sub_1Oo64zCWwOK68BLnfPDrQWIX'
  and (price -> 'discounts') is not null;
```

```sql+sqlite
select
  subscription_id,
  plan,
  json_extract(price, '$.discounts') as discounts
from
  stripe_subscription_item
where
  subscription_id = 'sub_1Oo64zCWwOK68BLnfPDrQWIX'
  and json_extract(price, '$.discounts') is not null;
```

### Retrieve plan details for a subscription item
Fetch detailed information about plans associated with a specific subscription, including their activity status, billing scheme, ID, usage type, and tiers.

```sql+postgres
select
  subscription_id,
  plan ->> 'active' as is_active,
  plan -> 'billingScheme' as billing_scheme,
  plan ->> 'id' as plan_id,
  plan ->> 'usageType' as usage_type,
  plan ->> 'tiers' as plan_tiers
from
  from
  stripe_subscription_item
where
  subscription_id = 'sub_1Oo64zCWwOK68BLnfPDrQWIX';
```

```sql+sqlite
select
  subscription_id,
  json_extract(plan, '$.active') as is_active,
  json_extract(plan, '$.billingScheme') as billing_scheme,
  json_extract(plan, '$.id') as plan_id,
  json_extract(plan, '$.usageType') as usage_type,
  json_extract(plan, '$.tiers') as plan_tiers
from
  stripe_subscription_item
where
  subscription_id = 'sub_1Oo64zCWwOK68BLnfPDrQWIX';
```