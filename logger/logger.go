/**
 * @author [Double]
 * @email [2637309949@qq.com.com]
 * @create date 2019-01-12 22:46:31
 * @modify date 2019-01-12 22:46:31
 * @desc [bulrush LoggerWriter plugin]
 */

package logger

import (
	"fmt"
	"os"

	"github.com/2637309949/bulrush-addition/logger/color"

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
			w *LevelWriter
		}
		transports []*Transport
	}
)

const (
	// Maxsize file size
	Maxsize = 1024 * 1024 * 10
	// ERRORLevel level
	ERRORLevel LOGLEVEL = iota + 1
	// WARNLevel level
	WARNLevel
	// INFOLevel level
	INFOLevel
	// VERBOSELevel level
	VERBOSELevel
	// DEBUGLevel level
	DEBUGLevel
	// SILLYLevel level
	SILLYLevel
	// HTTPLevel level
	HTTPLevel
)

// Error level
func (j *Journal) Error(format string, a ...interface{}) {
	var r = funk.Find(j.writers, func(x struct {
		l LOGLEVEL
		w *LevelWriter
	}) bool {
		return x.l == ERRORLevel
	}).(struct {
		l LOGLEVEL
		w *LevelWriter
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
		w *LevelWriter
	}) bool {
		return x.l == WARNLevel
	}).(struct {
		l LOGLEVEL
		w *LevelWriter
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
			w *LevelWriter
		}).l == INFOLevel
	}).(struct {
		l LOGLEVEL
		w *LevelWriter
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
		w *LevelWriter
	}) bool {
		return x.l == VERBOSELevel
	}).(struct {
		l LOGLEVEL
		w *LevelWriter
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
		w *LevelWriter
	}) bool {
		return x.l == DEBUGLevel
	}).(struct {
		l LOGLEVEL
		w *LevelWriter
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
		w *LevelWriter
	}) bool {
		return x.l == SILLYLevel
	}).(struct {
		l LOGLEVEL
		w *LevelWriter
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
		w *LevelWriter
	}) bool {
		return x.l == HTTPLevel
	}).(struct {
		l LOGLEVEL
		w *LevelWriter
	})
	if r.w != nil {
		r.w.Level = HTTPLevel
		j.fprintf(r.w, format, a...)
	}
}

func (j *Journal) fprintf(w *LevelWriter, format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
	fmt.Fprintln(w, "")
}

// create writer
func (j *Journal) createWriter(level LOGLEVEL) *LevelWriter {
	var writer *LevelWriter
	if j.level >= level {
		for _, transport := range j.transports {
			if transport.Level >= level {
				if transport.Type == "Print" {
					if writer != nil {
						writer = multiLevelWriter(writer, &color.Writer{
							W:     os.Stdout,
							Level: color.LOGLEVEL(transport.Level),
						})
					} else {
						writer = multiLevelWriter(&color.Writer{
							W:     os.Stdout,
							Level: color.LOGLEVEL(transport.Level),
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
			w *LevelWriter
		}{
			l: level,
			w: j.createWriter(level),
		}
		j.writers = append(j.writers, writer)
	}
	return j
}
