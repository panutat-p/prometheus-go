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
	Label   []string
	Counter *prometheus.CounterVec
}

func NewCounterVec(name string, labels []string) *CounterVec {
	return &CounterVec{
		Name:  name,
		Label: labels,
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
	c.Counter.WithLabelValues(label).Inc()
}
