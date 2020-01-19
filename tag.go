package blammo

type Tags struct {
	key    string
	value  interface{}
	parent *Tags
}
