package main

import (
	"fmt"
	"time"
)

//goroutine中使用recover 捕获 punic

func sayHello() {
	for i:= 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}

}
func test() {

	defer func () {
		if err := recover(); err != nil {
			fmt.Println("test发生错误",err)
		}
	}()

	var myMap map[int]string
	myMap[0] = "golang" //error
}

func main() {

	go sayHello()
	go test()

	for i:= 0; i < 10; i++ {
		fmt.Println("main() ok= ",i)
		time.Sleep(time.Second)
	}
}