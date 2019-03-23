package main

import "fmt"

func main() {

	fmt.Println("return:",a())
}

func a() (i int) {

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