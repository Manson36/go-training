package main

import "fmt"

//定义全局匿名函数
var (
	Fun = func (n1 int, n2 int)int{
		return n1*n2
	}
)
func main(){

	//匿名函数使用
	res:= func(n1 int ,n2 int )int {
		return n1+n2
	}(10,20)
	fmt.Println("res =",res )

	//使用方式2
	a:= func(n1 int ,n2 int ) int{
		return n1-n2
	}
	res2:=a(10,30)
	fmt.Println("res2=",res2)
	res3:=a(90,60)
	fmt.Println("res3=",res3)

	res4:=Fun(1,2)
	fmt.Println("res4=",res4)

}