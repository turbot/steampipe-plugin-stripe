---
title: "Steampipe Table: stripe_invoice - Query Stripe Invoices using SQL"
description: "Allows users to query Stripe Invoices, specifically the details of all invoices created, paid, or attempted for a customer."
---

# Table: stripe_invoice - Query Stripe Invoices using SQL

Stripe is a payment processing platform that allows businesses to accept payments online and in mobile apps. It provides the technical, fraud prevention, and banking infrastructure required to operate online payment systems. Invoices in Stripe represent a bill for goods or services, and they are often tied to a subscription, but can also be used for one-off billing.

## Table Usage Guide

The `stripe_invoice` table provides insights into Stripe invoices within the Stripe Payment processing platform. As a financial analyst or business owner, explore invoice-specific details through this table, including amounts, currency, customer details, and payment status. Utilize it to uncover information about invoices, such as those unpaid, partially paid, or fully paid, and the details of the associated customer.

## Examples

### List invoices
Explore the most recent financial transactions by listing the latest five invoices. This is useful for getting a quick overview of recent billing activity.

```sql
select
  *
from
  stripe_invoice
limit
  5
```

### Invoices created in the last 24 hours
Discover the segments that have generated invoices in the last day. This can be used to gain insights into recent billing activity and customer engagement.

```sql
select
  id,
  customer_name,
  amount_due
from
  stripe_invoice
where
  created > current_timestamp - interval '1 day'
order by
  created
```

### Invoices due in a given month
Analyze the settings to understand which invoices are due within a specific month. This can help businesses to manage their cash flow by anticipating incoming payments.

```sql
select
  id,
  customer_name,
  amount_due
from
  stripe_invoice
where
  due_date >= '2021-07-01'
  and due_date < '2021-08-01'
```

### All open invoices
Explore which invoices are still open by identifying the outstanding amounts due and unpaid. This can help in tracking payments and managing financial records effectively.

```sql
select
  id,
  customer_name,
  amount_due,
  amount_paid,
  amount_remaining
from
  stripe_invoice
where
  status = 'outstanding'
order by
  created
```