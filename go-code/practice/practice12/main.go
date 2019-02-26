package main

import "fmt"

func main(){

	//两个变量a和b，要求不使用中间变量，将两个值交换
	var a int =10
	var b int =20
	a= a+b
	b= a-b
	a=a-b
	fmt.Printf(" a =%v b=%v",a,b)

	//其他运算符说明
	c:=100
	fmt.Println("\n c的地址是",&c)
	var ptr *int=&c
	fmt.Println("ptr指向的值是",*ptr)
}