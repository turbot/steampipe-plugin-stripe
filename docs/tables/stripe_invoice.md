# Table: stripe_invoice

Invoices defined in the Stripe account.

## Examples

### List all invoices

```sql
select
  *
from
  stripe_invoice
```

### All unpaid invoices

```sql
select
  id,
  customer_name,
  amount_due,
  amount_paid,
  amount_outstanding
from
  stripe_invoice
where
  not paid
order by
  created
```
