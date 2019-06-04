/**
 * @author [double]
 * @email [2637309949@qq.com]
 * @create date 2019-03-19 17:52:35
 * @modify date 2019-03-19 17:52:35
 * @desc [description]
 */

package redis

import (
	"github.com/2637309949/bulrush"
	"github.com/go-redis/redis"
)

// Redis some common function
type Redis struct {
	Client *redis.Client
	API    *API
}

// New new redis instance
func New(cfg *bulrush.Config) *Redis {
	client := createClient(cfg)
	api := &API{
		Client: client,
	}
	return &Redis{
		Client: client,
		API:    api,
	}
}

// dialInfo with default params
func dialInfo(config *bulrush.Config) *redis.Options {
	options := &redis.Options{}
	options.Addr = config.Redis.Addr
	options.Password = config.Redis.Password
	options.DB = config.Redis.DB
	return options
}

// ping client
func ping(c *redis.Client) {
	if _, err := c.Ping().Result(); err != nil {
		panic(err)
	}
}

// createClient obtain a redis connecting
func createClient(config *bulrush.Config) *redis.Client {
	options := dialInfo(config)
	client := redis.NewClient(options)
	ping(client)
	return client
}
