package internal

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func StartCounter(ctx context.Context, name string, delay time.Duration) {
	c := promauto.NewCounter(prometheus.CounterOpts{
		Name: name,
		Help: "Counter for " + name,
	})
	for {
		c.Inc()
		time.Sleep(delay)
	}
}

func StartCounterVec(ctx context.Context, name string, label string, delay time.Duration) {
	c := promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: name,
			Help: "CounterVec for " + name,
		},
		[]string{label},
	)
	for {
		c.WithLabelValues(label).Inc()
		time.Sleep(delay)
	}
}
