package internal

import (
	"context"
	"time"
)

func StartCounters(ctx context.Context) {
	c1 := NewCounter("total_apple")
	c2 := NewCounter("total_banana")
	c3 := NewCounter("total_cherry")
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(200 * time.Millisecond):
			c1.Inc()
		case <-time.After(300 * time.Millisecond):
			c2.Inc()
		case <-time.After(400 * time.Millisecond):
			c3.Inc()
		}
	}
}

func StartCounterVecs(ctx context.Context) {
	c := NewCounterVec("total_fruit", []string{"apple", "banana", "cherry"})
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(200 * time.Millisecond):
			c.Inc("apple")
		case <-time.After(300 * time.Millisecond):
			c.Inc("banana")
		case <-time.After(400 * time.Millisecond):
			c.Inc("cherry")
		}
	}
}
