package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"

	"prometheus_go/internal"
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
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGKILL,
		syscall.SIGTERM,
	)
	defer cancel()

	h := internal.NewHandler()

	e := echo.New()
	e.GET("/", h.Health)
	e.GET("/metrics", h.PrometheusMetrics)
	e.GET("/increase", h.Increase)

	go func() {
		err := e.Start(":" + PORT)
		if err != nil {
			fmt.Println("[Echo]", err)
		}
	}()

	go internal.StartCounters(ctx)
	go internal.StartCounterVecs(ctx)

	<-SIGNAL_STOP
	{
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := e.Shutdown(ctx)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Exit")
}
