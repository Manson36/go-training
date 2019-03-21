package main

import "fmt"

var i = 0

func print2(i int) {
	fmt.Println(i)
}

func main() {

	for ; i< 5; i++ {

		defer print2(i)
	}
}