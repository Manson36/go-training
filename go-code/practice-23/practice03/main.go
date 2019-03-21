package main

import "fmt"

func main() {

	var x interface{}
	x = 1
	switch i := x.(type) {
	case nil:
		fmt.Printf("这里是nil，x的类型是%T",i)
	case int :
		fmt.Printf("这里是int，x的类型是%T",i)
	case float64:
		fmt.Printf("这里是float64，x的类型是%T",i)
	case bool :
		fmt.Printf("这里是bool，x的类型是%T",i)
	case string :
		fmt.Printf("这里是string，x的类型是%T",i)
	default:
		fmt.Println("未知类型")
	}
}