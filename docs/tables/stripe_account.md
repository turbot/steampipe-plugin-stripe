---
title: "Steampipe Table: stripe_account - Query Stripe Accounts using SQL"
description: "Allows users to query Stripe Accounts, providing detailed information about the accounts associated with your Stripe platform."
---

# Table: stripe_account - Query Stripe Accounts using SQL

Stripe Account is a resource within the Stripe payment processing platform that represents an account on Stripe. These accounts can be associated with individual users or businesses and contain information about the account's details, balances, capabilities, and settings. They are integral to managing and understanding the financial transactions occurring on your Stripe platform.

## Table Usage Guide

The `stripe_account` table provides insights into the accounts within your Stripe platform. As a financial analyst or platform administrator, explore account-specific details through this table, including balances, capabilities, and settings. Utilize it to uncover information about each account, such as the account type, its capabilities, and the currencies it can handle.

## Examples

### List all accounts
Explore all your Stripe accounts to gain a comprehensive overview and better manage your online transactions. This could be particularly useful for businesses with multiple accounts, helping to streamline their financial operations.

```sql
select
  *
from
  stripe_account
```

### Are card payments active for this account?
Explore whether card payments are enabled for a given account. This can be useful for businesses to ensure they can process card payments and maintain smooth operations.

```sql
select
  id,
  (capabilities -> 'card_payments')::bool as card_payments
from
  stripe_account
```