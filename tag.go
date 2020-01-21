package blammo

// Tags represent event metadata. Tags is an ordered collection of key-value
// pairs where every key is a string. Values are optional, and may be one of
// int, float64, and string. The same key may appear multiple times.
//
// Tags is internally represented by a singly linked list of key-value pairs. A
// Tags struct is immutable once created. The only thing that can be done is to
// create new nodes that point to old nodes. This allows us to continually
// recycle tag nodes that are widely used across an application.
type Tags struct {
	key    string
	value  interface{}
	parent *Tags
}
