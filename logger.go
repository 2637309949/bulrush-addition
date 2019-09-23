// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package addition

import (
	"github.com/2637309949/bulrush-addition/logger"
)

// RushLogger for app logger
var RushLogger = logger.CreateLogger(logger.SILLY, nil, []*logger.Transport{
	logger.NewPrintTransport(logger.Transport{
		Level:   logger.SILLY,
		Maxsize: logger.Maxsize,
	}),
})
