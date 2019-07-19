// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gormext

import (
	"errors"
	"fmt"

	addition "github.com/2637309949/bulrush-addition"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/thoas/go-funk"
)

type (
	// GORM Type Defined
	GORM struct {
		m    []*Profile
		conf *Config
		DB   *gorm.DB
		API  *API
	}
	// Config defined GORM Config
	Config struct {
		AutoMigrate bool   `json:"automigrate" yaml:"automigrate"`
		DBType      string `json:"dbType" yaml:"dbType"`
		URL         string `json:"url" yaml:"url"`
	}
	// Profile defined model profile
	Profile struct {
		DB        string
		Name      string
		Reflector interface{}
		BanHook   bool
		Opts      *Opts
		docs      *[]Doc
	}
)

// Docs defined Docs for bulrush
func (e *GORM) Docs() *[]Doc {
	docs := []Doc{}
	funk.ForEach(e.m, func(item *Profile) {
		docs = append(docs, *item.docs...)
	})
	return &docs
}

// Plugin defined plugin for bulrush
func (e *GORM) Plugin(r *gin.RouterGroup) *GORM {
	funk.ForEach(e.m, func(item *Profile) {
		if !item.BanHook {
			e.API.ALL(r, item.Name)
		}
	})
	return e
}

// Init e
func (e *GORM) Init(init func(*GORM)) *GORM {
	init(e)
	return e
}

// Register model
// should provide name and reflector paramters
func (e *GORM) Register(profile *Profile) *GORM {
	if profile.Name == "" {
		panic(errors.New("name params must be provided"))
	}
	if profile.Reflector == nil {
		panic(errors.New("reflector params must be provided"))
	}
	profile.docs = &[]Doc{}
	e.m = append(e.m, profile)
	if e.conf.AutoMigrate {
		if err := e.DB.AutoMigrate(profile.Reflector).Error; err != nil {
			addition.RushLogger.Error(fmt.Sprintf("Error in AutoMigrate:%v", err.Error()))
		}
	}
	return e
}

// Profile model profile
func (e *GORM) Profile(name string) *Profile {
	if m := funk.Find(e.m, func(profile *Profile) bool {
		return profile.Name == name
	}); m != nil {
		return m.(*Profile)
	}
	return nil
}

// Vars return array of Var
func (e *GORM) Vars(name string) interface{} {
	m := e.Profile(name)
	if m != nil {
		return addition.CreateSlice(m.Reflector)
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// Var return  Var
// reflect from reflector entity
func (e *GORM) Var(name string) interface{} {
	m := e.Profile(name)
	if m != nil {
		return addition.CreateObject(m.Reflector)
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// Conf set e conf
func (e *GORM) Conf(conf *Config) *GORM {
	db, err := gorm.Open(conf.DBType, conf.URL)
	if err != nil {
		panic(err)
	}
	e.conf = conf
	e.DB = db
	return e
}

// New New mongo instance
// Export Session, API and AutoHook
func New() *GORM {
	e := &GORM{}
	e.m = make([]*Profile, 0)
	e.API = &API{gorm: e, Opts: &Opts{}}
	return e
}
