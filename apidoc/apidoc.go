// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apidoc

import (
	"github.com/2637309949/bulrush-addition/apidoc/template"

	"github.com/gin-gonic/gin"
)

type (
	// APIDoc for apidoc
	APIDoc struct {
		URLPrefix string
	}
)

// New defined return Apidoc struct
func New() *APIDoc {
	return &APIDoc{
		URLPrefix: "/docs",
	}
}

// Init defined struct Init
func (api *APIDoc) Init(init func(*APIDoc)) *APIDoc {
	init(api)
	return api
}

func splitByWidth(str string, size int) []string {
	strLength := len(str)
	var splited []string
	var stop int
	for i := 0; i < strLength; i += size {
		stop = i + size
		if stop > strLength {
			stop = strLength
		}
		splited = append(splited, str[i:stop])
	}
	return splited
}

// Plugin for doc
func (api *APIDoc) Plugin(httpProxy *gin.Engine) *APIDoc {
	template.Handler.Prefix = api.URLPrefix
	httpProxy.GET(api.URLPrefix+"/*any", func(c *gin.Context) {
		template.Handler.ServeHTTP(c.Writer, c.Request)
	})
	return api
}
