package mongo

import (
	"errors"
	"fmt"

	"github.com/2637309949/bulrush"
	addition "github.com/2637309949/bulrush-addition"
	"github.com/globalsign/mgo"
)

// Mongo Type Defined
type Mongo struct {
	config    *bulrush.Config
	Session   *mgo.Session
	Hooks     *Hook
	manifests []interface{}
}

// New New mongo instance
func New(config *bulrush.Config) *Mongo {
	session := getSession(config)
	mgo := &Mongo{
		Hooks:     &Hook{},
		Session:   session,
		manifests: make([]interface{}, 0),
		config:    config,
	}
	mgo.Hooks.One = one(mgo)
	mgo.Hooks.List = list(mgo)
	mgo.Hooks.Create = create(mgo)
	mgo.Hooks.Update = update(mgo)
	mgo.Hooks.Delete = delete(mgo)
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
	mgo.manifests = append(mgo.manifests, manifest)
}

// Vars return array of Var
func (mgo *Mongo) Vars(name string) interface{} {
	manifest := bulrush.Find(mgo.manifests, func(item interface{}) bool {
		flag := item.(map[string]interface{})["name"].(string) == name
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
	manifest := bulrush.Find(mgo.manifests, func(item interface{}) bool {
		flag := item.(map[string]interface{})["name"].(string) == name
		return flag
	}).(map[string]interface{})
	if manifest != nil {
		entity := addition.LeftOkV(manifest["reflector"])
		ojt := addition.CreateObject(entity)
		return ojt
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// Model return instance
func (mgo *Mongo) Model(name string) *mgo.Collection {
	manifest := bulrush.Find(mgo.manifests, func(item interface{}) bool {
		flag := item.(map[string]interface{})["name"].(string) == name
		return flag
	}).(map[string]interface{})
	if manifest != nil {
		db := addition.Some(manifest["db"], mgo.config.GetString("mongo.opts.database", "bulrush")).(string)
		collect := addition.Some(manifest["collection"], name).(string)
		model := mgo.Session.DB(db).C(collect)
		return model
	}
	panic(fmt.Errorf("manifest %s not found", name))
}

// getMgoCfg create mgo config
func getMgoCfg(config *bulrush.Config) *mgo.DialInfo {
	addrs := config.GetStrList("mongo.addrs", nil)
	dial := &mgo.DialInfo{}
	dial.Addrs = addrs
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

// getSession obtain session
func getSession(config *bulrush.Config) *mgo.Session {
	addrs, _ := config.List("mongo.addrs")
	if addrs != nil && len(addrs) > 0 {
		dial := getMgoCfg(config)
		session := bulrush.LeftSV(mgo.DialWithInfo(dial)).(*mgo.Session)
		return session
	}
	return nil
}
