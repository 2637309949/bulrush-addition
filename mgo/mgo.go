// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mgoext

import (
	"errors"
	"fmt"
	"strings"

	addition "github.com/2637309949/bulrush-addition"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/jinzhu/inflection"
	"github.com/thoas/go-funk"
)

type (
	// Mongo Type Defined
	Mongo struct {
		m       []*Profile
		conf    *mgo.DialInfo
		Session *mgo.Session
		API     *API
	}
	// Profile defined model profile
	Profile struct {
		DB         string
		Collection string
		Name       string
		Reflector  interface{}
		BanHook    bool
		Opts       *Opts
		docs       *[]Doc
	}
)

// Docs defined Docs for bulrush
func (e *Mongo) Docs() *[]Doc {
	docs := []Doc{}
	funk.ForEach(e.m, func(item *Profile) {
		docs = append(docs, *item.docs...)
	})
	return &docs
}

// Plugin defined plugin for bulrush
func (e *Mongo) Plugin(r *gin.RouterGroup) *Mongo {
	funk.ForEach(e.m, func(item *Profile) {
		if !item.BanHook {
			e.API.ALL(r, item.Name)
		}
	})
	return e
}

// Init mgo
func (e *Mongo) Init(init func(*Mongo)) *Mongo {
	init(e)
	return e
}

// Register model
// should provide name and reflector paramters
func (e *Mongo) Register(profile *Profile) *Mongo {
	if profile.Name == "" {
		panic(errors.New("name params must be provided"))
	}
	if profile.Reflector == nil {
		panic(errors.New("reflector params must be provided"))
	}
	profile.docs = &[]Doc{}
	e.m = append(e.m, profile)
	return e
}

// Profile model profile
func (e *Mongo) Profile(name string) *Profile {
	if m := funk.Find(e.m, func(item *Profile) bool {
		return item.Name == name
	}); m != nil {
		return m.(*Profile)
	}
	return nil
}

// Vars return array of Var
func (e *Mongo) Vars(name string) interface{} {
	m := e.Profile(name)
	if m != nil {
		return addition.CreateSlice(m.Reflector)
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// Var return  Var
// reflect from reflector entity
func (e *Mongo) Var(name string) interface{} {
	m := e.Profile(name)
	if m != nil {
		return addition.CreateObject(m.Reflector)
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// Model return instance
// throw error if not exists these model
func (e *Mongo) Model(name string) *mgo.Collection {
	m := e.Profile(name)
	if m != nil {
		db := addition.Some(m.DB, e.conf.Database).(string)
		collect := addition.Some(m.Collection, name).(string)
		collect = strings.ToLower(collect)
		collect = inflection.Plural(collect)
		return e.Session.DB(db).C(collect)
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// Conf defined conf
func (e *Mongo) Conf(conf *mgo.DialInfo) *Mongo {
	session, err := mgo.DialWithInfo(conf)
	if err != nil {
		panic(err)
	}
	e.conf = conf
	e.Session = session
	return e
}

// New New mongo instance
// Export Session, API and AutoHook
func New() *Mongo {
	mgo := &Mongo{}
	mgo.m = make([]*Profile, 0)
	mgo.API = &API{mgo: mgo, Opts: &Opts{}}
	return mgo
}
