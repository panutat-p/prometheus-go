package internal

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Handler struct {
	Counter prometheus.Counter
}

func NewHandler() *Handler {
	return &Handler{
		Counter: promauto.NewCounter(prometheus.CounterOpts{
			Name: METRIC_HANDLER_NAME,
			Help: METRIC_HANDLER_DESCRIPTION,
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
			"metric":      METRIC_HANDLER_NAME,
			"description": METRIC_HANDLER_DESCRIPTION,
			"action":      "increase",
		},
	)
}
