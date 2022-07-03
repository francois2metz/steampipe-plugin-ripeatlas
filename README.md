# Ripe Atlas plugin for Steampipe

Use SQL to query probes, measurements and more from [RIPE Atlas][].

- **[Get started →](docs/index.md)**
- Documentation: [Table definitions & examples](docs/tables)

## Quick start

Install the plugin with [Steampipe][]:

    steampipe plugin install francois2metz/ripeatlas

## Development

To build the plugin and install it in your `.steampipe` directory

    make

Copy the default config file:

    cp config/ripeatlas.spc ~/.steampipe/config/ripeatlas.spc

## License

Apache 2

[steampipe]: https://steampipe.io
[RIPE Atlas]: https://atlas.ripe.net/
