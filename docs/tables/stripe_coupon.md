---
title: "Steampipe Table: stripe_coupon - Query Stripe Coupons using SQL"
description: "Allows users to query Stripe Coupons, specifically the properties and details of each coupon, providing insights into discounts and promotions."
---

# Table: stripe_coupon - Query Stripe Coupons using SQL

Stripe Coupon is a feature within Stripe that allows you to create, manage, and distribute discount codes for your customers. It provides a flexible way to set up and manage coupons for various Stripe services, including subscriptions, invoices, and more. Stripe Coupon helps you attract new customers, reward loyal ones, or reactivate past customers with discounts and promotions.

## Table Usage Guide

The `stripe_coupon` table provides insights into coupons within Stripe. As a business owner or marketing manager, explore coupon-specific details through this table, including discount values, duration, and associated metadata. Utilize it to uncover information about coupons, such as those with high discount values, the duration of the coupons, and the verification of redemption limits.

## Examples

### List all coupons
Explore all the promotional coupons available in your system to understand the various discount schemes you offer. This can help in assessing the effectiveness of your marketing strategies and plan future campaigns.

```sql
select
  *
from
  stripe_coupon
```

### Coupons that are currently valid
Explore which coupons are currently active. This can be useful for understanding which promotional offers are available for customers at a given time.

```sql
select
  id,
  name
from
  stripe_coupon
where
  valid
```

### Coupons by popularity
Discover the segments that are most popular based on the number of times they've been redeemed. This can help prioritize marketing efforts and understand customer behavior.

```sql
select
  id,
  name,
  times_redeemed
from
  stripe_coupon
order by
  times_redeemed desc
```