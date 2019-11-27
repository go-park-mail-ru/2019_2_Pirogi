package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
)

type Metrics struct {
	HitsTotal prometheus.Counter
	Hits      *prometheus.CounterVec
}

var ApiMetrics Metrics

func InitMetrics() {
	ApiMetrics.HitsTotal = prometheus.NewCounter(prometheus.CounterOpts{Name: "hits_total"})
	ApiMetrics.Hits = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "hits"}, []string{"status", "method", "path"})
	prometheus.MustRegister(ApiMetrics.HitsTotal, ApiMetrics.Hits)
}

func (m *Metrics) IncHitsTotal() {
	m.HitsTotal.Inc()
}

func (m *Metrics) IncHitOfResponse(status int, method, path string) {
	m.Hits.WithLabelValues(strconv.Itoa(status), method, path).Inc()
}
