package main

import "fmt"

//在GO中函数也是一种数据类型
func getSum(n1 int, n2 int )int{

	return n1+n2
}
func main (){

	a:= getSum
	fmt.Printf("a的数据类型是 %T，getSum的数据类型是%T \n",a,getSum)

	res:= a(10,40)
	fmt.Println("res=",res)
}