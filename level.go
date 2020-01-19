package blammo

type Level uint

const (
	// Debug is intended to be used for verbose logging information of
	// implementation details.
	Debug Level = iota

	// Info is intended to be used to report expected behaviors; it's used to
	// log usage and observe normal behaviors.
	Info

	// Warn is intended to be used to report events that are not along the
	// expected "happy path" of the application. These events should generally
	// represent failures of -other- systems.
	Warn

	// Error is intended to b e used to report things that the application was
	// not able to handle. These events should generally represent failures of
	// the system at hand.
	Error
)
