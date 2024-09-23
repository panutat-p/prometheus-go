package internal

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Job struct {
	Counter prometheus.Counter
}

func NewJob() *Job {
	return &Job{
		Counter: promauto.NewCounter(prometheus.CounterOpts{
			Name: METRIC_JOB_NAME,
			Help: METRIC_JOB_DESCRIPTION,
		}),
	}
}

func (m *Job) Run() {
	for {
		m.Counter.Inc()
		time.Sleep(1 * time.Second)
	}
}
