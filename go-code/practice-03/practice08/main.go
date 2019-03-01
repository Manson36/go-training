package main

import "fmt"
func getSum(n1 int, n2 int )int {

	return n1 + n2
}
//Go支持自定义数据类型

type myFunType func(int,int )int
func myFun2(funvar myFunType, num1 int, num2 int )int {

	return funvar(num1,num2)
}

func main(){

	type myInt int
	var num1 myInt
	var num2 int

	num1= 40
	num2 = int(num1)
	fmt.Println("num1 =",num1,"num2 =",num2)

	res3:= myFun2(getSum,500,600)
	fmt.Println("res3=",res3)
}