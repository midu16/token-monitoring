package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/user/token-monitoring/internal/collectors"
	"github.com/user/token-monitoring/internal/registry"
)

func main() {
	fmt.Println("Token Monitoring Server Starting...")

	// 1. Initialize Registry
	reg := registry.NewRegistry()

	// 2. Register Collectors
	httpCollector := collectors.NewHTTPInterceptor()
	logCollector := collectors.NewLogParser()

	reg.Register(httpCollector)
	reg.Register(logCollector)

	// 3. Start periodic collection in a background goroutine
	go func() {
		ticker := time.NewTicker(15 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fmt.Println("Executing periodic collection...")
				if err := reg.ExecuteAll(); err != nil {
					log.Printf("Error during collection: %v", err)
				}
			}
		}
	}()

	// 4. Setup HTTP Server for Prometheus metrics
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Metrics available at http://localhost:8081/metrics")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
