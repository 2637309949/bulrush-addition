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

type (
	// GORM Type Defined
	GORM struct {
		bulrush.PNBase
		m            []map[string]interface{}
		cfg          *Config
		DB           *jzgorm.DB
		AutoHook     bulrush.PNBase
		API          *api
		APIFeature1  *api
		APIFeature2  *api
		APIFeature3  *api
		APIFeature4  *api
		APIFeature5  *api
		APIFeature6  *api
		APIFeature7  *api
		APIFeature8  *api
		APIFeature9  *api
		APIFeature10 *api
		APIFeature11 *api
		APIFeature12 *api
		APIFeature13 *api
		APIFeature14 *api
		APIFeature15 *api
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
	gorm.API = &api{gorm: gorm}
	gorm.APIFeature1 = &api{gorm: gorm}
	gorm.APIFeature2 = &api{gorm: gorm}
	gorm.APIFeature3 = &api{gorm: gorm}
	gorm.APIFeature4 = &api{gorm: gorm}
	gorm.APIFeature5 = &api{gorm: gorm}
	gorm.APIFeature6 = &api{gorm: gorm}
	gorm.APIFeature7 = &api{gorm: gorm}
	gorm.APIFeature8 = &api{gorm: gorm}
	gorm.APIFeature9 = &api{gorm: gorm}
	gorm.APIFeature10 = &api{gorm: gorm}
	gorm.APIFeature11 = &api{gorm: gorm}
	gorm.APIFeature12 = &api{gorm: gorm}
	gorm.APIFeature13 = &api{gorm: gorm}
	gorm.APIFeature14 = &api{gorm: gorm}
	gorm.APIFeature15 = &api{gorm: gorm}
	return gorm
}
