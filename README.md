# firewall-exporter

A simple exporter for metal-stack firewalls.

## sample output

```text
# HELP firewall_info Provide information about the metal-stack firewall
# TYPE firewall_info gauge
firewall_info{name="shoot--pz9cjf--mwen-fel-firewall-aa1d3",image="firewall-ubuntu-3.0.20240202",size="n1-medium-x86",version="v2.3.2"} 1

# HELP firewall_condition Provide information about the metal-stack firewall conditions
# TYPE firewall_condition gauge
firewall_condition{condition="Connected",name="shoot--pz9cjf--mwen-fel-firewall-aa1d3"} 1
firewall_condition{condition="Created",name="shoot--pz9cjf--mwen-fel-firewall-aa1d3"} 1
firewall_condition{condition="Distance",name="shoot--pz9cjf--mwen-fel-firewall-aa1d3"} 1
firewall_condition{condition="MonitorDeployed",name="shoot--pz9cjf--mwen-fel-firewall-aa1d3"} 1
firewall_condition{condition="Ready",name="shoot--pz9cjf--mwen-fel-firewall-aa1d3"} 1
firewall_condition{condition="SeedConnected",name="shoot--pz9cjf--mwen-fel-firewall-aa1d3"} 1

# HELP firewall_last_run Provide the timestamp of the last run of the firewall collector
# TYPE firewall_last_run gauge
firewall_last_run{name="shoot--pz9cjf--mwen-fel-firewall-aa1d3",type="last-run"} 1.716383553e+09
firewall_last_run{name="shoot--pz9cjf--mwen-fel-firewall-aa1d3",type="last-run-against-seed"} 1.716383469e+09
```
