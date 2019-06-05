package logger

import (
	"io"

	"github.com/2637309949/bulrush-addition/logger/color"
)

// LevelWriter write with level
type LevelWriter struct {
	Level   LOGLEVEL
	writers []io.Writer
}

func (t *LevelWriter) Write(p []byte) (n int, err error) {
	for _, w := range t.writers {
		if r, ok := w.(*color.Writer); ok {
			r.Level = color.LOGLEVEL(t.Level)
		}
		n, err = w.Write(p)
		if err != nil {
			return
		}
		if n != len(p) {
			err = io.ErrShortWrite
			return
		}
	}
	return len(p), nil
}

func multiLevelWriter(writers ...io.Writer) *LevelWriter {
	allWriters := make([]io.Writer, 0, len(writers))
	for _, w := range writers {
		if mw, ok := w.(*LevelWriter); ok {
			allWriters = append(allWriters, mw.writers...)
		} else {
			allWriters = append(allWriters, w)
		}
	}
	return &LevelWriter{writers: allWriters}
}
