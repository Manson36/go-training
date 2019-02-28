package main

import "fmt"

//一个函数 test
func test(n1 int){

	n1=n1+1
	fmt.Println("test() n1=",n1)
}

//一个函数 getsum
func getsum(n1 int,n2 int) int {

	sum := n1 + n2
	fmt.Println("getsum sum=", sum)
	return sum
}

//编写函数 可以计算两个数的和和差，并返回结果
func getSumandSub(n1 int, n2 int)(int, int){

	sum:= n1+ n2
	sub:= n1- n2
	return sum,sub
}
func main(){

	n1:=10
	test(n1)
	fmt.Println("main() n1=",n1)

	sum:=getsum(10,20)
	fmt.Println("main() sum=",sum)

	res1,res2:=getSumandSub(1,2)
	fmt.Printf("res1=%v res2=%v",res1,res2)
}