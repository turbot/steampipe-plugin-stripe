---
title: "Steampipe Table: stripe_plan - Query Stripe Plans using SQL"
description: "Allows users to query Stripe Plans, specifically the details of each pricing plan for products, providing insights into billing and subscription patterns."
---

# Table: stripe_plan - Query Stripe Plans using SQL

Stripe Plans is a service within Stripe that allows businesses to set up recurring payments for their products or services. It provides a way to create and manage different pricing tiers for subscriptions, including details about the billing cycle, pricing, and currency. Stripe Plans helps businesses to automate their billing process and manage subscriptions efficiently.

## Table Usage Guide

The `stripe_plan` table provides insights into the pricing plans within Stripe. As a billing manager, explore plan-specific details through this table, including pricing, intervals, and associated products. Utilize it to uncover information about plans, such as their cost, billing frequency, and the product they are associated with.

## Examples

### List all plans
Explore the different plans available in your Stripe account. This can be useful for understanding your billing structure and identifying any plans that may need to be updated or changed.

```sql
select
  *
from
  stripe_plan
```

### List all plans with a trial period
Discover the segments that offer trial periods in your subscription plans. This can help you assess the elements within your business model that may be contributing to customer acquisition and retention.

```sql
select
  id,
  nickname,
  trial_period_days
from
  stripe_plan
where
  trial_period_days > 0
```

### List all products with their associated plans
Explore which products are linked to specific plans, allowing you to analyze the relationship between your offerings and their associated plans for better product management and strategic planning.

```sql
select
  p.id,
  p.name,
  pl.id,
  pl.nickname
from
  stripe_product as p,
  stripe_plan as pl
where
  p.id = pl.product_id
order by
  p.name,
  pl.nickname
```