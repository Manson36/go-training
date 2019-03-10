package main

import "fmt"

type A struct {
	name string
	age int
}

type B struct {
	name string
	Score float64
}

type C struct {
	A
	B
}

func main() {

	var c C
	c.A.name = "tom"
	fmt.Println(c)
}