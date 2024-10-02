package internal

import (
	"context"
	"time"
)

func StartCounters(ctx context.Context) {
	c1 := NewCounter("total_apple_count")
	c2 := NewCounter("total_banana_count")
	c3 := NewCounter("total_cherry_count")

	ticker1 := time.NewTicker(200 * time.Millisecond)
	ticker2 := time.NewTicker(300 * time.Millisecond)
	ticker3 := time.NewTicker(400 * time.Millisecond)
	defer ticker1.Stop()
	defer ticker2.Stop()
	defer ticker3.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker1.C:
			c1.Inc()
		case <-ticker2.C:
			c2.Inc()
		case <-ticker3.C:
			c3.Inc()
		}
	}
}

func StartCounterVecs(ctx context.Context) {
	c := NewCounterVec("total_fruit_count", []string{"name", "color"})

	ticker1 := time.NewTicker(200 * time.Millisecond)
	ticker2 := time.NewTicker(300 * time.Millisecond)
	ticker3 := time.NewTicker(400 * time.Millisecond)
	defer ticker1.Stop()
	defer ticker2.Stop()
	defer ticker3.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker1.C:
			c.Inc("apple", "red")
		case <-ticker2.C:
			c.Inc("banana", "yellow")
		case <-ticker3.C:
			c.Inc("cherry", "red")
		}
	}
}
