package main

import "fmt"

func main(){

	var n1 int32 = 5
	var n2 int32 = 20

	switch n1{
	case n2,10:
		fmt.Println("ok1")
	case 5:
		fmt.Println("ok2")
	default:
		fmt.Println("没有匹配到")
	}
}