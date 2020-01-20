package blammo

import (
	"time"
)

// Event is a single log record. A log line if you're writing to file, a
// database row if you're writing to a database, etc. Everything internally is
// expressed as an event.
//
// Event is exported to support the implementation of custom log writers. Most
// users should not need to handle this type directly.
type Event struct {
	// severity of the event
	Level Level

	// time at which the event occured
	Time time.Time

	// where the event occurred in the system
	Path *Path

	// message to be logged
	Text string

	// key-value pairs to log as extra metadata
	Tags *Tags
}
