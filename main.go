package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"

	klog "k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	firewallv2 "github.com/metal-stack/firewall-controller-manager/api/v2"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var (
	scheme = runtime.NewScheme()
)

func main() {

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	cs, err := client.New(config, client.Options{Scheme: scheme})
	if err != nil {
		panic(err.Error())
	}
	metalCollector := newfirewallCollector(cs)
	prometheus.MustRegister(metalCollector)

	http.Handle("/metrics", promhttp.Handler())
	klog.Info("Beginning to serve on port :9080")
	server := &http.Server{
		Addr:              ":9080",
		ReadHeaderTimeout: 1 * time.Minute,
	}
	klog.Fatal(server.ListenAndServe())
}

func init() {
	sb := runtime.NewSchemeBuilder(firewallv2.AddToScheme)
	utilruntime.Must(sb.AddToScheme(scheme))
}
