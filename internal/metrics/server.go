package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"ozonva/ova-competition-api/internal/config"
)

func RunMetricsServer(metricsConfig *config.MetricsConfig) error {
	mux := http.NewServeMux()
	mux.Handle(fmt.Sprintf("/%s", metricsConfig.MetricsPath), promhttp.Handler())

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", metricsConfig.MetricsPort),
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
