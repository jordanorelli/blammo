package blammo

import (
	"fmt"
)

type Log struct {
	dw    EventWriter
	iw    EventWriter
	ew    EventWriter
	path  *Path
	clock Clock
	tags  *Tags
}

func NewLog(name string, options ...Option) *Log {
	l := &Log{
		path: NewPath(name),
	}
	for _, opt := range options {
		opt.apply(l)
	}
	return l
}

func format(t string, args ...interface{}) string {
	if len(args) == 0 {
		return t
	}
	return fmt.Sprintf(t, args...)
}

func (l *Log) event(lvl Level, t string, args ...interface{}) *Event {
	return &Event{
		Time:  l.clock.Now(),
		Level: lvl,
		Path:  l.path,
		Tags:  l.tags,
		Text:  format(t, args...),
	}
}

func (l *Log) Debug(t string, args ...interface{}) {
	if l.dw != nil {
		l.dw.WriteEvent(l.event(Debug, t, args...))
	}
}

func (l *Log) Info(t string, args ...interface{}) {
	if l.iw != nil {
		l.iw.WriteEvent(l.event(Info, t, args...))
	}
}

func (l *Log) Error(t string, args ...interface{}) {
	if l.ew != nil {
		l.ew.WriteEvent(l.event(Error, t, args...))
	}
}

func (l *Log) Child(name string) *Log {
	return &Log{
		dw:    l.dw,
		iw:    l.iw,
		ew:    l.ew,
		path:  l.path.Child(name),
		clock: l.clock,
		tags:  l.tags,
	}
}

// Tag creates a new Tags struct having the given key as its final key, with no
// associated value. All existing keys continue to exist.
func (l *Log) Tag(key string) *Log {
	return &Log{
		dw:    l.dw,
		iw:    l.iw,
		ew:    l.ew,
		path:  l.path,
		clock: l.clock,
		tags:  &Tags{key: key, parent: l.tags},
	}
}

// TagInt creates a new Tags struct having the given key-value pair as the
// final key-value pair. All existing key-value pairs continue to exist.
func (l *Log) TagInt(key string, v int) *Log {
	return &Log{
		dw:    l.dw,
		iw:    l.iw,
		ew:    l.ew,
		path:  l.path,
		clock: l.clock,
		tags:  &Tags{key: key, value: v, parent: l.tags},
	}
}

// TagString creates a new Tags struct having the given key-value pair as the
// final key-value pair. All existing key-value pairs continue to exist.
func (l *Log) TagString(key, v string) *Log {
	return &Log{
		dw:    l.dw,
		iw:    l.iw,
		ew:    l.ew,
		path:  l.path,
		clock: l.clock,
		tags:  &Tags{key: key, value: v, parent: l.tags},
	}
}

// TagFloat creates a new Tags struct having the given key-value pair as the
// final key-value pair. All existing key-value pairs continue to exist.
func (l *Log) TagFloat(key string, v float64) *Log {
	return &Log{
		dw:    l.dw,
		iw:    l.iw,
		ew:    l.ew,
		path:  l.path,
		clock: l.clock,
		tags:  &Tags{key: key, value: v, parent: l.tags},
	}
}

type Option interface{ apply(*Log) }

type optionFn func(*Log)

func (fn optionFn) apply(l *Log) { fn(l) }

func DebugWriter(w EventWriter) Option { return optionFn(func(l *Log) { l.dw = w }) }
func InfoWriter(w EventWriter) Option  { return optionFn(func(l *Log) { l.iw = w }) }
func ErrorWriter(w EventWriter) Option { return optionFn(func(l *Log) { l.ew = w }) }
func UserClock(c Clock) Option         { return optionFn(func(l *Log) { l.clock = c }) }
