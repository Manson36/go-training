package main

import (
	"fmt"
	"strconv"
	"time"
)

func test03(){

		str:= ""
		for i:=0; i<10000; i++{
			str +="hello" + strconv.Itoa(i)
		}
	}
func main(){

	//统计函数Tset03的执行时间
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Printf("执行的时间为%v",end-start)
}