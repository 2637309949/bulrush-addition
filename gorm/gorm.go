// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gorm

import (
	"errors"
	"fmt"

	"github.com/2637309949/bulrush"
	addition "github.com/2637309949/bulrush-addition"
	"github.com/gin-gonic/gin"
	jzgorm "github.com/jinzhu/gorm"
	"github.com/thoas/go-funk"
)

// GORM Type Defined
type GORM struct {
	bulrush.PNBase
	m        []map[string]interface{}
	cfg      *Config
	DB       *jzgorm.DB
	API      *api
	AutoHook bulrush.PNBase
}

// Plugin defined plugin for bulrush
func (gorm *GORM) Plugin() bulrush.PNRet {
	return func(r *gin.RouterGroup) {
		funk.ForEach(gorm.m, func(item map[string]interface{}) {
			if autoHook, exists := item["autoHook"]; exists == false || autoHook == true {
				collection := item["name"].(string)
				gorm.API.ALL(r, collection)
			}
		})
	}
}

// Config defined GORM Config
type Config struct {
	DBType string `json:"dbType" yaml:"dbType"`
	URL    string `json:"url" yaml:"url"`
}

// Register model
// should provide name and reflector paramters
func (gorm *GORM) Register(manifest map[string]interface{}) {
	if _, ok := manifest["name"]; !ok {
		panic(errors.New("name params must be provided"))
	}
	if _, ok := manifest["reflector"]; !ok {
		panic(errors.New("reflector params must be provided"))
	}
	gorm.m = append(gorm.m, manifest)
}

// Vars return array of Var
func (gorm *GORM) Vars(name string) interface{} {
	m := funk.Find(gorm.m, func(item map[string]interface{}) bool {
		flag := item["name"].(string) == name
		return flag
	}).(map[string]interface{})
	if m != nil {
		return addition.CreateSlice(addition.LeftOkV(m["reflector"]))
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// Var return  Var
// reflect from reflector entity
func (gorm *GORM) Var(name string) interface{} {
	m := funk.Find(gorm.m, func(item map[string]interface{}) bool {
		return item["name"].(string) == name
	}).(map[string]interface{})
	if m != nil {
		return addition.CreateObject(addition.LeftOkV(m["reflector"]))
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// obtain mongo connect session
func createSession(cfg *Config) *jzgorm.DB {
	db, err := jzgorm.Open(cfg.DBType, cfg.URL)
	if err != nil {
		panic(err)
	}
	return db
}

// New New mongo instance
// Export Session, API and AutoHook
func New(bulCfg *bulrush.Config) *GORM {
	cf, err := bulCfg.Unmarshal("sql", Config{})
	if err != nil {
		panic(err)
	}
	conf := cf.(Config)
	db := createSession(&conf)
	gorm := &GORM{
		m:   make([]map[string]interface{}, 0),
		cfg: &conf,
		API: &api{},
		DB:  db,
	}
	gorm.API.gorm = gorm
	return gorm
}
