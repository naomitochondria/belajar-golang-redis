package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0,
})

// func TestConnection(t *testing.T) {
// 	assert.NotNil(t, client)

// 	err := client.Close()
// 	assert.Nil(t, err)
// }

var ctx = context.Background()

func TestStoreAccess(t *testing.T) {
	var (
		key         string = "key"
		value       string = "value"
		timeExpired int    = 3
		timeSleep   int    = 10
	)

	client.SetEx(ctx, key, value, time.Second*time.Duration(timeExpired))

	time.Sleep(time.Duration(timeSleep) * time.Second)

	result, err := client.Get(ctx, key).Result()
	fmt.Println("err: ", err)
	fmt.Println("result: ", result == "")
}

func TestPing(t *testing.T) {
	result, err := client.Ping(ctx).Result()
	assert.Nil(t, err)

	assert.Equal(t, "PONG", result)
}

// func TestString(t *testing.T) {
// 	client.SetEx(ctx, "name", "Caca", time.Second*3)

// 	result, err := client.Get(ctx, "name").Result()
// 	assert.Nil(t, err)
// 	assert.Equal(t, "Caca", result)

// 	time.Sleep(5 * time.Second)
// 	result, err = client.Get(ctx, "name").Result()
// 	assert.NotNil(t, err)
// 	assert.Empty(t, result)
// }

// func TestList(t *testing.T) {
// 	client.RPush(ctx, "names", "A")
// 	client.RPush(ctx, "names", "B")
// 	client.RPush(ctx, "names", "C")

// 	assert.Equal(t, "A", client.LPop(ctx, "names").Val())
// 	assert.Equal(t, "B", client.LPop(ctx, "names").Val())
// 	assert.Equal(t, "C", client.LPop(ctx, "names").Val())

// 	client.Del(ctx, "names")
// }

// func TestSet(t *testing.T) {
// 	client.SAdd(ctx, "students", "F")
// 	client.SAdd(ctx, "students", "F")
// 	client.SAdd(ctx, "students", "G")
// 	client.SAdd(ctx, "students", "H")

// 	assert.Equal(t, int64(3), client.SCard(ctx, "students").Val())
// 	assert.Equal(t, []string{"F", "G", "H"}, client.SMembers(ctx, "students").Val())
// }

// func TestSortedSet(t *testing.T) {
// 	client.ZAdd(ctx, "scores", redis.Z{Score: 100, Member: "Eko"})
// 	client.ZAdd(ctx, "scores", redis.Z{Score: 85, Member: "Budi"})
// 	client.ZAdd(ctx, "scores", redis.Z{Score: 95, Member: "Joko"})

// 	assert.Equal(t, []string{"Budi", "Joko", "Eko"},
// 		client.ZRange(ctx, "scores", 0, -1).Val())
// 	assert.Equal(t, "Eko", client.ZPopMax(ctx, "scores").Val()[0].Member)
// 	assert.Equal(t, "Joko", client.ZPopMax(ctx, "scores").Val()[0].Member)
// 	assert.Equal(t, "Budi", client.ZPopMax(ctx, "scores").Val()[0].Member)
// }

// func TestHash(t *testing.T) {
// 	client.HSet(ctx, "user:1", "id", "1")
// 	client.HSet(ctx, "user:1", "name", "Eko")
// 	client.HSet(ctx, "user:1", "email", "eko@gmail.com")

// 	user := client.HGetAll(ctx, "user:1").Val()
// 	assert.Equal(t, "1", user["id"])
// 	assert.Equal(t, "Eko", user["name"])
// 	assert.Equal(t, "eko@gmail.com", user["email"])

// 	client.Del(ctx, "user:1")
// }

// func TestGeoPoint(t *testing.T) {
// 	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
// 		Name:      "Toko A",
// 		Longitude: 106.818489,
// 		Latitude:  -6.178966,
// 	})
// 	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
// 		Name:      "Toko B",
// 		Longitude: 106.821568,
// 		Latitude:  -6.180662,
// 	})

// 	distance := client.GeoDist(ctx, "sellers", "Toko A", "Toko B", "km").Val()
// 	assert.Equal(t, 0.3892, distance)

// 	sellers := client.GeoSearch(ctx, "sellers", &redis.GeoSearchQuery{
// 		Longitude:  106.819143,
// 		Latitude:   -6.180182,
// 		Radius:     5,
// 		RadiusUnit: "km",
// 	}).Val()

// 	assert.Equal(t, []string{"Toko A", "Toko B"}, sellers)
// }

// func TestHyperLogLog(t *testing.T) {
// 	client.PFAdd(ctx, "visitors", "A", "B", "C")
// 	client.PFAdd(ctx, "visitors", "D", "A", "C")
// 	client.PFAdd(ctx, "visitors", "C", "A", "B")

// 	assert.Equal(t, int64(4), client.PFCount(ctx, "visitors").Val())
// }

// func TestPipeline(t *testing.T) {
// 	client.Pipelined(ctx, func(pipeliner redis.Pipeliner) error {
// 		pipeliner.SetEx(ctx, "name", "A", time.Second*3)
// 		pipeliner.SetEx(ctx, "address", "B", time.Second*2)

// 		return nil
// 	})

// 	assert.Equal(t, "A", client.Get(ctx, "name").Val())
// 	assert.Equal(t, "B", client.Get(ctx, "address").Val())
// }

// func TestTransaction(t *testing.T) {
// 	client.TxPipelined(ctx, func(pipeliner redis.Pipeliner) error {
// 		pipeliner.SetEx(ctx, "name", "A", time.Second*3)
// 		pipeliner.SetEx(ctx, "address", "B", time.Second*2)

// 		return nil
// 	})

// 	assert.Equal(t, "A", client.Get(ctx, "name").Val())
// 	assert.Equal(t, "B", client.Get(ctx, "address").Val())
// }

// func TestStream(t *testing.T) {
// 	for i := 1; i <= 10; i++ {
// 		client.XAdd(ctx, &redis.XAddArgs{
// 			Stream: "members",
// 			Values: map[string]interface{}{
// 				"numberKey":   i,
// 				"numberValue": i * i,
// 			},
// 		})
// 	}
// }

// func TestCreateConsumerGroup(t *testing.T) {
// 	client.XGroupCreate(ctx, "members", "group-1", "0")
// 	client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-1")
// 	client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-2")
// }

// func TestGetStream(t *testing.T) {
// 	result := client.XReadGroup(ctx, &redis.XReadGroupArgs{
// 		Group:    "group-1",
// 		Consumer: "consumer-1",
// 		Streams:  []string{"members", ">"},
// 		Count:    6,
// 		Block:    time.Second * 5,
// 	}).Val()

// 	for _, stream := range result {
// 		for _, message := range stream.Messages {
// 			fmt.Println(message.Values)
// 		}
// 	}
// }

// func TestPublishPubSub(t *testing.T) {
// 	err := client.Publish(ctx, "channel-1", "Subscribe these!").Err()
// 	assert.Nil(t, err)

// 	subscriber := client.Subscribe(ctx, "channel-1")
// 	defer subscriber.Close()

// 	message, err := subscriber.ReceiveMessage(ctx)
// 	assert.Nil(t, err)
// 	fmt.Println("Accepting:", message.Payload)
// }
