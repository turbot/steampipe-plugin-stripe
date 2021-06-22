# Table: stripe_coupon

Coupons issued in the Stripe account.

## Examples

### List all coupons

```sql
select
  *
from
  stripe_coupon
```

### Coupons that are currently valid

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
