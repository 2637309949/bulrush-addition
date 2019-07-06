// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apidoc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/thoas/go-funk"

	"github.com/2637309949/bulrush-addition/apidoc/template"
	gormext "github.com/2637309949/bulrush-addition/gorm"
	mgoext "github.com/2637309949/bulrush-addition/mgo"
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

func (api *APIDoc) writeSysDoc(docs []interface{}) *APIDoc {
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
func (api *APIDoc) Plugin(httpProxy *gin.Engine, ext1 *gormext.GORM, ext2 *mgoext.Mongo) *APIDoc {
	template.Handler.Prefix = api.Prefix
	docs := []interface{}{}
	funk.ForEach(*ext1.Docs(), func(doc gormext.Doc) {
		docs = append(docs, doc)
	})
	funk.ForEach(*ext2.Docs(), func(doc mgoext.Doc) {
		docs = append(docs, doc)
	})
	api.writeSysDoc(docs)
	httpProxy.GET(api.Prefix+"/*any", func(c *gin.Context) {
		template.Handler.ServeHTTP(c.Writer, c.Request)
	})
	return api
}
