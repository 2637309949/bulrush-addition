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
		m        []*Profile
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
	// Profile defined model profile
	Profile struct {
		DB        string
		Name      string
		Reflector interface{}
		BanHook   bool
		Opts      *Opts
	}
)

// Plugin defined plugin for bulrush
func (gorm *GORM) Plugin() bulrush.PNRet {
	return func(r *gin.RouterGroup) {
		funk.ForEach(gorm.m, func(item *Profile) {
			if !item.BanHook {
				gorm.API.ALL(r, item.Name)
			}
		})
	}
}

// Register model
// should provide name and reflector paramters
func (gorm *GORM) Register(profile *Profile) *GORM {
	if profile.Name == "" {
		panic(errors.New("name params must be provided"))
	}
	if profile.Reflector == nil {
		panic(errors.New("reflector params must be provided"))
	}
	gorm.m = append(gorm.m, profile)
	return gorm
}

// Profile model profile
func (gorm *GORM) Profile(name string) *Profile {
	if m := funk.Find(gorm.m, func(profile *Profile) bool {
		return profile.Name == name
	}); m != nil {
		return m.(*Profile)
	}
	return nil
}

// Vars return array of Var
func (gorm *GORM) Vars(name string) interface{} {
	m := gorm.Profile(name)
	if m != nil {
		return addition.CreateSlice(addition.LeftOkV(m.Reflector))
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// Var return  Var
// reflect from reflector entity
func (gorm *GORM) Var(name string) interface{} {
	m := gorm.Profile(name)
	if m != nil {
		return addition.CreateObject(addition.LeftOkV(m.Reflector))
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
	gorm.m = make([]*Profile, 0)
	gorm.cfg = conf
	gorm.DB = db
	gorm.API = &API{gorm: gorm, Opts: &Opts{}}
	return gorm
}
