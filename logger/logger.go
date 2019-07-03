// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package logger

import (
	"fmt"
	"os"

	"github.com/thoas/go-funk"
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

func (j *Journal) fprintf(w *MutiWriter, format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
}

func (j *Journal) appendTransport(transports ...*Transport) *Journal {
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
