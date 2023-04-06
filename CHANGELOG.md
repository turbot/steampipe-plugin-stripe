## v0.4.0 [2023-04-06]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#21](https://github.com/turbot/steampipe-plugin-stripe/pull/21))

## v0.3.1 [2022-10-03]

_Bug fixes_

- Fixed `account_balance` -> `balance` column names in `stripe_customer` table document examples. ([#19](https://github.com/turbot/steampipe-plugin-stripe/pull/19)) (Thanks to [@fabpot](https://github.com/fabpot) for the fixes!)

## v0.3.0 [2022-09-28]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#17](https://github.com/turbot/steampipe-plugin-stripe/pull/17))
- Recompiled plugin with Go version `1.19`. ([#16](https://github.com/turbot/steampipe-plugin-stripe/pull/16))

## v0.2.1 [2022-05-23]

_Bug fixes_

- Fixed the Slack community links in README and docs/index.md files. ([#12](https://github.com/turbot/steampipe-plugin-stripe/pull/12))

## v0.2.0 [2022-04-28]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#10](https://github.com/turbot/steampipe-plugin-stripe/pull/10))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#9](https://github.com/turbot/steampipe-plugin-stripe/pull/9))

## v0.1.0 [2021-11-23]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) and Go version 1.17 ([#5](https://github.com/turbot/steampipe-plugin-stripe/pull/5))

## v0.0.2 [2021-09-22]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v161--2021-09-21) ([#2](https://github.com/turbot/steampipe-plugin-stripe/pull/2))

## v0.0.1 [2021-07-12]

_What's new?_

- New tables added
  - [stripe_account](https://hub.steampipe.io/plugins/turbot/stripe/tables/stripe_account)
  - [stripe_coupon](https://hub.steampipe.io/plugins/turbot/stripe/tables/stripe_coupon)
  - [stripe_customer](https://hub.steampipe.io/plugins/turbot/stripe/tables/stripe_customer)
  - [stripe_invoice](https://hub.steampipe.io/plugins/turbot/stripe/tables/stripe_invoice)
  - [stripe_plan](https://hub.steampipe.io/plugins/turbot/stripe/tables/stripe_plan)
  - [stripe_product](https://hub.steampipe.io/plugins/turbot/stripe/tables/stripe_product)
  - [stripe_subscription](https://hub.steampipe.io/plugins/turbot/stripe/tables/stripe_subscription)
