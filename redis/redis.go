// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package redisext

import (
	addition "github.com/2637309949/bulrush-addition"
	"github.com/go-redis/redis"
)

// Redis some common function
type Redis struct {
	Client *redis.Client
	API    *API
}

// New new redis instance
func New() *Redis {
	return &Redis{API: &API{}}
}

// Init redis
func (r *Redis) Init(init func(*Redis)) *Redis {
	init(r)
	return r
}

// Conf set e conf
func (r *Redis) Conf(opts *redis.Options) *Redis {
	client := redis.NewClient(opts)
	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}
	addition.RushLogger.Info("redis:Connection has been established successfully, URL:%v", opts.Addr)
	r.Client = client
	r.API.Client = client
	return r
}
