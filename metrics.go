package main

import (
	"context"
	"net/http"
	"time"

<<<<<<< HEAD
	pm "github.com/deathowl/go-metrics-prometheus"
	"github.com/prometheus/client_golang/prometheus"
=======
	promMetrics "github.com/MeteoGroup/go-metrics-prometheus"
>>>>>>> origin/master
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/trivago/gollum/core"
)

func startPrometheusMetricsService(address string) func() {
<<<<<<< HEAD
	registry := prometheus.NewRegistry()
	srv := &http.Server{Addr: address}
	quit := make(chan struct{})

	// Start updates
	go func() {
		client := pm.NewPrometheusProvider(core.MetricsRegistry, "gollum", "", registry, 0)
		for {
			select {
			case <-time.After(time.Second):
				client.UpdatePrometheusMetricsOnce()
=======
	srv := &http.Server{Addr: address}
	quit := make(chan struct{})

	flushInterval := 3 * time.Second
	promClient := promMetrics.NewPrometheusProvider(core.MetricsRegistry, "gollum", "", flushInterval)

	// Start updates
	go func() {
		for {
			select {
			case <-time.After(flushInterval):
				if err := promClient.UpdatePrometheusMetricsOnce(); err != nil {
					logrus.WithError(err).Warn("Error updating metrics")
				}
>>>>>>> origin/master
			case <-quit:
				return
			}
		}
	}()

	// Start http
	go func() {
		opts := promhttp.HandlerOpts{
			ErrorLog:      logrus.StandardLogger(),
			ErrorHandling: promhttp.ContinueOnError,
		}
<<<<<<< HEAD
		http.Handle("/prometheus", promhttp.HandlerFor(registry, opts))
=======
		http.Handle("/prometheus", promhttp.HandlerFor(promClient.PromRegistry, opts))
>>>>>>> origin/master

		err := srv.ListenAndServe()
		if err != nil {
			logrus.WithError(err).Error("Failed to start metrics http server")
		}
	}()

	logrus.WithField("address", address).Info("Started metric service")

	// Return stop function
	return func() {
		close(quit)
		if err := srv.Shutdown(context.Background()); err != nil {
			logrus.WithError(err).Error("Failed to shutdown metrics http server")
		}
	}
}
