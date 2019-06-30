// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package logger

import (
	"fmt"
)

// Bold returns a Bold string
func Bold(message string) string {
	return fmt.Sprintf("\x1b[1m%s\x1b[21m", message)
}

// Black returns a black string
func Black(message string) string {
	return fmt.Sprintf("\x1b[30m%s\x1b[0m", message)
}

// White returns a white string
func White(message string) string {
	return fmt.Sprintf("\x1b[37m%s\x1b[0m", message)
}

// Cyan returns a cyan string
func Cyan(message string) string {
	return fmt.Sprintf("\x1b[36m%s\x1b[0m", message)
}

// Blue returns a blue string
func Blue(message string) string {
	return fmt.Sprintf("\x1b[34m%s\x1b[0m", message)
}

// Red returns a red string
func Red(message string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", message)
}

// Green returns a green string
func Green(message string) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m", message)
}

// Yellow returns a yellow string
func Yellow(message string) string {
	return fmt.Sprintf("\x1b[33m%s\x1b[0m", message)
}

// Gray returns a gray string
func Gray(message string) string {
	return fmt.Sprintf("\x1b[37m%s\x1b[0m", message)
}

// Magenta returns a magenta string
func Magenta(message string) string {
	return fmt.Sprintf("\x1b[35m%s\x1b[0m", message)
}

// BlackBold returns a black Bold string
func BlackBold(message string) string {
	return fmt.Sprintf("\x1b[30m%s\x1b[0m", Bold(message))
}

// WhiteBold returns a white Bold string
func WhiteBold(message string) string {
	return fmt.Sprintf("\x1b[37m%s\x1b[0m", Bold(message))
}

// CyanBold returns a cyan Bold string
func CyanBold(message string) string {
	return fmt.Sprintf("\x1b[36m%s\x1b[0m", Bold(message))
}

// BlueBold returns a blue Bold string
func BlueBold(message string) string {
	return fmt.Sprintf("\x1b[34m%s\x1b[0m", Bold(message))
}

// RedBold returns a red Bold string
func RedBold(message string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", Bold(message))
}

// GreenBold returns a green Bold string
func GreenBold(message string) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m", Bold(message))
}

// YellowBold returns a yellow Bold string
func YellowBold(message string) string {
	return fmt.Sprintf("\x1b[33m%s\x1b[0m", Bold(message))
}

// GrayBold returns a gray Bold string
func GrayBold(message string) string {
	return fmt.Sprintf("\x1b[37m%s\x1b[0m", Bold(message))
}

// MagentaBold returns a magenta Bold string
func MagentaBold(message string) string {
	return fmt.Sprintf("\x1b[35m%s\x1b[0m", Bold(message))
}

func toLevelString(level LOGLEVEL) string {
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

func toColorString(level LOGLEVEL, text string) string {
	switch level {
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
