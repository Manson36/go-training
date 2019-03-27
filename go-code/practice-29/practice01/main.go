package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func initRds() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d","39.104.124.215", 6379),
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalln(err)
	}

	return client
}
