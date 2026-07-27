// Package billing implements a deterministic cost ledger and pricing model that
// drive kiri's billing and time-travel simulations.
package billing

import (
	"sync"
	"time"

	"github.com/Brilhante29/kiri-aws/internal/clock"
)

// Resource represents a cloud-agnostic resource.
type Resource struct {
	ARN          string
	Provider     string
	Service      string
	ResourceType string
	CreatedAt    time.Time
	DeletedAt    *time.Time
}

// UsageEvent represents a discrete usage event (e.g. data transfer, API request).
type UsageEvent struct {
	ResourceARN string
	MetricName  string
	Value       float64
	Timestamp   time.Time
}

// Cost represents the monetary value.
type Cost struct {
	Amount   float64
	Currency string
}

// Ledger holds the lifecycle of resources and their usage events.
type Ledger struct {
	mu        sync.RWMutex
	resources map[string]*Resource
	usage     []UsageEvent
}

var globalLedger = &Ledger{
	resources: make(map[string]*Resource),
}

// Global returns the singleton ledger instance.
func Global() *Ledger {
	return globalLedger
}

// Reset clears the global ledger (mostly for testing).
func Reset() {
	globalLedger.mu.Lock()
	defer globalLedger.mu.Unlock()
	globalLedger.resources = make(map[string]*Resource)
	globalLedger.usage = nil
}

// RecordResourceCreated marks a resource as created in the ledger.
func (l *Ledger) RecordResourceCreated(arn, provider, serviceName, resourceType string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.resources[arn] = &Resource{
		ARN:          arn,
		Provider:     provider,
		Service:      serviceName,
		ResourceType: resourceType,
		CreatedAt:    clock.Now(),
	}
}

// RecordResourceDeleted marks a resource as deleted.
func (l *Ledger) RecordResourceDeleted(arn string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if res, ok := l.resources[arn]; ok {
		now := clock.Now()
		res.DeletedAt = &now
	}
}

// RecordUsage increments a specific metric for a resource.
func (l *Ledger) RecordUsage(arn, metricName string, value float64) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.usage = append(l.usage, UsageEvent{
		ResourceARN: arn,
		MetricName:  metricName,
		Value:       value,
		Timestamp:   clock.Now(),
	})
}
