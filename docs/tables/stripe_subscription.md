# Table: stripe_subscription

Subscriptions managed in the Stripe account.

## Examples

### List all subscriptions

```sql
select
  *
from
  stripe_subscription
```

### Subscriptions currently in the trial period

```sql
select
  *
from
  stripe_subscription
where
  status = 'trialing'
order by
  created desc
```

### Subscriptions set to cancel at the end of this period

```sql
select
  *
from
  stripe_subscription
where
  cancel_at_period_end
```

### Subscriptions created in the last 7 days

```sql
select
  *
from
  stripe_subscription
where
  created > current_timestamp - interval '7 days'
order by
  created
```
