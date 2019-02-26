package main

import "fmt"

func main(){
	//golang 的指针类型
	var i int =10
	fmt.Println("i的地址=",&i)
	var ptr *int =&i
	fmt.Printf("ptr=%v\n",ptr)
	fmt.Printf("ptr的地址=%v\n",&ptr)
	fmt.Printf("ptr指向的值是= %v \n",*ptr)

	var num int = 9
	fmt.Printf("num address =%v \n",&num)

	var ptr1 *int
	ptr1 = &num
	*ptr1 = 10
	fmt.Println("num=",num)


}