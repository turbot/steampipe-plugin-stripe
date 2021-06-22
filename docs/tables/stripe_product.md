# Table: stripe_product

Query information about products defined in the Stripe account.

## Examples

### List all products

```sql
select
  *
from
  stripe_product
```

### All products that are not active

```sql
select
  *
from
  stripe_product
where
  not active
```
