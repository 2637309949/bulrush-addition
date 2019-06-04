package mgo

import (
	"errors"
	"fmt"

	"github.com/2637309949/bulrush"
	addition "github.com/2637309949/bulrush-addition"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/thoas/go-funk"
)

// Mongo Type Defined
type Mongo struct {
	m        []map[string]interface{}
	cfg      *bulrush.Config
	Session  *mgo.Session
	API      *api
	AutoHook bulrush.PNBase
}

// New New mongo instance
// Export Session, API and AutoHook
func New(config *bulrush.Config) *Mongo {
	session := createSession(config)
	mgo := &Mongo{
		m:       make([]map[string]interface{}, 0),
		cfg:     config,
		Session: session,
		API:     &api{},
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
	statement := funk.Find(mgo.m, func(item map[string]interface{}) bool {
		flag := item["name"].(string) == name
		return flag
	}).(map[string]interface{})
	if statement != nil {
		return addition.CreateSlice(addition.LeftOkV(statement["reflector"]))
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// Var return  Var
// reflect from reflector entity
func (mgo *Mongo) Var(name string) interface{} {
	statement := funk.Find(mgo.m, func(item map[string]interface{}) bool {
		return item["name"].(string) == name
	}).(map[string]interface{})
	if statement != nil {
		return addition.CreateObject(addition.LeftOkV(statement["reflector"]))
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// Model return instance
// throw error if not exists these model
func (mgo *Mongo) Model(name string) *mgo.Collection {
	statement := funk.Find(mgo.m, func(item map[string]interface{}) bool {
		return item["name"].(string) == name
	}).(map[string]interface{})
	if statement != nil {
		db := addition.Some(statement["db"], mgo.cfg.Mongo.Database).(string)
		collect := addition.Some(statement["collection"], name).(string)
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
func dialInfo(config *bulrush.Config) *mgo.DialInfo {
	dial := &mgo.DialInfo{}
	dial.Addrs = config.Mongo.Addrs
	dial.Timeout = config.Mongo.Timeout
	dial.Database = config.Mongo.Database
	dial.ReplicaSetName = config.Mongo.ReplicaSetName
	dial.Source = config.Mongo.Source
	dial.Service = config.Mongo.Service
	dial.ServiceHost = config.Mongo.ServiceHost
	dial.Mechanism = config.Mongo.Mechanism
	dial.Username = config.Mongo.Username
	dial.Password = config.Mongo.Password
	dial.PoolLimit = config.Mongo.PoolLimit
	dial.PoolTimeout = config.Mongo.PoolTimeout
	dial.ReadTimeout = config.Mongo.ReadTimeout
	dial.WriteTimeout = config.Mongo.WriteTimeout
	dial.AppName = config.Mongo.AppName
	dial.FailFast = config.Mongo.FailFast
	dial.Direct = config.Mongo.Direct
	dial.MinPoolSize = config.Mongo.MinPoolSize
	dial.MaxIdleTimeMS = config.Mongo.MaxIdleTimeMS
	return dial
}

// obtain mongo connect session
func createSession(cfg *bulrush.Config) *mgo.Session {
	dial := dialInfo(cfg)
	session, err := mgo.DialWithInfo(dial)
	if err != nil {
		panic(err)
	}
	return session
}
