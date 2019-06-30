// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package logger

import (
	"fmt"
	"io"
)

// LevelWriter defined level writer
type LevelWriter struct {
	W     io.Writer
	Level LOGLEVEL
}

func (c *LevelWriter) Write(p []byte) (int, error) {
	level := toLevelString(c.Level)
	colorLevel := toColorString(c.Level, level)
	pbyte := []byte(fmt.Sprintf("%s %s \n", colorLevel, string(p)))
	return c.W.Write(pbyte)
}

// MutiWriter write with level
type MutiWriter struct {
	writers []io.Writer
	Level   LOGLEVEL
}

func (t *MutiWriter) Write(p []byte) (n int, err error) {
	for _, w := range t.writers {
		if r, ok := w.(*LevelWriter); ok {
			r.Level = t.Level
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

func multiLevelWriter(writers ...io.Writer) *MutiWriter {
	allWriters := make([]io.Writer, 0, len(writers))
	for _, w := range writers {
		if mw, ok := w.(*MutiWriter); ok {
			allWriters = append(allWriters, mw.writers...)
		} else {
			allWriters = append(allWriters, w)
		}
	}
	return &MutiWriter{writers: allWriters}
}
