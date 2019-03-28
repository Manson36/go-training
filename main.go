package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main(){
	fmt.Println(redis.NewClient)
}