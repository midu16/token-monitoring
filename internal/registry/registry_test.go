package registry

import (
	"testing"
)

// MockCollector is a test implementation of the Collector interface.
type MockCollector struct {
	name           string
	collectCalled bool
}

func (m *MockCollector) Name() string {
	return m.name
}

func (m *MockCollector) Collect() error {
	m.collectCalled = true
	return nil
}

func TestRegistry_RegisterAndExecute(t *testing.T) {
	reg := NewRegistry()
	mock := &MockCollector{name: "test-collector"}

	err := reg.Register(mock)
	if err != nil {
		t.Fatalf("failed to register collector: %v", err)
	}

	err = reg.ExecuteAll()
	if err != nil {
		t.Fatalf("failed to execute collectors: %v", err)
	}

	if !mock.collectCalled {
		t.Error("expected Collect() to be *called* on the mock collector")
	}
}

func TestRegistry_DuplicateRegistration(t *testing.T) {
	reg := NewRegistry()
	mock1 := &MockCollector{name: "duplicate"}
	mock2 := &MockCollector{name: "duplicate"}

	_ = reg.Register(mock1)
	err := reg.Register(mock2)

	if err == nil {
		t.Error("expected error when registering duplicate collector name, got nil")
	}
}
