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

### List all products that are not active

```sql
select
  *
from
  stripe_product
where
  not active
```
