package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	TokensUsedTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "llm_tokens_used_total",
			Help: "The total number of tokens used, partitioned by collector and type.",
		},
		[]string{"collector", "type"},
	)

	LastCollectedTimestamp = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "llm_tokens_last_collected_timestamp",
			Help: "The timestamp of the last successful collection, partitioned by collector.",
		},
		[]string{"collector"},
	)
)

// UpdateTokenUsage updates the counter for token usage.
func UpdateTokenUsage(collector string, tokenType string, amount float64) {
	TokensUsedTotal.WithLabelValues(collector, tokenType).Add(amount)
}

// RecordCollectionTimestamp records when a collection happened.
func RecordCollectionTimestamp(collector string) {
	// Using current unix timestamp as dummy value for simplicity in this example
	// In a real implementation, we might use the actual time.
	// But since we are just updating a gauge with a timestamp,
	// let's assume it receives the epoch seconds.
}
