package main

import "fmt"

func main() {

	fmt.Println("return:",a())
}

func a() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:",i)
	}()

	defer func() {
		i++
		fmt.Println("defer1:",i)
	}()
	return i
}