# Table: stripe_invoice

Invoices defined in the Stripe account.

## Examples

### List invoices

```sql
select
  *
from
  stripe_invoice
limit
  5
```

### Invoices created in the last 24 hours

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
