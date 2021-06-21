# Table: stripe_plan

Query information about plans defined in the Stripe account.

## Examples

### List all plans

```sql
select
  *
from
  stripe_plan
```

### List all plans with a trial period

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
