/**
 * @author [Double]
 * @email [2637309949@qq.com.com]
 * @create date 2019-01-12 22:46:31
 * @modify date 2019-01-12 22:46:31
 * @desc [Writer]
 */

package color

import (
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
	ERRORLevel LOGLEVEL = iota + 2
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

func (c *Writer) getColorString(message string) string {
	switch c.Level {
	case ERRORLevel:
		return RedBold(message)
	case WARNLevel:
		return YellowBold(message)
	case INFOLevel:
		return WhiteBold(message)
	case VERBOSELevel:
		return YellowBold(message)
	case DEBUGLevel:
		return BlueBold(message)
	case SILLYLevel:
		return Cyan(message)
	case HTTPLevel:
		return GreenBold(message)
	default:
		return message
	}
}

func (c *Writer) Write(p []byte) (int, error) {
	pstring := c.getColorString(string(p))
	pbyte := []byte(pstring)
	return c.W.Write(pbyte)
}
