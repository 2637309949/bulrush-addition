// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package addition

import (
	"os"
	"path"

	"github.com/2637309949/bulrush-addition/logger"
)

// BLOGPATH logger path if you want
var BLOGPATH = os.Getenv("BLOG_PATH")

// RushLogger for bulrush framework use
var RushLogger *logger.Journal

func initLogger() {
	var transports []*logger.Transport
	if BLOGPATH != "" {
		transports = []*logger.Transport{
			&logger.Transport{
				Dirname: path.Join(path.Join(".", BLOGPATH), "error"),
				Level:   logger.ERRORLevel,
				Maxsize: logger.Maxsize,
			},
			&logger.Transport{
				Dirname: path.Join(path.Join(".", BLOGPATH), "combined"),
				Level:   logger.INFOLevel,
				Maxsize: logger.Maxsize,
			},
			&logger.Transport{
				Level: logger.INFOLevel,
			},
		}
	} else {
		transports = []*logger.Transport{
			&logger.Transport{
				Level: logger.SILLYLevel,
			},
		}
	}
	RushLogger = logger.CreateLogger(logger.SILLYLevel, nil, transports)
}

// initLogger
func init() {
	initLogger()
}
