package main

import "fmt"

//闭包 累加器
func Addupper() func (int)int{

	var n int =10
	var str string ="Hello"

	return func(x int) int {
		n= n+ x
		str+= string(36)
		fmt.Println("str=",str)
		return n
	}
}
func main(){

	f:= Addupper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println(f(4))
}