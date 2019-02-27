package main

import "fmt"

func main(){

	//switch 穿透 fallthrough 不判断case条件 直接输出下一个内容
	var num int =10

	switch num{
	case 10:
		fmt.Println("ok1")
		fallthrough
	case 20:
		fmt.Println("ok2")
		fallthrough
	case 30:
		fmt.Println("ok3")
	}

	var x interface{}
	var y =10.0
	x=y

	switch i:= x.(type) {
	case nil:
		fmt.Printf("x 的类型是：%T",i)
	case int:
		fmt.Printf("x 的类型是int")
	case float64:
		fmt.Printf("x 的类型是float64")
	case func(int) float64:
		fmt.Printf("x 的类型是func(int)")
	case bool, string:
		fmt.Printf("x 的类型是bool或string")
	default:
		fmt.Printf("未知型")

	}
}