package main

import "fmt"

//定义全局变量
	var age int =50
	var Name string = "Tom"

	//函数
	func test(){
		//函数内的定义变量只在函数内生效
		age:= 10
		Name:= "Jack"
		fmt.Println("age =",age)
		fmt.Println("Name =",Name)
	}

func main(){

	fmt.Println("age =",age)
	fmt.Println("Name =",Name)

	test()

	//变量在代码块中，那么作用域就在代码块

	for i:= 1;i< 10;i++{
		fmt.Println("i=",i )

	}
	var i int //局部变量
	for i=1; i< 10; i++{
		fmt.Println("i=",i )
	}

	fmt.Println("i =",i )
}