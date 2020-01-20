package blammo

type EventWriter interface{ WriteEvent(*Event) }

type NullWriter struct{}

func (w NullWriter) WriteEvent(*Event) {}
