// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mgo

import (
	"errors"
	"fmt"
	"time"

	"github.com/2637309949/bulrush"
	addition "github.com/2637309949/bulrush-addition"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/thoas/go-funk"
)

// Mongo Type Defined
type Mongo struct {
	m        []map[string]interface{}
	cfg      *conf
	Session  *mgo.Session
	API      *api
	AutoHook bulrush.PNBase
}

type conf struct {
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

// New New mongo instance
// Export Session, API and AutoHook
func New(config *bulrush.Config) *Mongo {
	conf := &conf{}
	if err := config.Unmarshal(conf); err != nil {
		panic(err)
	}
	session := createSession(conf)
	mgo := &Mongo{
		m:       make([]map[string]interface{}, 0),
		cfg:     conf,
		API:     &api{},
		Session: session,
	}
	mgo.API.mgo = mgo
	mgo.AutoHook = autoHook(mgo)
	return mgo
}

// Register model
// should provide name and reflector paramters
func (mgo *Mongo) Register(manifest map[string]interface{}) {
	if _, ok := manifest["name"]; !ok {
		panic(errors.New("name params must be provided"))
	}
	if _, ok := manifest["reflector"]; !ok {
		panic(errors.New("reflector params must be provided"))
	}
	mgo.m = append(mgo.m, manifest)
}

// Vars return array of Var
func (mgo *Mongo) Vars(name string) interface{} {
	m := funk.Find(mgo.m, func(item map[string]interface{}) bool {
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
func (mgo *Mongo) Var(name string) interface{} {
	m := funk.Find(mgo.m, func(item map[string]interface{}) bool {
		return item["name"].(string) == name
	}).(map[string]interface{})
	if m != nil {
		return addition.CreateObject(addition.LeftOkV(m["reflector"]))
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// Model return instance
// throw error if not exists these model
func (mgo *Mongo) Model(name string) *mgo.Collection {
	m := funk.Find(mgo.m, func(item map[string]interface{}) bool {
		return item["name"].(string) == name
	}).(map[string]interface{})
	if m != nil {
		db := addition.Some(m["db"], mgo.cfg.Database).(string)
		collect := addition.Some(m["collection"], name).(string)
		return mgo.Session.DB(db).C(collect)
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// autoHook Automatic routing
func autoHook(mgo *Mongo) bulrush.PNBase {
	return bulrush.PNQuick(func(r *gin.RouterGroup) {
		funk.ForEach(mgo.m, func(item map[string]interface{}) {
			if autoHook, exists := item["autoHook"]; exists == false || autoHook == true {
				collection := item["name"].(string)
				mgo.API.ALL(r, collection)
			}
		})
	})
}

// dialInfo with default params
func dialInfo(conf *conf) *mgo.DialInfo {
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

// obtain mongo connect session
func createSession(cfg *conf) *mgo.Session {
	dial := dialInfo(cfg)
	session, err := mgo.DialWithInfo(dial)
	if err != nil {
		panic(err)
	}
	return session
}
