package mongoclient

import (
	"context"
	"fmt"
	testutils "stock-contexts/pkg/shared/utils/test"
	"testing"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func TestConnectToMongo(t *testing.T) {
	testutils.LoadEnv()
	client := NewConnectedMongoDriver()
	err := client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		t.Errorf("Couldn't connect to the database %v", err)
	} else {
		fmt.Println("Connected!")
	}
}
