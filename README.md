# firewall-exporter

A simple exporter for metal-stack firewalls.

## sample output

```text
# HELP firewall_info Provide information about the metal-stack firewall
# TYPE firewall_info gauge
firewall_info{image="firewall-ubuntu-3.0.20240202",size="n1-medium-x86",version="v2.3.2"} 1
```
