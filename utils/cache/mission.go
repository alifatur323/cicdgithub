package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

type Customer struct {
	ID       uint   `json:"ID"`
	Username string `json:"Username"`
	Email    string `json:"Email"`
}

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := json.Marshal(Customer{ID: 1, Username: "john_doe", Email: "johndoe@email.com"})
	if err != nil {
		fmt.Println(err)
	}

	userID := 1 // Ganti dengan ID data yang ingin Anda hapus
	key := fmt.Sprintf("user:%d", userID)

	val, err := client.Get(key).Result()

	err = client.Del(key).Err()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
	fmt.Println("Success Delete")
}
