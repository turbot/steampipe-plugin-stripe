![image](https://hub.steampipe.io/images/plugins/turbot/stripe-social-graphic.png)

# Stripe Plugin for Steampipe

Use SQL to query infrastructure including servers, networks, facilities and more from Stripe.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/stripe)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/stripe/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-stripe/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install stripe
```

Run a query:

```sql
select id, name, description from stripe_product;
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-stripe.git
cd steampipe-plugin-stripe
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/stripe.spc
```

Try it!

```
steampipe query
> .inspect stripe
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-stripe/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Stripe Plugin](https://github.com/turbot/steampipe-plugin-stripe/labels/help%20wanted)
