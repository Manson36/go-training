package main

import (
	"fmt"
)

//函数的defer将语句放入栈时，也会将相关的值同时拷贝入栈

func sum(n1 int, n2 int )int {

	defer fmt.Println("ok1 ,n1 =",n1)
	defer fmt.Println("ok2 ,n2 =",n2)

	n1++//n1 = 11
	n2++//n2 =21
	res:= n1 + n2//res = 32
	fmt.Println("ok3 ,res =",res)
	return res

}

func main(){

	res := sum(10,20)
	fmt.Println("res=",res)
}