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
	config *bulrush.Config
}

// New new redis instance
func New(config *bulrush.Config) *Redis {
	client := obClient(config)
	rds := &Redis{
		Client: client,
		Hooks: &Hooks{
			Client: client,
		},
		config: config,
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
