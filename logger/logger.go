// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package logger

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/thoas/go-funk"
)

// These flags define which text to prefix to each log entry generated by the Logger.
// Bits are or'ed together to control what's printed.
// There is no control over the order they appear (the order listed
// here) or the format they present (as described in the comments).
// The prefix is followed by a colon only when Llongfile or Lshortfile
// is specified.
// For example, flags Ldate | Ltime (or LstdFlags) produce,
//	2009/01/23 01:23:23 message
// while flags Ldate | Ltime | Lmicroseconds | Llongfile produce,
//	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

type (
	// LEVEL level tags
	LEVEL int
	// FormatFunc log format
	FormatFunc func(map[string]string) string
	// Transport for Journal
	Transport struct {
		Type    string
		Dirname string
		Level   LEVEL
		Maxsize int64
	}
	// Journal logger
	Journal struct {
		mu      sync.Mutex
		flag    int
		prefix  string
		buf     []byte
		level   LEVEL
		format  FormatFunc
		writers []struct {
			l LEVEL
			w *MutiWriter
		}
		transports []*Transport
	}
)

// Maxsize file size
const Maxsize = 1024 * 1024 * 10

const (
	// ERROR level
	ERROR LEVEL = iota + 1
	// WARN level
	WARN
	// INFO level
	INFO
	// VERBOSE level
	VERBOSE
	// DEBUG level
	DEBUG
	// SILLY level
	SILLY
	// HTTP level
	HTTP
)

// Error level
func (j *Journal) Error(format string, a ...interface{}) {
	var r = funk.Find(j.writers, func(x struct {
		l LEVEL
		w *MutiWriter
	}) bool {
		return x.l == ERROR
	}).(struct {
		l LEVEL
		w *MutiWriter
	})
	if r.w != nil {
		r.w.Level = ERROR
		j.fprintf(r.w, format, a...)
	}
}

// Warn level
func (j *Journal) Warn(format string, a ...interface{}) {
	var r = funk.Find(j.writers, func(x struct {
		l LEVEL
		w *MutiWriter
	}) bool {
		return x.l == WARN
	}).(struct {
		l LEVEL
		w *MutiWriter
	})
	if r.w != nil {
		r.w.Level = WARN
		j.fprintf(r.w, format, a...)
	}
}

// Info level
func (j *Journal) Info(format string, a ...interface{}) {
	var r = funk.Find(j.writers, func(x interface{}) bool {
		return x.(struct {
			l LEVEL
			w *MutiWriter
		}).l == INFO
	}).(struct {
		l LEVEL
		w *MutiWriter
	})
	if r.w != nil {
		r.w.Level = INFO
		j.fprintf(r.w, format, a...)
	}
}

// Verbose level
func (j *Journal) Verbose(format string, a ...interface{}) {
	var r = funk.Find(j.writers, func(x struct {
		l LEVEL
		w *MutiWriter
	}) bool {
		return x.l == VERBOSE
	}).(struct {
		l LEVEL
		w *MutiWriter
	})
	if r.w != nil {
		r.w.Level = VERBOSE
		j.fprintf(r.w, format, a...)
	}
}

// Debug level
func (j *Journal) Debug(format string, a ...interface{}) {
	var r = funk.Find(j.writers, func(x struct {
		l LEVEL
		w *MutiWriter
	}) bool {
		return x.l == DEBUG
	}).(struct {
		l LEVEL
		w *MutiWriter
	})
	if r.w != nil {
		r.w.Level = DEBUG
		j.fprintf(r.w, format, a...)
	}
}

// Silly level
func (j *Journal) Silly(format string, a ...interface{}) {
	var r = funk.Find(j.writers, func(x struct {
		l LEVEL
		w *MutiWriter
	}) bool {
		return x.l == SILLY
	}).(struct {
		l LEVEL
		w *MutiWriter
	})
	if r.w != nil {
		r.w.Level = SILLY
		j.fprintf(r.w, format, a...)
	}
}

// HTTP level
func (j *Journal) HTTP(format string, a ...interface{}) {
	var r = funk.Find(j.writers, func(x struct {
		l LEVEL
		w *MutiWriter
	}) bool {
		return x.l == HTTP
	}).(struct {
		l LEVEL
		w *MutiWriter
	})
	if r.w != nil {
		r.w.Level = HTTP
		j.fprintf(r.w, format, a...)
	}
}

// Init Journal
func (j *Journal) Init(init func(*Journal)) *Journal {
	init(j)
	return j
}

// SetFlags sets the output flags for the logger.
func (j *Journal) SetFlags(flag int) {
	j.mu.Lock()
	defer j.mu.Unlock()
	j.flag = flag
}

