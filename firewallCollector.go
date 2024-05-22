package main

import (
	"context"

	firewallv2 "github.com/metal-stack/firewall-controller-manager/api/v2"
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type firewallCollector struct {
	firewallInfo *prometheus.Desc
	c            client.Client
}

func newfirewallCollector(c client.Client) *firewallCollector {
	return &firewallCollector{
		firewallInfo: prometheus.NewDesc("firwall_info",
			"Provide information about the metal-stack firewall",
			[]string{"image", "size", "version"}, nil,
		),
		c: c}
}

func (collector *firewallCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.firewallInfo
}

func (collector *firewallCollector) Collect(ch chan<- prometheus.Metric) {

	firewallList := &firewallv2.FirewallMonitorList{}
	err := collector.c.List(context.Background(), firewallList, &client.ListOptions{Namespace: "firewall"})
	if err != nil {
		panic(err)
	}
	for _, firewall := range firewallList.Items {
		ch <- prometheus.MustNewConstMetric(collector.firewallInfo, prometheus.GaugeValue, 1.0, firewall.Image, firewall.Size, firewall.ControllerStatus.ControllerVersion)
	}

}
