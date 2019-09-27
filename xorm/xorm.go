// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package xormext

import (
	"errors"
	"fmt"

	addition "github.com/2637309949/bulrush-addition"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/jinzhu/inflection"
	"github.com/thoas/go-funk"
)

type (
	// XORM Type Defined
	XORM struct {
		m   []*Profile
		c   *Config
		DB  *xorm.EngineGroup
		API *API
	}
	// Config defined XORM Config
	Config struct {
		AutoMigrate bool     `json:"automigrate" yaml:"automigrate"`
		DBType      string   `json:"dbType" yaml:"dbType"`
		URL         []string `json:"url" yaml:"url"`
	}
	// Profile defined model profile
	Profile struct {
		DB        string
		Name      string
		Reflector interface{}
		AutoHook  bool
		Opts      *Opts
		docs      *[]Doc
	}
)

// ComplexMapper defined mapper
type ComplexMapper struct {
}

// Obj2Table defined mapper
func (m ComplexMapper) Obj2Table(o string) string {
	return inflection.Plural(o)
}

// Table2Obj defined mapper
func (m ComplexMapper) Table2Obj(t string) string {
	return inflection.Singular(t)
}

// Docs defined Docs for bulrush
func (e *XORM) Docs() *[]Doc {
	docs := []Doc{}
	funk.ForEach(e.m, func(item *Profile) {
		docs = append(docs, *item.docs...)
	})
	return &docs
}

// Plugin defined plugin for bulrush
func (e *XORM) Plugin(r *gin.RouterGroup) *XORM {
	funk.ForEach(e.m, func(item *Profile) {
		if item.AutoHook {
			e.API.ALL(r, item.Name)
		}
	})
	return e
}

// Init e
func (e *XORM) Init(init func(*XORM)) *XORM {
	init(e)
	return e
}

// Register model
// should provide name and reflector paramters
func (e *XORM) Register(profile *Profile) *XORM {
	if profile.Name == "" {
		panic(errors.New("name params must be provided"))
	}
	if profile.Reflector == nil {
		panic(errors.New("reflector params must be provided"))
	}
	profile.docs = &[]Doc{}
	e.m = append(e.m, profile)
	if e.c.AutoMigrate {
		if err := e.DB.Sync2(profile.Reflector); err != nil {
			addition.RushLogger.Error(fmt.Sprintf("Error in AutoMigrate:%s", err))
		}
	}
	return e
}

// Profile model profile
func (e *XORM) Profile(name string) *Profile {
	if m := funk.Find(e.m, func(profile *Profile) bool {
		return profile.Name == name
	}); m != nil {
		return m.(*Profile)
	}
	panic(fmt.Errorf("profile %s not found", name))
}

// Vars return array of Var
func (e *XORM) Vars(name string) interface{} {
	return addition.CreateSlice(e.Profile(name).Reflector)
}

// Var return  Var
// reflect from reflector entity
func (e *XORM) Var(name string) interface{} {
	return addition.CreateObject(e.Profile(name).Reflector)
}

// Model return instance
func (e *XORM) Model(name string) *xorm.Session {
	m := e.Profile(name)
	if has, _ := e.DB.Table(m.Reflector).Exist(); !has {
		e.DB.Sync2(m.Reflector)
	}
	return e.DB.Table(m.Reflector)
}

// Conf set e conf
func (e *XORM) Conf(conf *Config) *XORM {
	db, err := xorm.NewEngineGroup(conf.DBType, conf.URL)
	if err != nil {
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}
	addition.RushLogger.Info("%v:Connection has been established successfully, URL:%v", conf.DBType, conf.URL)
	e.c = conf
	e.DB = db
	return e
}

// New New mongo instance
// Export Session, API and AutoHook
func New() *XORM {
	e := &XORM{}
	e.m = make([]*Profile, 0)
	e.API = &API{xorm: e, Opts: &Opts{}}
	return e
}
