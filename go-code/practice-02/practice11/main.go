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
func main(){

	n1:=10
	test(n1)
	fmt.Println("main() n1=",n1)

	sum:=getsum(10,20)
	fmt.Println("main() sum=",sum)
}