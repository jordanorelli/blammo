package blammo

import (
	"time"
)

// Clock is used to allow for the injection of false clocks into log writers.
type Clock interface {
	Now() time.Time
}

// SystemClock is the default system clock
type SystemClock struct{}

// Now returns the current system time
func (c SystemClock) Now() time.Time { return time.Now() }

// FixedClock is a Clock implementation that always returns the same time.
type FixedClock time.Time

// Now always returns the same time
func (c FixedClock) Now() time.Time { return time.Time(c) }
