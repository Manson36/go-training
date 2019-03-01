package main

import "fmt"

//基本数据类型和数组都是值传递，即在函数内修改不会影响原来的值

func test02(n int){

	n= n + 10
	fmt.Println("test02 n=",n)
}

func main(){

	num:=20
	test02(num)
	fmt.Println("main num =",num )
}