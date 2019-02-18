package addition

import (
	"github.com/2637309949/bulrush"
	"github.com/go-redis/redis"
)

// rdsHooks -
type rdsHooks struct {
	Client *redis.Client
}

// Rds some common function
type Rds struct {
	Client *redis.Client
	Hooks  *rdsHooks
	config *bulrush.Config
}

// NewRds new redis instance
func NewRds(config *bulrush.Config) *Rds {
	client := obClient(config)
	rds := &Rds{
		Client: client,
		Hooks: &rdsHooks{
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
