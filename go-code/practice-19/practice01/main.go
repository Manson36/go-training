package main

import (
	"fmt"
	"strconv"
	"time"
)

//goroutine快速入门
func Test() {
	for i:= 0; i< 10; i++ {
		fmt.Println("hello world"+ strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {

	go Test()

	for i:= 0; i< 10; i++ {
		fmt.Println("hello golang"+ strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}