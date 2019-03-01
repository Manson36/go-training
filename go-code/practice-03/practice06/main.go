package main

import "fmt"

func test03(n *int) {

	*n = *n + 10
	fmt.Println("test03 n=", *n)
}

func main(){

	num:= 20
	test03(&num)
	fmt.Println("main num =",num )
}