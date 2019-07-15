// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mgoext

import (
	"errors"
	"fmt"

	addition "github.com/2637309949/bulrush-addition"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
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
func (ext *Mongo) Docs() *[]Doc {
	docs := []Doc{}
	funk.ForEach(ext.m, func(item *Profile) {
		docs = append(docs, *item.docs...)
	})
	return &docs
}

// Plugin defined plugin for bulrush
func (ext *Mongo) Plugin(r *gin.RouterGroup) *Mongo {
	funk.ForEach(ext.m, func(item *Profile) {
		if !item.BanHook {
			ext.API.ALL(r, item.Name)
		}
	})
	return ext
}

// Init mgo
func (ext *Mongo) Init(init func(*Mongo)) *Mongo {
	init(ext)
	return ext
}

// Register model
// should provide name and reflector paramters
func (ext *Mongo) Register(profile *Profile) *Mongo {
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
func (ext *Mongo) Profile(name string) *Profile {
	if m := funk.Find(ext.m, func(item *Profile) bool {
		return item.Name == name
	}); m != nil {
		return m.(*Profile)
	}
	return nil
}

// Vars return array of Var
func (ext *Mongo) Vars(name string) interface{} {
	m := ext.Profile(name)
	if m != nil {
		return addition.CreateSlice(m.Reflector)
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// Var return  Var
// reflect from reflector entity
func (ext *Mongo) Var(name string) interface{} {
	m := ext.Profile(name)
	if m != nil {
		return addition.CreateObject(m.Reflector)
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// Model return instance
// throw error if not exists these model
func (ext *Mongo) Model(name string) *mgo.Collection {
	m := ext.Profile(name)
	if m != nil {
		db := addition.Some(m.DB, ext.conf.Database).(string)
		collect := addition.Some(m.Collection, name).(string)
		return ext.Session.DB(db).C(collect)
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// Conf defined conf
func (ext *Mongo) Conf(conf *mgo.DialInfo) *Mongo {
	session, err := mgo.DialWithInfo(conf)
	if err != nil {
		panic(err)
	}
	ext.conf = conf
	ext.Session = session
	return ext
}

// New New mongo instance
// Export Session, API and AutoHook
func New() *Mongo {
	mgo := &Mongo{}
	mgo.m = make([]*Profile, 0)
	mgo.API = &API{mgo: mgo, Opts: &Opts{}}
	return mgo
}
