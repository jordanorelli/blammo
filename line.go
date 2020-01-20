package blammo

import (
	"bytes"
	"io"
	"strconv"
	"strings"
	"sync"
	"time"
)

type LineWriter struct {
	pool sync.Pool
	out  struct {
		sync.Mutex
		io.Writer
	}
}

func NewLineWriter(w io.Writer) *LineWriter {
	lw := new(LineWriter)
	lw.pool = sync.Pool{
		New: func() interface{} { return new(bytes.Buffer) },
	}
	lw.out.Writer = w
	return lw
}

func (l *LineWriter) WriteEvent(e *Event) {
	buf := l.pool.Get().(*bytes.Buffer)
	buf.Reset()

	buf.WriteString(e.Time.Format(time.RFC3339))
	buf.WriteRune(' ')

	switch e.Level {
	case Debug:
		buf.WriteString("d ")
	case Info:
		buf.WriteString("i ")
	case Error:
		buf.WriteString("e ")
	default:
		buf.WriteString("? ")
	}

	buf.WriteRune('[')
	buf.WriteString(e.Path.String())
	buf.WriteRune(']')
	buf.WriteRune(' ')

	if e.Tags == nil {
		buf.WriteString("[] ")
	} else {
		buf.WriteRune('[')
		writeTags(buf, e.Tags)
		buf.WriteString("] ")
	}

	buf.WriteString(strings.ReplaceAll(e.Text, string('\n'), "\n"))

	l.out.Lock()
	l.out.Write(buf.Bytes())
	l.out.Unlock()
	l.pool.Put(buf)
}

func writeTags(buf *bytes.Buffer, tags *Tags) {
	if tags.parent != nil {
		writeTags(buf, tags.parent)
		buf.WriteRune('+')
	}
	buf.WriteString(tags.key)
	if tags.value != nil {
		buf.WriteRune('=')
		switch v := tags.value.(type) {
		case int:
			buf.WriteString(strconv.Itoa(v))
		case string:
			buf.WriteString(MakeSafeName(v))
		case float64:
			buf.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
		default:
			buf.WriteString("???")
		}
	}
}
