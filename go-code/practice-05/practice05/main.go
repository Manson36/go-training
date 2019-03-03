package main

import (
	"fmt"
	"time"
)

	//sleep使用时间常亮
func main(){

	i := 0
	for {
		i++
		fmt.Println(i)
		//休眠
		time.Sleep(time.Millisecond* 100)

		if i==100{

			break
		}
	}
}