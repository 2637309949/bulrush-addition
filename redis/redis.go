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

// Hooks -
type Hooks struct {
	Client *redis.Client
}

// Redis some common function
type Redis struct {
	Client *redis.Client
	Hooks  *Hooks
}

// New new redis instance
func New(cfg *bulrush.Config) *Redis {
	client := createClient(cfg)
	hooks := &Hooks{
		Client: client,
	}
	return &Redis{
		Client: client,
		Hooks:  hooks,
	}
}

// createClient obtain a redis connecting
func createClient(config *bulrush.Config) *redis.Client {
	addrs := config.GetString("redis.addrs", "")
	options := &redis.Options{}
	options.Addr = addrs
	options.Password = config.GetString("redis.opts.password", "")
	options.DB = config.GetInt("redis.opts.db", 0)
	client := redis.NewClient(options)
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}
