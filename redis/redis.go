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
	client := obClient(cfg)
	rds := &Redis{
		Client: client,
		Hooks: &Hooks{
			Client: client,
		},
	}
	return rds
}

// obClient obtain a redis connecting
func obClient(config *bulrush.Config) *redis.Client {
	addrs := config.GetString("redis.addrs", "")
	if addrs != "" {
		options := &redis.Options{}
		options.Addr = addrs
		options.Password = config.GetString("redis.opts.password", "")
		options.DB = config.GetInt("redis.opts.db", 0)
		client := redis.NewClient(options)
		if _, err := client.Ping().Result(); err != nil {
			panic(err)
		}
		return client
	}
	return nil
}
