package internal

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Counter struct {
	Name    string
	Counter prometheus.Counter
}

func NewCounter(name string) *Counter {
	if name == "" {
		panic("empty name")
	}
	return &Counter{
		Name: name,
		Counter: promauto.NewCounter(prometheus.CounterOpts{
			Name: name,
			Help: "Counter for " + name,
		}),
	}
}

func (c *Counter) Inc() {
	c.Counter.Inc()
}

type CounterVec struct {
	Name    string
	Labels  map[string]struct{}
	Counter *prometheus.CounterVec
}

func NewCounterVec(name string, labels []string) *CounterVec {
	m := make(map[string]struct{})
	for _, label := range labels {
		if label == "" {
			panic("empty label")
		}
		m[label] = struct{}{}
	}
	return &CounterVec{
		Name:   name,
		Labels: m,
		Counter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: name,
				Help: "CounterVec for " + name,
			},
			labels,
		),
	}
}

func (c *CounterVec) Inc(labels ...string) {
	c.Counter.WithLabelValues(labels...).Inc()
}
