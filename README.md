mackerel-plugin-unbound
=======================

[Unbound](https://nlnetlabs.nl/projects/unbound/about/) custom metrics plugin for mackerel.io agent.

## Synopsis

```shell
mackerel-plugin-unbound [-unbound-control=<path to unbound-control>] [-tempfile=<tempfile>]
```

## Example of mackerel-agent.conf

```toml
[plugin.metrics.unbound]
command = "/path/to/mackerel-plugin-unbound"
```
