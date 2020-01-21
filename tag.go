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

// Tag creates a new Tags struct having the given key as its final key, with no
// associated value. All existing keys continue to exist.
func (t *Tags) Tag(key string) *Tags {
	return &Tags{key: key, parent: t}
}

// TagInt creates a new Tags struct having the given key-value pair as the
// final key-value pair. All existing key-value pairs continue to exist.
func (t *Tags) TagInt(key string, v int) *Tags {
	return &Tags{key: key, value: v, parent: t}
}

// TagString creates a new Tags struct having the given key-value pair as the
// final key-value pair. All existing key-value pairs continue to exist.
func (t *Tags) TagString(key, v string) *Tags {
	return &Tags{key: key, value: v, parent: t}
}

// TagFloat creates a new Tags struct having the given key-value pair as the
// final key-value pair. All existing key-value pairs continue to exist.
func (t *Tags) TagFloat(key string, v float64) *Tags {
	return &Tags{key: key, value: v, parent: t}
}
