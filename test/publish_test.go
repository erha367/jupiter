package test

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"reflect"
	"strings"
	"testing"
)

func TestPublish(t *testing.T) {
	// redis连接
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// redis订阅
	Subscribe(client)
	client.Close()
}

func Subscribe(client *redis.Client) {
	// 订阅mychannel 确定
	pubsub := client.PSubscribe("__keyevent@*__:expired")
	// 阻塞程序
	for {
		_, err := pubsub.Receive()
		if err != nil {

			return
		}
		ch := pubsub.Channel()
		for msg := range ch {
			fmt.Println(msg.Channel, msg.Payload)
			fmt.Println(reflect.TypeOf(msg.Payload))
			kv := strings.Split(msg.Payload, ``)
			fmt.Println(kv)
		}
	}
}