// formatHeader writes log header to buf in following order:
//   * l.prefix (if it's not blank),
//   * date and/or time (if corresponding flags are provided),
//   * file and line number (if corresponding flags are provided).
func (j *Journal) formatHeader(buf *[]byte, t time.Time, file string, line int) {
	*buf = append(*buf, j.prefix...)
	if j.flag&(Ldate|Ltime|Lmicroseconds) != 0 {
		if j.flag&LUTC != 0 {
			t = t.UTC()
		}
		if j.flag&Ldate != 0 {
			year, month, day := t.Date()
			itoa(buf, year, 4)
			*buf = append(*buf, '/')
			itoa(buf, int(month), 2)
			*buf = append(*buf, '/')
			itoa(buf, day, 2)
			*buf = append(*buf, ' ')
		}
		if j.flag&(Ltime|Lmicroseconds) != 0 {
			hour, min, sec := t.Clock()
			itoa(buf, hour, 2)
			*buf = append(*buf, ':')
			itoa(buf, min, 2)
			*buf = append(*buf, ':')
			itoa(buf, sec, 2)
			if j.flag&Lmicroseconds != 0 {
				*buf = append(*buf, '.')
				itoa(buf, t.Nanosecond()/1e3, 6)
			}
			*buf = append(*buf, ' ')
		}
	}
	if j.flag&(Lshortfile|Llongfile) != 0 {
		if j.flag&Lshortfile != 0 {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
		}
		*buf = append(*buf, file...)
		*buf = append(*buf, ':')
		itoa(buf, line, -1)
		*buf = append(*buf, ": "...)
	}
}

// Output writes the output for a logging event. The string s contains
// the text to print after the prefix specified by the flags of the
// Logger. A newline is appended if the last character of s is not
// already a newline. Calldepth is used to recover the PC and is
// provided for generality, although at the moment on all pre-defined
// paths it will be 2.
func (j *Journal) output(w *MutiWriter, calldepth int, s string) error {
	now := time.Now()
	var file string
	var line int
	j.mu.Lock()
	defer j.mu.Unlock()
	if j.flag&(Lshortfile|Llongfile) != 0 {
		j.mu.Unlock()
		var ok bool
		_, file, line, ok = runtime.Caller(calldepth)
		if !ok {
			file = "???"
			line = 0
		}
		j.mu.Lock()
	}
	j.buf = j.buf[:0]
	j.formatHeader(&j.buf, now, file, line)
	j.buf = append(j.buf, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		j.buf = append(j.buf, '\n')
	}
	_, err := w.Write(j.buf)
	return err
}

func (j *Journal) fprintf(w *MutiWriter, format string, a ...interface{}) {
	j.output(w, 2, fmt.Sprintf(format, a...))
}

// AppendTransports defined append transports
func (j *Journal) AppendTransports(transports ...*Transport) *Journal {
	for _, transport := range transports {
		if transport.Level == 0 {
			transport.Level = INFO
		}
		if transport.Dirname == "" {
			transport.Type = "Print"
		} else if transport.Dirname != "" {
			transport.Type = "File"
		}
	}
	j.transports = append(j.transports, transports...)
	levels := []LEVEL{ERROR, WARN, INFO, VERBOSE, DEBUG, SILLY, HTTP}
	j.writers = make([]struct {
		l LEVEL
		w *MutiWriter
	}, 0)
	for _, level := range levels {
		writer := struct {
			l LEVEL
			w *MutiWriter
		}{
			l: level,
			w: j.createWriter(level),
		}
		j.writers = append(j.writers, writer)
	}
	return j
}

// create writer
func (j *Journal) createWriter(level LEVEL) *MutiWriter {
	var writer *MutiWriter
	if j.level >= level {
		for _, transport := range j.transports {
			if transport.Level >= level {
				if transport.Type == "Print" {
					if writer != nil {
						writer = multiLevelWriter(writer, &LevelWriter{
							W:     os.Stdout,
							Level: transport.Level,
						})
					} else {
						writer = multiLevelWriter(&LevelWriter{
							W:     os.Stdout,
							Level: transport.Level,
						})
					}

				}
				if transport.Type == "File" {
					f, _ := OpenFile(transport.Dirname, transport.Maxsize, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
					if writer != nil {
						writer = multiLevelWriter(writer, f)
					} else {
						writer = multiLevelWriter(f)
					}
				}
			}
		}
	}
	return writer
}

// CreateLogger logger
func CreateLogger(level LEVEL, format FormatFunc, transports []*Transport) *Journal {
	j := &Journal{}
	for _, transport := range transports {
		if transport.Level == 0 {
			transport.Level = INFO
		}
		if transport.Dirname == "" {
			transport.Type = "Print"
		} else if transport.Dirname != "" {
			transport.Type = "File"
		}
	}
	j.format = format
	j.level = level
	j.transports = transports
	levels := []LEVEL{ERROR, WARN, INFO, VERBOSE, DEBUG, SILLY, HTTP}
	for _, level := range levels {
		writer := struct {
			l LEVEL
			w *MutiWriter
		}{
			l: level,
			w: j.createWriter(level),
		}
		j.writers = append(j.writers, writer)
	}
	return j
}
