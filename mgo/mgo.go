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
		db := addition.Some(statement["db"], mgo.cfg.GetString("mongo.opts.database", "bulrush")).(string)
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
	dial.Addrs = config.GetStrList("mongo.addrs", nil)
	dial.Timeout = config.GetDurationFromSecInt("mongo.opts.timeout", 0)
	dial.Database = config.GetString("mongo.opts.database", "")
	dial.ReplicaSetName = config.GetString("mongo.opts.replicaSetName", "")
	dial.Source = config.GetString("mongo.opts.source", "")
	dial.Service = config.GetString("mongo.opts.service", "")
	dial.ServiceHost = config.GetString("mongo.opts.serviceHost", "")
	dial.Mechanism = config.GetString("mongo.opts.mechanism", "")
	dial.Username = config.GetString("mongo.opts.username", "")
	dial.Password = config.GetString("mongo.opts.password", "")
	dial.PoolLimit = config.GetInt("mongo.opts.poolLimit", 0)
	dial.PoolTimeout = config.GetDurationFromSecInt("mongo.opts.poolTimeout", 0)
	dial.ReadTimeout = config.GetDurationFromSecInt("mongo.opts.readTimeout", 0)
	dial.WriteTimeout = config.GetDurationFromSecInt("mongo.opts.writeTimeout", 0)
	dial.AppName = config.GetString("mongo.opts.appName", "")
	dial.FailFast = config.GetBool("mongo.opts.failFast", false)
	dial.Direct = config.GetBool("mongo.opts.direct", false)
	dial.MinPoolSize = config.GetInt("mongo.opts.minPoolSize", 0)
	dial.MaxIdleTimeMS = config.GetInt("mongo.opts.maxIdleTimeMS", 0)
	return dial
}

// obtain mongo connect session
func createSession(config *bulrush.Config) *mgo.Session {
	dial := dialInfo(config)
	session, err := mgo.DialWithInfo(dial)
	if err != nil {
		panic(err)
	}
	return session
}
