// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"path"

	"github.com/2637309949/bulrush"
	xormext "github.com/2637309949/bulrush-addition/xorm"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

// Conf defined app conf
var Conf = bulrush.LoadConfig(path.Join("local.yaml"))

// XORMExt defined xorm object
var XORMExt = xormext.
	New().
	Init(func(ext *xormext.XORM) {
		fmt.Println(123)
		cfg := &xormext.Config{}
		if err := Conf.Unmarshal("sql", cfg); err != nil {
			panic(err)
		}
		ext.Conf(cfg)
		logger := xorm.NewSimpleLogger(os.Stdout)
		logger.ShowSQL(true)
		ext.DB.SetLogger(logger)
		ext.Register(&xormext.Profile{
			Name:      "User",
			Reflector: &User{},
			AutoHook:  true,
			Opts: &xormext.Opts{
				RouteHooks: &xormext.RouteHooks{
					List: &xormext.ListHook{
						Cond: func(cond map[string]interface{}, c *gin.Context, info struct{ Name string }) map[string]interface{} {
							return cond
						},
					},
				},
			},
		})
	})
