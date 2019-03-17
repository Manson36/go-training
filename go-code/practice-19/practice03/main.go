package main

import (
	"fmt"
	"sync"
	"time"
)

//改进：解决资源争夺的问题，对全局变量加锁
var (
	myMap = make(map[int]int,10)
	lock sync.Mutex
)
func test(n int){

	res :=1
	for i := 1; i< n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {

	//开启多个协程[200]个
	for i:= 1; i<= 200; i++ {
		go test(i)
	}
	time.Sleep(time.Second*10)

	lock.Lock()
	for i,v := range myMap {
		fmt.Printf("myMap[%d]= %d\n",i,v )
	}
	lock.Unlock()
}