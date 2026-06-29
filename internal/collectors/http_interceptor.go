package collectors

import (
	"fmt"
	"math/rand"

	"github.com/user/token-monitoring/internal/exporter"
)

type HTTPInterceptor struct {
	name string
}

func NewHTTPInterceptor() *HTTPInterceptor {
	return &HTTPInterceptor{
		name: "http-interceptor",
	}
}

func (h *HTTPInterceptor) Name() string {
	return h.name
}

func (h *HTTPInterceptor) Collect() error {
	// Simulating finding usage in HTTP headers or body
	tokensUsed := float64(rand.Intn(100) + 1)
	tokenType := "completion" // placeholder

	exporter.UpdateTokenUsage(h.name, tokenType, tokensUsed)
	exporter.UpdateNLWP(h.name, float64(rand.Intn(10)+1))
	fmt.Printf("[Collector: %s] Detected %f tokens used\n", h.name, tokensUsed)

	return nil
}
