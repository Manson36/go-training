package main

import "fmt"

func main(){

	//内置函数new的使用
	num1 := 100
	fmt.Printf("num1的类型%T，值%v，地址%v\n",num1,num1,&num1)

	num2 := new(int) //*int
	//num2的类型是 指针
	//num2的值 是 地址
	//num2的 地址是 地址
	*num2 = 1000
	fmt.Printf("num2的类型%T，值%v，地址%v\n",num2,num2,&num2)
	fmt.Printf("num2指向的值是%v ",*num2)
}