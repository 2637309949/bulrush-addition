// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gormext

import (
	"errors"
	"fmt"

	"github.com/2637309949/bulrush"
	addition "github.com/2637309949/bulrush-addition"
	"github.com/gin-gonic/gin"
	jzgorm "github.com/jinzhu/gorm"
	"github.com/thoas/go-funk"
)

type (
	// GORM Type Defined
	GORM struct {
		bulrush.PNBase
		m        []map[string]interface{}
		cfg      *Config
		DB       *jzgorm.DB
		AutoHook bulrush.PNBase
		API      *API
	}
	// Config defined GORM Config
	Config struct {
		DBType string `json:"dbType" yaml:"dbType"`
		URL    string `json:"url" yaml:"url"`
	}
)

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

// Register model
// should provide name and reflector paramters
func (gorm *GORM) Register(manifest map[string]interface{}) *GORM {
	if _, ok := manifest["name"]; !ok {
		panic(errors.New("name params must be provided"))
	}
	if _, ok := manifest["reflector"]; !ok {
		panic(errors.New("reflector params must be provided"))
	}
	gorm.m = append(gorm.m, manifest)
	return gorm
}

// Profile model profile
func (gorm *GORM) Profile(name string) map[string]interface{} {
	m := funk.Find(gorm.m, func(item map[string]interface{}) bool {
		flag := item["name"].(string) == name
		return flag
	})
	if m != nil {
		profile := m.(map[string]interface{})
		return profile
	}
	return nil
}

// Vars return array of Var
func (gorm *GORM) Vars(name string) interface{} {
	m := gorm.Profile(name)
	if m != nil {
		return addition.CreateSlice(addition.LeftOkV(m["reflector"]))
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// Var return  Var
// reflect from reflector entity
func (gorm *GORM) Var(name string) interface{} {
	m := gorm.Profile(name)
	if m != nil {
		return addition.CreateObject(addition.LeftOkV(m["reflector"]))
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// openDB get gorm connect session
func openSession(cfg *Config) *jzgorm.DB {
	db, err := jzgorm.Open(cfg.DBType, cfg.URL)
	if err != nil {
		panic(err)
	}
	return db
}

// New New mongo instance
// Export Session, API and AutoHook
func New(bulCfg *bulrush.Config) *GORM {
	conf := &Config{}
	if err := bulCfg.Unmarshal("sql", conf); err != nil {
		panic(err)
	}
	db := openSession(conf)
	gorm := &GORM{}
	gorm.m = make([]map[string]interface{}, 0)
	gorm.cfg = conf
	gorm.DB = db
	gorm.API = &API{gorm: gorm, Opts: &APIOpts{}}
	return gorm
}
