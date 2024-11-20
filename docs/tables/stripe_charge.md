---
title: "Steampipe Table: stripe_charge - Query Stripe Charges using SQL"
description: "Allows users to query Stripe charges, providing detailed information on charge amounts, payment methods, customers, and more."
---

# Table: stripe_charge - Query Stripe Charges using SQL

Stripe charges represent payments made by customers to businesses through the Stripe platform. Each charge contains detailed information about the transaction, including amounts, currency, customer details, payment methods, and dispute status. The `stripe_charge` table in Steampipe enables you to query and analyze charges, helping you manage transactions, refunds, and customer details.

## Table Usage Guide

The `stripe_charge` table is useful for finance teams, data analysts, and developers who need insights into Stripe charge data. You can query various attributes such as charge amounts, payment methods, refunds, and customer details. This table is especially helpful for monitoring transactions, analyzing payment trends, managing disputes, and ensuring payment processes are functioning as expected.

## Examples

### Basic charge information
Retrieve basic information about Stripe charges, including the amount, currency, and status.

```sql+postgres
select
  id,
  amount,
  currency,
  created,
  status
from
  stripe_charge;
```

```sql+sqlite
select
  id,
  amount,
  currency,
  created,
  status
from
  stripe_charge;
```

### List charges by customer
Retrieve charges made by a specific customer, identified by the `customer` field.

```sql+postgres
select
  id,
  customer,
  amount,
  currency,
  status
from
  stripe_charge
where
  customer = 'cus_12345ABC';
```

```sql+sqlite
select
  id,
  customer,
  amount,
  currency,
  status
from
  stripe_charge
where
  customer = 'cus_12345ABC';
```

### List refunded charges
Identify charges that have been refunded and retrieve the refund amount.

```sql+postgres
select
  id,
  amount,
  amount_refunded,
  refunded,
  created
from
  stripe_charge
where
  refunded = true;
```

```sql+sqlite
select
  id,
  amount,
  amount_refunded,
  refunded,
  created
from
  stripe_charge
where
  refunded = 1;
```

### List charges by payment method
Fetch charges based on the payment method used, such as credit card or bank transfer.

```sql+postgres
select
  id,
  payment_method,
  amount,
  currency,
  status
from
  stripe_charge
where
  payment_method = 'pm_123456789';
```

```sql+sqlite
select
  id,
  payment_method,
  amount,
  currency,
  status
from
  stripe_charge
where
  payment_method = 'pm_123456789';
```

### List disputed charges
Identify charges that are disputed and view the dispute details.

```sql+postgres
select
  id,
  amount,
  currency,
  disputed,
  dispute
from
  stripe_charge
where
  disputed = true;
```

```sql+sqlite
select
  id,
  amount,
  currency,
  disputed,
  dispute
from
  stripe_charge
where
  disputed = 1;
```