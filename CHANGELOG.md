## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#44](https://github.com/turbot/steampipe-plugin-stripe/pull/44))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#44](https://github.com/turbot/steampipe-plugin-stripe/pull/44))

## v0.6.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#36](https://github.com/turbot/steampipe-plugin-stripe/pull/36))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#36](https://github.com/turbot/steampipe-plugin-stripe/pull/36))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-stripe/blob/main/docs/LICENSE). ([#36](https://github.com/turbot/steampipe-plugin-stripe/pull/36))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#35](https://github.com/turbot/steampipe-plugin-stripe/pull/35))

## v0.5.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#27](https://github.com/turbot/steampipe-plugin-stripe/pull/27))

## v0.5.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#25](https://github.com/turbot/steampipe-plugin-stripe/pull/25))
- Recompiled plugin with Go version `1.21`. ([#25](https://github.com/turbot/steampipe-plugin-stripe/pull/25))

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
