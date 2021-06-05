package identityrepo

import (
	"fmt"
	identitydomain "stock-contexts/pkg/customer-identity/domain"
	redisabstractclient "stock-contexts/pkg/shared/infra/repo/redis/abstractClient"
	redisclient "stock-contexts/pkg/shared/infra/repo/redis/config"
	"time"

	"github.com/go-redis/redis/v8"
)

var hashName = "activeUsers"

type IRedisRepo interface {
	GetClient() *redis.Client
	SaveAuthenticatedUser(user identitydomain.User)
	ReadTokens(userName identitydomain.UserName) ([]string, error)
}

type RedisRepo struct {
	Abstract *redisabstractclient.Abstract
}

func (c *RedisRepo) GetClient() *redis.Client {
	return c.Abstract.Client
}

func (c *RedisRepo) ConstructKeyWithToken(user identitydomain.User) string {
	return fmt.Sprintf("refresh-%s.%s.%s", user.GetRefreshToken(), hashName, user.GetUserName().GetValue())
}

func (c *RedisRepo) SaveAuthenticatedUser(user identitydomain.User) {
	m, _ := time.ParseDuration("5h30m")
	err := c.Abstract.SetOne(c.ConstructKeyWithToken(user), user.GetAccessToken(), time.Duration(m.Minutes()))
	fmt.Println("err", err)
}

func (c *RedisRepo) ReadTokens(userName identitydomain.UserName) ([]string, error) {
	a := fmt.Sprintf("*.%s.%s", hashName, userName.GetValue())
	// fmt.Println("search key", a)
	b, err := c.Abstract.GetAllKeys(a)
	fmt.Println("result", b)
	return b, err
}

// NewCRUDRepo define ReadRepo Instance
func NewCRUDRepo() *RedisRepo {
	return &RedisRepo{
		Abstract: &redisabstractclient.Abstract{
			Client: redisclient.NewConnectedRedis(),
		},
	}
}
