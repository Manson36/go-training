package main

import (
	"fmt"
	"time"
)

//需求：计算1-200各个数的阶乘，放到map中，最后显示出来，用goroutine完成

//思路：编写函数，计算阶乘，存到map中;启动多个协程，统计的结果放入到map中
var (
	myMap = make(map[int]int,10)
)
func test(n int){

	res :=1
	for i := 1; i< n; i++ {
		res *= i
	}
	myMap[n] = res
}

func main() {

	//开启多个协程[200]个
	for i:= 1; i<= 200; i++ {
		go test(i)
	}
	time.Sleep(time.Second*10)

	for i,v := range myMap {
		fmt.Printf("myMap[%d]= %d",i,v )
	}
}