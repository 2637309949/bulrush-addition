// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package color

import (
	"fmt"
	"io"
)

type (
	// LOGLEVEL level type
	LOGLEVEL int
	// Writer for color console
	Writer struct {
		W     io.Writer
		Level LOGLEVEL
	}
)

const (
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

func getLevelTag(level LOGLEVEL) string {
	switch level {
	case ERRORLevel:
		return "ERROR"
	case WARNLevel:
		return "WARN"
	case INFOLevel:
		return "INFO"
	case VERBOSELevel:
		return "VERBOSE"
	case DEBUGLevel:
		return "DEBUG"
	case SILLYLevel:
		return "SILLY"
	case HTTPLevel:
		return "HTTP"
	}
	return string(level)
}

func (c *Writer) getColorString(text string) string {
	switch c.Level {
	case ERRORLevel:
		return Red(text)
	case WARNLevel:
		return Yellow(text)
	case INFOLevel:
		return White(text)
	case VERBOSELevel:
		return Yellow(text)
	case DEBUGLevel:
		return Blue(text)
	case SILLYLevel:
		return Cyan(text)
	case HTTPLevel:
		return Green(text)
	default:
		return text
	}
}

func (c *Writer) Write(p []byte) (int, error) {
	levelTag := getLevelTag(c.Level)
	levelTagColor := c.getColorString(levelTag)
	pbyte := []byte(fmt.Sprintf("%s %s \n", levelTagColor, string(p)))
	return c.W.Write(pbyte)
}
