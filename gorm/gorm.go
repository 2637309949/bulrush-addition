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
		m   []*Profile
		DB  *gorm.DB
		API *API
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
		docs      *[]Doc
	}
)

// Docs defined Docs for bulrush
func (ext *GORM) Docs() *[]Doc {
	docs := []Doc{}
	funk.ForEach(ext.m, func(item *Profile) {
		docs = append(docs, *item.docs...)
	})
	return &docs
}

// Plugin defined plugin for bulrush
func (ext *GORM) Plugin(r *gin.RouterGroup) *GORM {
	funk.ForEach(ext.m, func(item *Profile) {
		if !item.BanHook {
			ext.API.ALL(r, item.Name)
		}
	})
	return ext
}

// Init ext
func (ext *GORM) Init(init func(*GORM)) *GORM {
	init(ext)
	return ext
}

// Register model
// should provide name and reflector paramters
func (ext *GORM) Register(profile *Profile) *GORM {
	if profile.Name == "" {
		panic(errors.New("name params must be provided"))
	}
	if profile.Reflector == nil {
		panic(errors.New("reflector params must be provided"))
	}
	profile.docs = &[]Doc{}
	ext.m = append(ext.m, profile)
	return ext
}

// Profile model profile
func (ext *GORM) Profile(name string) *Profile {
	if m := funk.Find(ext.m, func(profile *Profile) bool {
		return profile.Name == name
	}); m != nil {
		return m.(*Profile)
	}
	return nil
}

// Vars return array of Var
func (ext *GORM) Vars(name string) interface{} {
	m := ext.Profile(name)
	if m != nil {
		return addition.CreateSlice(m.Reflector)
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// Var return  Var
// reflect from reflector entity
func (ext *GORM) Var(name string) interface{} {
	m := ext.Profile(name)
	if m != nil {
		return addition.CreateObject(m.Reflector)
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// Conf set ext conf
func (ext *GORM) Conf(conf *Config) *GORM {
	db, err := gorm.Open(conf.DBType, conf.URL)
	if err != nil {
		panic(err)
	}
	ext.DB = db
	return ext
}

// New New mongo instance
// Export Session, API and AutoHook
func New() *GORM {
	ext := &GORM{}
	ext.m = make([]*Profile, 0)
	ext.API = &API{gorm: ext, Opts: &Opts{}}
	return ext
}
