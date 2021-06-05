package redisclient

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func NewConnectedRedis() *redis.Client {
	if Client == nil {
		config := NewDBConfig()
		// fmt.Println("config", config)
		database := redis.NewClient(
			&redis.Options{
				Addr:       fmt.Sprintf("%s:%s", config.Addr, config.Port),
				Username:   config.Username,
				Password:   config.Password,
				DB:         config.DBName,
				MaxRetries: 2,
			})
		fmt.Println("Connected to Redis!")
		Client = database
		return database
	}
	return Client
}

// Client is singleton of db client/pool
