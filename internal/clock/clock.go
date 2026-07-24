package clock

import (
	"sync"
	"time"
)

var (
	mu     sync.RWMutex
	offset time.Duration
)

// Now returns the current time, applying any virtual offset.
func Now() time.Time {
	mu.RLock()
	defer mu.RUnlock()
	return time.Now().Add(offset)
}

// Advance fast-forwards the global virtual clock by the specified duration.
func Advance(d time.Duration) {
	mu.Lock()
	defer mu.Unlock()
	offset += d
}

// Reset clears any virtual offset.
func Reset() {
	mu.Lock()
	defer mu.Unlock()
	offset = 0
}

// Offset returns the current virtual time offset.
func Offset() time.Duration {
	mu.RLock()
	defer mu.RUnlock()
	return offset
}
