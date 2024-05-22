package main

import (
	"context"

	firewallv2 "github.com/metal-stack/firewall-controller-manager/api/v2"
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type firewallCollector struct {
	firewallInfo      *prometheus.Desc
	firewallCondition *prometheus.Desc
	lastRun           *prometheus.Desc

	c client.Client
}

func newfirewallCollector(c client.Client) *firewallCollector {
	return &firewallCollector{
		firewallInfo: prometheus.NewDesc("firewall_info",
			"Provide information about the metal-stack firewall",
			[]string{"name", "image", "size", "version"}, nil,
		),
		firewallCondition: prometheus.NewDesc("firewall_condition",
			"Provide information about the metal-stack firewall conditions",
			[]string{"name", "condition"}, nil,
		),
		lastRun: prometheus.NewDesc("firewall_last_run",
			"Provide the timestamp of the last run of the firewall collector",
			[]string{"name", "type"}, nil,
		),

		c: c}
}

func (collector *firewallCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.firewallInfo
	ch <- collector.firewallCondition
}

func (collector *firewallCollector) Collect(ch chan<- prometheus.Metric) {

	firewallList := &firewallv2.FirewallMonitorList{}
	err := collector.c.List(context.Background(), firewallList, &client.ListOptions{Namespace: "firewall"})
	if err != nil {
		panic(err)
	}
	for _, firewall := range firewallList.Items {
		ch <- prometheus.MustNewConstMetric(collector.firewallInfo, prometheus.GaugeValue, 1.0, firewall.Name, firewall.Image, firewall.Size, firewall.ControllerStatus.ControllerVersion)

		for _, condition := range firewall.Conditions {
			ch <- prometheus.MustNewConstMetric(collector.firewallCondition, prometheus.GaugeValue, string2bool(string(condition.Status)), firewall.Name, string(condition.Type))
		}

		ch <- prometheus.MustNewConstMetric(collector.lastRun, prometheus.GaugeValue, float64(firewall.ControllerStatus.Updated.Unix()), firewall.Name, "last-run")
		ch <- prometheus.MustNewConstMetric(collector.lastRun, prometheus.GaugeValue, float64(firewall.ControllerStatus.SeedUpdated.Unix()), firewall.Name, "last-run-against-seed")
	}

}

func string2bool(b string) float64 {
	if b == "True" {
		return 1.0
	}
	return 0.0
}
