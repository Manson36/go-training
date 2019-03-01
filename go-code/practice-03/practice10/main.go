package main

import "fmt"

//go支持可变参数

func sum(n1 int ,args... int )int {
	sum:=n1
	//遍历args 0, 4 i = 5 args[5]
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}

	for i := len(args) - 1; i >= 0; i-- {
		sum += args[i]
	}
	return sum
}
func main(){

	res:=sum(11,22,3,3,3,3)
	fmt.Println("res=",res )
}