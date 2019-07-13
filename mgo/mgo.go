// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mgoext

import (
	"errors"
	"fmt"
	"time"

	addition "github.com/2637309949/bulrush-addition"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/thoas/go-funk"
)

type (
	// Mongo Type Defined
	Mongo struct {
		m       []*Profile
		cfg     *Config
		Session *mgo.Session
		API     *API
	}
	// Config defined mgo config
	Config struct {
		Addrs          []string      `json:"addrs" yaml:"addrs"`
		Timeout        time.Duration `json:"timeout" yaml:"timeout"`
		Database       string        `json:"database" yaml:"database"`
		ReplicaSetName string        `json:"replicaSetName" yaml:"replicaSetName"`
		Source         string        `json:"source" yaml:"source"`
		Service        string        `json:"service" yaml:"service"`
		ServiceHost    string        `json:"serviceHost" yaml:"serviceHost"`
		Mechanism      string        `json:"mechanism" yaml:"mechanism"`
		Username       string        `json:"username" yaml:"username"`
		Password       string        `json:"password" yaml:"password"`
		PoolLimit      int           `json:"poolLimit" yaml:"poolLimit"`
		PoolTimeout    time.Duration `json:"poolTimeout" yaml:"poolTimeout"`
		ReadTimeout    time.Duration `json:"readTimeout" yaml:"readTimeout"`
		WriteTimeout   time.Duration `json:"writeTimeout" yaml:"writeTimeout"`
		AppName        string        `json:"appName" yaml:"appName"`
		FailFast       bool          `json:"failFast" yaml:"failFast"`
		Direct         bool          `json:"direct" yaml:"direct"`
		MinPoolSize    int           `json:"minPoolSize" yaml:"minPoolSize"`
		MaxIdleTimeMS  int           `json:"maxIdleTimeMS" yaml:"maxIdleTimeMS"`
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
		db := addition.Some(m.DB, ext.cfg.Database).(string)
		collect := addition.Some(m.Collection, name).(string)
		return ext.Session.DB(db).C(collect)
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// dialInfo with default params
func dialInfo(conf *Config) *mgo.DialInfo {
	dial := &mgo.DialInfo{}
	dial.Addrs = conf.Addrs
	dial.Timeout = conf.Timeout
	dial.Database = conf.Database
	dial.ReplicaSetName = conf.ReplicaSetName
	dial.Source = conf.Source
	dial.Service = conf.Service
	dial.ServiceHost = conf.ServiceHost
	dial.Mechanism = conf.Mechanism
	dial.Username = conf.Username
	dial.Password = conf.Password
	dial.PoolLimit = conf.PoolLimit
	dial.PoolTimeout = conf.PoolTimeout
	dial.ReadTimeout = conf.ReadTimeout
	dial.WriteTimeout = conf.WriteTimeout
	dial.AppName = conf.AppName
	dial.FailFast = conf.FailFast
	dial.Direct = conf.Direct
	dial.MinPoolSize = conf.MinPoolSize
	dial.MaxIdleTimeMS = conf.MaxIdleTimeMS
	return dial
}

// Conf defined conf
func (ext *Mongo) Conf(conf *Config) *Mongo {
	dial := dialInfo(conf)
	session, err := mgo.DialWithInfo(dial)
	if err != nil {
		panic(err)
	}
	ext.cfg = conf
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
