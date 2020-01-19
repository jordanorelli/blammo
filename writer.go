package blammo

type EventWriter interface {
	WriteEvent(*Event) error
}
