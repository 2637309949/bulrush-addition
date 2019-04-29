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
	"io"
	"os"
	"path"

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
			w io.Writer
		}
		transports []*Transport
	}
)

const (
	// Maxsize file size
	Maxsize = 1024 * 1024 * 10
)

const (
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

// Info level
func (j *Journal) Info(format string, a ...interface{}) {
	var r = funk.Find(j.writers, func(x interface{}) bool {
		return x.(struct {
			l LOGLEVEL
			w io.Writer
		}).l == INFOLevel
	}).(struct {
		l LOGLEVEL
		w io.Writer
	})
	if r.w != nil {
		fmt.Fprintf(r.w, format, a...)
		fmt.Fprint(r.w, "\n")
	}
}

// Error level
func (j *Journal) Error(format string, a ...interface{}) {
	var r = funk.Find(j.writers, func(x interface{}) bool {
		return x.(struct {
			l LOGLEVEL
			w io.Writer
		}).l == ERRORLevel
	}).(struct {
		l LOGLEVEL
		w io.Writer
	})
	if r.w != nil {
		fmt.Fprintf(r.w, format, a...)
		fmt.Fprint(r.w, "\n")
	}
}

// create writer
func (j *Journal) createWriter(level LOGLEVEL) io.Writer {
	var writer io.Writer
	if j.level >= level {
		for _, transport := range j.transports {
			if transport.Level >= level {
				if transport.Type == "Print" {
					if writer != nil {
						writer = io.MultiWriter(writer, os.Stdout)
					} else {
						writer = io.MultiWriter(os.Stdout)
					}
				}
				if transport.Type == "File" {
					f, _ := OpenFile(transport.Dirname, transport.Maxsize, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
					if writer != nil {
						writer = io.MultiWriter(writer, f)
					} else {
						writer = io.MultiWriter(f)
					}
				}
			}
		}
	}
	return writer
}

// CreateLogger logger
func (j *Journal) CreateLogger(level LOGLEVEL, format FormatFunc, transports []*Transport) *Journal {
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

	// create writers
	infoWriter := struct {
		l LOGLEVEL
		w io.Writer
	}{
		l: INFOLevel,
		w: j.createWriter(INFOLevel),
	}
	errorWriter := struct {
		l LOGLEVEL
		w io.Writer
	}{
		l: ERRORLevel,
		w: j.createWriter(ERRORLevel),
	}
	j.writers = append(j.writers, infoWriter, errorWriter)
	return j
}

// CreateLogger log to console and file
func CreateLogger(dirPath string) *Journal {
	j := &Journal{}
	j.CreateLogger(
		INFOLevel,
		nil,
		[]*Transport{
			&Transport{
				Dirname: path.Join(dirPath, "error"),
				Level:   ERRORLevel,
				Maxsize: Maxsize,
			},
			&Transport{
				Dirname: path.Join(dirPath, "combined"),
				Level:   INFOLevel,
				Maxsize: Maxsize,
			},
			&Transport{
				Level: INFOLevel,
			},
		},
	)
	return j
}

// CreateHTTPLogger log to console and file
func CreateHTTPLogger(dirPath string) *Journal {
	j := &Journal{}
	j.CreateLogger(
		INFOLevel,
		nil,
		[]*Transport{
			&Transport{
				Dirname: path.Join(dirPath, "http"),
				Level:   HTTPLevel,
				Maxsize: Maxsize,
			},
			&Transport{
				Level: INFOLevel,
			},
		},
	)
	return j
}

// CreateConsoleLogger log to console
func CreateConsoleLogger() *Journal {
	j := &Journal{}
	j.CreateLogger(
		INFOLevel,
		nil,
		[]*Transport{
			&Transport{
				Level: INFOLevel,
			},
		},
	)
	return j
}