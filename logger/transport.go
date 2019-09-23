// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package logger

import "os"

// NewPrintTransport defined PrintTransport
func NewPrintTransport(t Transport) *Transport {
	return &Transport{
		Dirname: t.Dirname,
		Level:   t.Level,
		Maxsize: t.Maxsize,
		Writer:  NewLevelWriter(t.Level),
	}
}

// NewFileTransport defined FileTransport
func NewFileTransport(t Transport) *Transport {
	return &Transport{
		Dirname: t.Dirname,
		Level:   t.Level,
		Maxsize: t.Maxsize,
		Writer:  NewRotateWriter(t.Dirname, t.Maxsize, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600),
	}
}
