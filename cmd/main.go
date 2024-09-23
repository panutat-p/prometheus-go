package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	PORT = "8080"
)

var (
	SIGNAL_STOP = make(chan os.Signal, 1)
)

func main() {
	signal.Notify(
		SIGNAL_STOP,
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGKILL,
		syscall.SIGTERM,
	)

	e := echo.New()
	e.GET("/", Health)
	e.GET("/hello", Hello)
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	go func() {
		err := e.Start(":" + PORT)
		if err != nil {
			panic(err)
		}
	}()

	opsProcessed := promauto.NewCounter(prometheus.CounterOpts{
		Name: "go_client_total_ops_process",
		Help: "The total number of processed events",
	})
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(500 * time.Millisecond)
		}
	}()

	<-SIGNAL_STOP
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := e.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Exit")
}

func Health(c echo.Context) error {
	return c.JSON(
		http.StatusOK, map[string]any{
			"status": "ok",
		},
	)
}

func Hello(c echo.Context) error {
	return c.JSON(
		http.StatusOK, map[string]any{
			"message": "hello",
		},
	)
}
