package main

import (
	"fmt"
)

func main(){

	//四种初始化数组

	var intArr01 [3]int = [3]int{1,2,3}
	fmt.Println("intArr01=",intArr01)

	var intArr02 = [3]int{5,6,7}
	fmt.Println("intArr02=",intArr02)

	//[...]的规定写法
	var intArr03 = [...]int{8,9,10}
	fmt.Println("intArr03=",intArr03)

	var intArr04 = [...]int{1:300,2:600,0:900}
	fmt.Println("intArr04=",intArr04)

	//类型推倒
	var intArr05 = [...]string{1: "Jhon",2: "jack",0: "Mary"}
	fmt.Println("intArr05=",intArr05)

}