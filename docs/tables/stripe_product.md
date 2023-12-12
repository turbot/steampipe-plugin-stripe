---
title: "Steampipe Table: stripe_product - Query Stripe Products using SQL"
description: "Allows users to query Stripe Products, specifically product details like ID, name, description, active status, and more."
---

# Table: stripe_product - Query Stripe Products using SQL

Stripe Products is a resource within the Stripe payment processing service that represents the different products or services that a business offers. It provides a way to define and manage the catalog of items for which a business accepts payments. Stripe Products facilitates the organization of these items, including their pricing, description, and other related information.

## Table Usage Guide

The `stripe_product` table provides insights into the various products or services within Stripe's payment processing service. As a financial analyst or business owner, explore product-specific details through this table, including product ID, name, description, and active status. Utilize it to manage and organize your product catalog, monitor active products, and gain a comprehensive overview of your business offerings.

## Examples

### List all products
Explore all available products in your inventory to manage and oversee your product catalog more efficiently. This can help you keep track of your offerings and make strategic decisions based on the comprehensive overview.

```sql+postgres
select
  *
from
  stripe_product;
```

```sql+sqlite
select
  *
from
  stripe_product;
```

### All products that are not active
Explore which of your Stripe products are currently inactive. This can be useful for identifying items that may need updating or reactivation, helping to ensure your product offerings remain current and relevant.

```sql+postgres
select
  *
from
  stripe_product
where
  not active;
```

```sql+sqlite
select
  *
from
  stripe_product
where
  active = 0;
```