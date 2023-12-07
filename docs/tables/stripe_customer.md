---
title: "Steampipe Table: stripe_customer - Query Stripe Customers using SQL"
description: "Allows users to query Stripe Customers, providing insights into customer data including payment methods, transactions, and associated metadata."
---

# Table: stripe_customer - Query Stripe Customers using SQL

Stripe is a technology company that builds economic infrastructure for the internet. Businesses of every size from new startups to public companies use the software to accept payments and manage their businesses online. Stripe combines a payments platform with applications that put revenue data at the heart of business operations.

## Table Usage Guide

The `stripe_customer` table provides insights into customer data within Stripe. As a financial analyst, you can explore customer-specific details through this table, including payment methods, transactions, and associated metadata. Utilize it to uncover information about customers, such as their payment history, preferred payment methods, and transaction patterns.

## Examples

### List all customers
Explore all customer data to gain insights into your customer base and understand their activity and behavior better. This can assist in creating targeted marketing strategies and improving customer service.

```sql+postgres
select
  *
from
  stripe_customer;
```

```sql+sqlite
select
  *
from
  stripe_customer;
```

### Customers added in the last week
Discover the segments that represent new customers by identifying those who have been added in the past week. This can help businesses understand recent growth patterns and tailor their marketing efforts accordingly.

```sql+postgres
select
  id,
  name,
  created
from
  stripe_customer
where
  created > (current_timestamp - interval '7 days')
order by
  created desc;
```

```sql+sqlite
select
  id,
  name,
  created
from
  stripe_customer
where
  created > (strftime('%s','now') - 7*24*60*60)
order by
  created desc;
```

### All customers with a credit on their account
Discover the segments that consist of customers who have a credit balance on their account. This is useful to identify potential areas for revenue recovery or customer engagement.

```sql+postgres
select
  id,
  name,
  balance,
  currency
from
  stripe_customer
where
  balance < 0;
```

```sql+sqlite
select
  id,
  name,
  balance,
  currency
from
  stripe_customer
where
  balance < 0;
```

### All customers with an outstanding balance to add to their next invoice
Gain insights into customers who have a pending balance, which will be added to their next invoice. This helps in understanding the financial standing of the customers and planning future billing accordingly.

```sql+postgres
select
  id,
  name,
  balance,
  currency
from
  stripe_customer
where
  balance > 0;
```

```sql+sqlite
select
  id,
  name,
  balance,
  currency
from
  stripe_customer
where
  balance > 0;
```