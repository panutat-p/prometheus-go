package internal

import (
	"slices"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Counter struct {
	Name    string
	Counter prometheus.Counter
}

func NewCounter(name string) *Counter {
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
	Labels  []string
	Counter *prometheus.CounterVec
}

func NewCounterVec(name string, labels []string) *CounterVec {
	return &CounterVec{
		Name:   name,
		Labels: labels,
		Counter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: name,
				Help: "CounterVec for " + name,
			},
			labels,
		),
	}
}

func (c *CounterVec) Inc(label string) {
	if !slices.Contains(c.Labels, label) {
		panic("invalid label: " + label)
	}
	c.Counter.WithLabelValues(label).Inc()
}
