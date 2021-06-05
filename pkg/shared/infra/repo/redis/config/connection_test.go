package redisclient

import (
	"context"
	"fmt"
	testutils "stock-contexts/pkg/shared/utils/test"
	"testing"
)

var ctx = context.Background()

func TestNewConnectedRedis(t *testing.T) {
	testutils.LoadEnv()
	client := NewConnectedRedis()
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		t.Errorf("Received %v", err)
	}
	fmt.Println(pong)
}
