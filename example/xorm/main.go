// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"path"

	"github.com/2637309949/bulrush"
	addition "github.com/2637309949/bulrush-addition"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/go-events"
)

// Logger defined plugin
var Logger = addition.RushLogger

func main() {
	app := bulrush.Default()
	app.PostUse(XORMExt)
	app.Config(path.Join("local.yaml"))
	app.Use(func(event events.EventEmmiter) {
		event.On(bulrush.EventsRunning, func(message ...interface{}) {
			Logger.Info("running %v", message)
		})
	})
	app.Run()
}
