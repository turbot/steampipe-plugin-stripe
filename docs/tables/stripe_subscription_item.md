---
title: "Steampipe Table: stripe_subscription_item - Query Stripe Subscriptions using SQL"
description: "Query the items associaated with a Stripe Subscription."
---

# Table: stripe_subscription_item - Query Stripe Subscription Items using SQL

Stripe Subscriptions is a service within Stripe that allows businesses to manage recurring billing for their customers. Subscription Items provide details including price, quantity, and billing thresholds.

## Examples

### List details for a subscription

```sql+postgres
select
  *
from
  stripe_subscription
where
  subscription_id = 'sub_1Oo64zCWwOK68BLnfPDrQWIX'
```
