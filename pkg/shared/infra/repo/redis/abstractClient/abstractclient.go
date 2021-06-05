package redisabstractclient

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// Abstract define an abstaction layer for connecting redis
type Abstract struct {
	Client *redis.Client
}

func (a Abstract) HandleErr(err error) error {
	var message string
	if err == nil {
		return nil
	}
	switch {
	case err == redis.Nil:
		message = "key does not exist"
	case err != nil:
		fmt.Println("error message", err)
		message = err.Error()
	default:
		message = "Unexpected Error"
	}
	return fmt.Errorf("Message: %s", message)
}

func (a Abstract) GetOne(key string) (string, error) {
	val, err := a.Client.Get(ctx, key).Result()
	if val == "" {
		return "", fmt.Errorf("Message Value is empty")
	}
	return val, a.HandleErr(err)
}

func (a Abstract) GetAllKeys(key string) ([]string, error) {
	var empty []string
	val, err := a.Client.Keys(ctx, key).Result()
	if len(val) == 0 {
		return empty, fmt.Errorf("Message Value Slice is empty")
	}
	return val, a.HandleErr(err)
}

func (a Abstract) Exists(key string) (bool, string) {
	val, _ := a.GetOne(key)
	if val != "" {
		return true, val
	}
	return false, ""
}

func (a Abstract) getAllKeyValue() {}

func (a *Abstract) SetOne(key string, value interface{}, time time.Duration) error {
	err := a.Client.Set(ctx, key, value, 0).Err()
	return a.HandleErr(err)
}

func (a *Abstract) Delete(key string) error {
	err := a.Client.Del(ctx, key).Err()
	return a.HandleErr(err)
}
