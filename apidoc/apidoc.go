// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apidoc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/2637309949/bulrush-addition/apidoc/template"
	gormext "github.com/2637309949/bulrush-addition/gorm"
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
	if err == nil {
		template.WriteFile(path.Join("/", APIData), apiData, 0777)
	}
	apiProject, err := ioutil.ReadFile(path.Join(dir, APIProject))
	if err == nil {
		template.WriteFile(path.Join("/", APIProject), apiProject, 0777)
	}
	return api
}

func (api *APIDoc) writeSysDoc(docs *[]gormext.Doc) *APIDoc {
	apiData := struct {
		API interface{} `json:"api" yaml:"api"`
	}{
		API: docs,
	}
	apiDataByte, err := json.Marshal(apiData)
	apiDataString := fmt.Sprintf("define(%s)", string(apiDataByte))
	template.WriteFile(path.Join("/", APIDataSys), []byte(apiDataString), 0777)
	if err != nil {
		panic(err)
	}
	return api
}

// Init defined struct Init
func (api *APIDoc) Init(init func(*APIDoc)) *APIDoc {
	init(api)
	return api
}

// Plugin for doc
func (api *APIDoc) Plugin(httpProxy *gin.Engine, gormext *gormext.GORM) *APIDoc {
	template.Handler.Prefix = api.Prefix
	api.writeSysDoc(gormext.Docs())
	httpProxy.GET(api.Prefix+"/*any", func(c *gin.Context) {
		template.Handler.ServeHTTP(c.Writer, c.Request)
	})
	return api
}
