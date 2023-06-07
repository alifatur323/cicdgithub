package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func main() {
	// connect to redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis port address
		Password: "",               // redis password
		DB:       0,                // index db redis
	})

	// close connection
	defer func(client *redis.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)

	err := client.Set(context.Background(), "key", "Alifatur", time.Hour).Err()

	if err != nil {
		log.Fatal(err)
	}

	// fetch data from redis
	val, err := client.Get(context.Background(), "key").Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("data dari redis", val)

}
