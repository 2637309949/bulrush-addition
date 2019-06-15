// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package redis

import (
	"time"

	"github.com/2637309949/bulrush"
	"github.com/go-redis/redis"
)

// Redis some common function
type Redis struct {
	Client *redis.Client
	API    *API
}

type conf struct {
	Network            string        `json:"network" yaml:"network"`
	Addr               string        `json:"addrs" yaml:"addrs"`
	Password           string        `json:"password" yaml:"password"`
	DB                 int           `json:"db" yaml:"db"`
	MaxRetries         int           `json:"maxRetries" yaml:"maxRetries"`
	MinRetryBackoff    time.Duration `json:"minRetryBackoff" yaml:"minRetryBackoff"`
	MaxRetryBackoff    time.Duration `json:"maxRetryBackoff" yaml:"maxRetryBackoff"`
	DialTimeout        time.Duration `json:"dialTimeout" yaml:"dialTimeout"`
	ReadTimeout        time.Duration `json:"readTimeout" yaml:"readTimeout"`
	WriteTimeout       time.Duration `json:"writeTimeout" yaml:"writeTimeout"`
	PoolSize           int           `json:"poolSize" yaml:"poolSize"`
	MinIdleConns       int           `json:"minIdleConns" yaml:"minIdleConns"`
	MaxConnAge         time.Duration `json:"maxConnAge" yaml:"maxConnAge"`
	PoolTimeout        time.Duration `json:"poolTimeout" yaml:"poolTimeout"`
	IdleTimeout        time.Duration `json:"idleTimeout" yaml:"idleTimeout"`
	IdleCheckFrequency time.Duration `json:"idleCheckFrequency" yaml:"idleCheckFrequency"`
}

// New new redis instance
func New(bulCfg *bulrush.Config) *Redis {
	cf, err := bulCfg.Unmarshal("redis", conf{})
	if err != nil {
		panic(err)
	}
	conf := cf.(conf)
	client := createClient(&conf)
	api := &API{
		Client: client,
	}
	return &Redis{
		Client: client,
		API:    api,
	}
}

// dialInfo with default params
func dialInfo(conf *conf) *redis.Options {
	options := &redis.Options{}
	options.Addr = conf.Addr
	options.Password = conf.Password
	options.DB = conf.DB
	return options
}

// ping client
func ping(c *redis.Client) {
	if _, err := c.Ping().Result(); err != nil {
		panic(err)
	}
}

// createClient obtain a redis connecting
func createClient(conf *conf) *redis.Client {
	options := dialInfo(conf)
	client := redis.NewClient(options)
	ping(client)
	return client
}
