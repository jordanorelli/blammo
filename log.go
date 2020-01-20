package blammo

import (
	"fmt"
)

type Log struct {
	lvl   Level
	dw    EventWriter
	iw    EventWriter
	ew    EventWriter
	path  *Path
	clock Clock
	tags  *Tags
}

type Options struct {
	Debug EventWriter
	Info  EventWriter
	Error EventWriter
}

var defaults Options

func init() {
	defaults.Debug = NullWriter{}
}

func NewLog() *Log {
	return new(Log)
}

func format(t string, args ...interface{}) string {
	if len(args) == 0 {
		return t
	}
	return fmt.Sprintf(t, args...)
}

func (l *Log) Debug(t string, args ...interface{}) {
}

func (l *Log) Info(t string, args ...interface{}) {
}

func (l *Log) Error(t string, args ...interface{}) {
}

func (l *Log) Child(name string) *Log {
	return l
}

func (l *Log) Tag(key string, value interface{}) *Log {
	return l
}
