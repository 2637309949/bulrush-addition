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
	// LOGLEVEL level tags
	LOGLEVEL int
	// FormatFunc log format
	FormatFunc func(map[string]string) string
	// Transport for Journal
	Transport struct {
		Type    string
		Dirname string
		Level   LOGLEVEL
		Maxsize int64
	}
	// Journal logger
	Journal struct {
		level   LOGLEVEL
		format  FormatFunc
		writers []struct {
			l LOGLEVEL
			w *MutiWriter
		}
		transports []*Transport
	}
)

const (
	// Maxsize file size
	Maxsize = 1024 * 1024 * 10
	// ERRORLevel level
	ERRORLevel LOGLEVEL = 1
	// WARNLevel level
	WARNLevel LOGLEVEL = 2
	// INFOLevel level
	INFOLevel LOGLEVEL = 3
	// VERBOSELevel level
	VERBOSELevel LOGLEVEL = 4
	// DEBUGLevel level
	DEBUGLevel LOGLEVEL = 5
	// SILLYLevel level
	SILLYLevel LOGLEVEL = 6
	// HTTPLevel level
	HTTPLevel LOGLEVEL = 7
)

// Error level
func (j *Journal) Error(format string, a ...interface{}) {
	var r = funk.Find(j.writers, func(x struct {
		l LOGLEVEL
		w *MutiWriter
	}) bool {
		return x.l == ERRORLevel
	}).(struct {
		l LOGLEVEL
		w *MutiWriter
	})
	if r.w != nil {
		r.w.Level = ERRORLevel
		j.fprintf(r.w, format, a...)
	}
}

// Warn level
func (j *Journal) Warn(format string, a ...interface{}) {
	var r = funk.Find(j.writers, func(x struct {
		l LOGLEVEL
		w *MutiWriter
	}) bool {
		return x.l == WARNLevel
	}).(struct {
		l LOGLEVEL
		w *MutiWriter
	})
	if r.w != nil {
		r.w.Level = WARNLevel
		j.fprintf(r.w, format, a...)
	}
}

// Info level
func (j *Journal) Info(format string, a ...interface{}) {
	var r = funk.Find(j.writers, func(x interface{}) bool {
		return x.(struct {
			l LOGLEVEL
			w *MutiWriter
		}).l == INFOLevel
	}).(struct {
		l LOGLEVEL
		w *MutiWriter
	})
	if r.w != nil {
		r.w.Level = INFOLevel
		j.fprintf(r.w, format, a...)
	}
}

// Verbose level
func (j *Journal) Verbose(format string, a ...interface{}) {
	var r = funk.Find(j.writers, func(x struct {
		l LOGLEVEL
		w *MutiWriter
	}) bool {
		return x.l == VERBOSELevel
	}).(struct {
		l LOGLEVEL
		w *MutiWriter
	})
	if r.w != nil {
		r.w.Level = VERBOSELevel
		j.fprintf(r.w, format, a...)
	}
}

// Debug level
func (j *Journal) Debug(format string, a ...interface{}) {
	var r = funk.Find(j.writers, func(x struct {
		l LOGLEVEL
		w *MutiWriter
	}) bool {
		return x.l == DEBUGLevel
	}).(struct {
		l LOGLEVEL
		w *MutiWriter
	})
	if r.w != nil {
		r.w.Level = DEBUGLevel
		j.fprintf(r.w, format, a...)
	}
}

// Silly level
func (j *Journal) Silly(format string, a ...interface{}) {
	var r = funk.Find(j.writers, func(x struct {
		l LOGLEVEL
		w *MutiWriter
	}) bool {
		return x.l == SILLYLevel
	}).(struct {
		l LOGLEVEL
		w *MutiWriter
	})
	if r.w != nil {
		r.w.Level = SILLYLevel
		j.fprintf(r.w, format, a...)
	}
}

// HTTP level
func (j *Journal) HTTP(format string, a ...interface{}) {
	var r = funk.Find(j.writers, func(x struct {
		l LOGLEVEL
		w *MutiWriter
	}) bool {
		return x.l == HTTPLevel
	}).(struct {
		l LOGLEVEL
		w *MutiWriter
	})
	if r.w != nil {
		r.w.Level = HTTPLevel
		j.fprintf(r.w, format, a...)
	}
}

func (j *Journal) fprintf(w *MutiWriter, format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
}

// create writer
func (j *Journal) createWriter(level LOGLEVEL) *MutiWriter {
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
func CreateLogger(level LOGLEVEL, format FormatFunc, transports []*Transport) *Journal {
	j := &Journal{}
	for _, transport := range transports {
		if transport.Level == 0 {
			transport.Level = INFOLevel
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
	levels := []LOGLEVEL{ERRORLevel, WARNLevel, INFOLevel, VERBOSELevel, DEBUGLevel, SILLYLevel, HTTPLevel}
	for _, level := range levels {
		writer := struct {
			l LOGLEVEL
			w *MutiWriter
		}{
			l: level,
			w: j.createWriter(level),
		}
		j.writers = append(j.writers, writer)
	}
	return j
}
