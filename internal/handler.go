package internal

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Handler struct {
	MetricName string
	Counter    prometheus.Counter
}

func NewHandler() *Handler {
	return &Handler{
		MetricName: "total_http_request",
		Counter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "total_http_request",
			Help: "Counter for total_http_request",
		}),
	}
}

func (h *Handler) PrometheusMetrics(c echo.Context) error {
	promhttp.Handler().ServeHTTP(c.Response(), c.Request())
	return nil
}

func (h *Handler) Health(c echo.Context) error {
	return c.JSON(
		http.StatusOK, map[string]any{
			"status": "ok",
		},
	)
}

func (h *Handler) Increase(c echo.Context) error {
	h.Counter.Inc()
	return c.JSON(
		http.StatusOK, map[string]any{
			"metric": h.MetricName,
			"action": "increase",
		},
	)
}
