// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apidoc

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/2637309949/bulrush-addition/apidoc/template"

	"github.com/gin-gonic/gin"
)

type (
	// APIDoc for apidoc
	APIDoc struct {
		Prefix string
	}
)

// New defined return Apidoc struct
func New() *APIDoc {
	return &APIDoc{
		Prefix: "/docs",
	}
}

// Doc defined apidoc
func (api *APIDoc) Doc(dir string) *APIDoc {
	apiData, err := ioutil.ReadFile(path.Join(dir, APIData))
	fmt.Println(path.Join(dir, APIData))
	fmt.Println(path.Join("/", APIData))
	if err == nil {
		template.WriteFile(path.Join("/", APIData), apiData, 0777)
	}
	apiProject, err := ioutil.ReadFile(path.Join(dir, APIProject))
	if err == nil {
		template.WriteFile(path.Join("/", APIProject), apiProject, 0777)
	}
	return api
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
	template.Handler.Prefix = api.Prefix
	httpProxy.GET(api.Prefix+"/*any", func(c *gin.Context) {
		template.Handler.ServeHTTP(c.Writer, c.Request)
	})
	return api
}
