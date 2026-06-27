package registry

// Collector defines the standard contract for all token usage trackers.
type Collector interface {
	Name() string
	Collect() error
}
