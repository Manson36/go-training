package main

import (
	"fmt"
	"time"
)

//使用select解决阻塞问题

func main() {

	intChan := make(chan int ,10)
	for i:= 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string,5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d",i)
	}
	for {
		select {
		case v:= <- intChan :
			fmt.Println("从intChan读取的数据",v)
			time.Sleep(time.Second)
		case v := <- stringChan :
			fmt.Println("从stringchan中读取的数据",v)
			time.Sleep(time.Second)
		default:
			fmt.Println("都读取不到了")
			time.Sleep(time.Second)
			return

		}

	}
}