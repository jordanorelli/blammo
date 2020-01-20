package blammo

import (
	"bytes"
	"testing"
	"time"
)

func TestLineWriter(t *testing.T) {
	var buf bytes.Buffer
	w := NewLineWriter(&buf)
	w.WriteEvent(&Event{
		Level: Debug,
		Time:  time.Now(),
		Path:  NewPath("alice").Child("bob").Child("carol"),
		Text:  "hey you farthead",
		Tags: &Tags{
			key: "poop",
			parent: &Tags{
				key:   "num_poops",
				value: 27,
			},
		},
	})
	t.Error(buf.String())
}
