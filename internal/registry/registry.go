package registry

import (
	"fmt"
	"sync"
)

// Registry manages a set of active collectors and aggregates their results.
type Registry struct {
	mu         sync.RWMutex
	collectors []Collector
}

// NewRegistry creates a new instance of the Registry.
func NewRegistry() *Registry {
	return &Registry{
		collectors: []Collector{},
	}
}

// Register adds a new collector to the registry.
func (r *Registry) Register(c Collector) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, existing := range r.collectors {
		if existing.Name() == c.Name() {
			return fmt.Errorf("collector with name %s already registered", c.Name())
		}
	}
	r.collectors = append(r.collectors, c)
	return nil
}

// ExecuteAll executes the Collect method on all registered collectors.
func (r *Registry) ExecuteAll() error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, c := range r.collectors {
		if err := c.Collect(); err != nil {
			return fmt.Errorf("collector %s failed: %w", c.Name(), err)
		}
	}
	return nil
}

// Collectors returns a list of all registered collectors.
func (r *Registry) Collectors() []Collector {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.collectors
}
