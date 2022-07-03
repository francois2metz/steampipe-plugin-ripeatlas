---
organization: francois2metz
category: ["Internet"]
brand_color: "#080037"
display_name: "RIPE Atlas"
short_name: "ripeatlas"
description: "Steampipe plugin for querying probes from RIPE Atlas."
og_description: "Query Scalingo with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/francois2metz/ripeatlas-social-graphic.png"
icon_url: "/images/plugins/francois2metz/ripeatlas.svg"
---

# RIPE Atlas + Steampipe

[RIPE Atlas][https://atlas.ripe.net/] employs a global network of probes that measure Internet connectivity and reachability, providing an unprecedented understanding of the state of the Internet in real time.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  address_v4,
  is_public
from
  ripeatlas_probe
where
  id=1
```

```
+---------------+-----------+
| address_v4    | is_public |
+---------------+-----------+
| 45.138.229.91 | true      |
+---------------+-----------+
```

## Documentation

- **[Table definitions & examples →](/plugins/francois2metz/ripeatlas/tables)**

## Get started

### Install

Download and install the latest RIPE Atlas plugin:

```bash
steampipe plugin install francois2metz/ripeatlas
```

### Configuration

Installing the latest ripe atlas plugin will create a config file (`~/.steampipe/config/ripeatlas.spc`) with a single connection named `ripeatlas`:

```hcl
connection "ripeatlas" {
    plugin = "francois2metz/ripeatlas"

    # API Key
    # Get it on: https://atlas.ripe.net/keys/
    # key = "00000000-0000-0000-0000-000000000000"
}
```

You can also use environment variables:

- `RIPEATLAS_KEY` for the KEY (ex: 00000000-0000-0000-0000-000000000000)

## Get Involved

* Open source: https://github.com/francois2metz/steampipe-plugin-ripeatlas
