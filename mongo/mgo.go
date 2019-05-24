package mongo

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
	API      *API
	AutoHook bulrush.PNBase
}

// New New mongo instance
func New(config *bulrush.Config) *Mongo {
	sess := getSession(config)
	mgo := &Mongo{
		m:       make([]map[string]interface{}, 0),
		cfg:     config,
		Session: sess,
		API:     &API{},
	}
	mgo.API.mgo = mgo
	mgo.AutoHook = autoHook(mgo)
	return mgo
}

// Register model
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
	manifest := funk.Find(mgo.m, func(item map[string]interface{}) bool {
		flag := item["name"].(string) == name
		return flag
	}).(map[string]interface{})
	if manifest != nil {
		target := addition.LeftOkV(manifest["reflector"])
		ojt := addition.CreateSlice(target)
		return ojt
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// Var return  Var
func (mgo *Mongo) Var(name string) interface{} {
	m := funk.Find(mgo.m, func(item map[string]interface{}) bool {
		return item["name"].(string) == name
	}).(map[string]interface{})
	if m != nil {
		entity := addition.LeftOkV(m["reflector"])
		ojt := addition.CreateObject(entity)
		return ojt
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// Model return instance
func (mgo *Mongo) Model(name string) *mgo.Collection {
	m := funk.Find(mgo.m, func(item map[string]interface{}) bool {
		flag := item["name"].(string) == name
		return flag
	}).(map[string]interface{})
	if m != nil {
		db := addition.Some(m["db"], mgo.cfg.GetString("mongo.opts.database", "bulrush")).(string)
		collect := addition.Some(m["collection"], name).(string)
		model := mgo.Session.DB(db).C(collect)
		return model
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// autoHook Automatic routing
func autoHook(mgo *Mongo) bulrush.PNBase {
	return bulrush.PNQuick(func(r *gin.RouterGroup) {
		funk.ForEach(mgo.m, func(item map[string]interface{}) {
			if autoHook, exists := item["autoHook"]; exists == false || autoHook == true {
				name := item["name"].(string)
				mgo.API.List(r, name)
				mgo.API.One(r, name)
				mgo.API.Create(r, name)
				mgo.API.Update(r, name)
				mgo.API.Delete(r, name)
			}
		})
	})
}

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

func getSession(config *bulrush.Config) *mgo.Session {
	if addrs, _ := config.List("mongo.addrs"); addrs != nil && len(addrs) > 0 {
		dial := dialInfo(config)
		return bulrush.LeftSV(mgo.DialWithInfo(dial)).(*mgo.Session)
	}
	panic(fmt.Errorf("mongo.addrs not found"))
}
