package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var (
	//Initializing counter vector, which is collector that bundles a set of counters
	//that share same Desc (metric meta-data) but have different label vals

	//HttpReq Here we're declaring a one-dimensional counter vector for total number of http_requests
	//partitioned by their endpoint label (i.e. each http_requests_total counter
	//instance will differ in value of endpoint label)
	HttpReq = prometheus.NewCounterVec(prometheus.CounterOpts{
		//Desc data
		Name: "http_requests_total",
		Help: "How many http requests processed, partitioned by endpoint",
	},
		//label names
		[]string{"endpoint"},
	)

	FirstServiceCount, SecondServiceCount prometheus.Counter

	// TasksQueued Number of tasks queued, using gauge since this can go up or down
	TasksQueued = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tasks_queued",
		Help: "Number of tasks queued at any given time",
	})
)

func init() {
	//Registering collectors, which allows metrics to be exposed
	//via HTTP server
	prometheus.MustRegister(HttpReq)
	prometheus.MustRegister(TasksQueued)

	//Creating handles for specific label values
	FirstServiceCount = HttpReq.WithLabelValues("/first_service")
	SecondServiceCount = HttpReq.WithLabelValues("/second_service")

}
func Start() {

	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":6060", nil))

}
