package blammo

type Tags struct {
	key    string
	value  interface{}
	parent *Tags
}

func (t *Tags) Tag(key string) *Tags {
	return &Tags{key: key, parent: t}
}

func (t *Tags) TagInt(key string, v int) *Tags {
	return &Tags{key: key, value: v, parent: t}
}

func (t *Tags) TagString(key, v string) *Tags {
	return &Tags{key: key, value: v, parent: t}
}

func (t *Tags) TagFloat(key string, v float64) *Tags {
	return &Tags{key: key, value: v, parent: t}
}
