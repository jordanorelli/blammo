package blammo

import (
	"bytes"
	"testing"
	"time"
)

func TestLineWriter(t *testing.T) {
	refTime := time.Date(2020, time.January, 13, 12, 26, 47, 999999, time.UTC)

	tests := []struct {
		name  string
		event Event
		line  string
	}{
		{
			name:  "empty event",
			event: Event{},
			line:  `0001-01-01T00:00:00Z d [] []`,
		},
		{
			name: "just a time",
			event: Event{
				Time: refTime,
			},
			line: `2020-01-13T12:26:47Z d [] []`,
		},
		{
			name: "root path only",
			event: Event{
				Time: refTime,
				Path: NewPath("root"),
			},
			line: `2020-01-13T12:26:47Z d [root] []`,
		},
		{
			name: "child path",
			event: Event{
				Time: refTime,
				Path: NewPath("root").Child("kid"),
			},
			line: `2020-01-13T12:26:47Z d [root/kid] []`,
		},
		{
			name: "another child path",
			event: Event{
				Time: refTime,
				Path: NewPath("root").Child("kid").Child("fart"),
			},
			line: `2020-01-13T12:26:47Z d [root/kid/fart] []`,
		},
		{
			name: "a message",
			event: Event{
				Time: refTime,
				Path: NewPath("root").Child("kid").Child("fart"),
				Text: "this is a message",
			},
			line: `2020-01-13T12:26:47Z d [root/kid/fart] [] this is a message`,
		},
		{
			name: "a message with an empty tag",
			event: Event{
				Time: refTime,
				Path: NewPath("root").Child("kid").Child("fart"),
				Text: "this is a message",
				Tags: &Tags{key: "alert"},
			},
			line: `2020-01-13T12:26:47Z d [root/kid/fart] [alert] this is a message`,
		},
		{
			name: "a message with two empty tags",
			event: Event{
				Time: refTime,
				Path: NewPath("root").Child("kid").Child("fart"),
				Text: "this is a message",
				Tags: &Tags{
					key:    "zombo-dot-com",
					parent: &Tags{key: "alert"},
				},
			},
			line: `2020-01-13T12:26:47Z d [root/kid/fart] [alert+zombo-dot-com] this is a message`,
		},
		{
			name: "a message with an int tag",
			event: Event{
				Time: refTime,
				Path: NewPath("root").Child("kid").Child("fart"),
				Text: "this is a message",
				Tags: &Tags{
					key:   "num-users",
					value: 15,
				},
			},
			line: `2020-01-13T12:26:47Z d [root/kid/fart] [num-users=15] this is a message`,
		},
		{
			name: "a message with a variety of tags",
			event: Event{
				Time: refTime,
				Path: NewPath("root").Child("kid").Child("fart"),
				Text: "this is a message",
				Tags: &Tags{
					key:   "num-users",
					value: 15,
					parent: &Tags{
						key:   "pi",
						value: 3.14,
						parent: &Tags{
							key:   "request-id",
							value: "b49d31c7-d3bb-4bd3-96fe-34e7c7d2b0a4",
						},
					},
				},
			},
			line: `2020-01-13T12:26:47Z d [root/kid/fart] [request-id=b49d31c7-d3bb-4bd3-96fe-34e7c7d2b0a4+pi=3.14+num-users=15] this is a message`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buf bytes.Buffer
			w := NewLineWriter(&buf)
			w.WriteEvent(&test.event)
			line := buf.String()
			if line != test.line {
				t.Log("expected line does not match observed line")
				t.Logf("expected line: '%s'", test.line)
				t.Logf("observed line: '%s'", line)
				t.Fail()
			}
		})
	}
}
