# firewall-exporter

A simple exporter for metal-stack firewalls.

## sample output

```text
# HELP firwall_info Provide information about the metal-stack firewall
# TYPE firwall_info gauge
firwall_info{image="firewall-ubuntu-3.0.20240202",size="n1-medium-x86",version="v2.3.2"} 1
```
