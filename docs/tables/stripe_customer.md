# Table: stripe_customer

Customers defined in the Stripe account.

## Examples

### List all customers

```sql
select
  *
from
  stripe_customer
```

### Customers added in the last week

```sql
select
  id,
  name,
  created
from
  stripe_customer
where
  created > (current_timestamp - interval '7 days')
order by
  created desc
```

### All customers with a credit on their account

```sql
select
  id,
  name,
  account_balance,
  currency
from
  stripe_customer
where
  account_balance < 0
```

### All customers with an outstanding balance to add to their next invoice

```sql
select
  id,
  name,
  account_balance,
  currency
from
  stripe_customer
where
  account_balance > 0
```
