# Table: stripe_account

Get Stripe account information for the authenticated caller.

## Examples

### List all accounts

```sql
select
  *
from
  stripe_account
```

### Are card payments active for this account?

```sql
select
  id,
  (capabilities -> 'card_payments')::bool as card_payments
from
  stripe_account
```
