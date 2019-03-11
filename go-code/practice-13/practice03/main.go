package main

import "fmt"

func main() {

	var x interface{}
	var b float32 = 2.2
	x = b

	if y ,ok := x.(float32); ok {
		fmt.Println("convert success")
		fmt.Println("y 的类型是 %T 值是 %v",y,y)
	}else{
		fmt.Println("convert fail")
	}
	fmt.Println("继续执行")
}