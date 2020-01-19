package blammo

type Codec interface {
	Encode(*Event, []byte) error
	Decode(*Event, []byte) error
}
