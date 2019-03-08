package main

import "fmt"

type integer int

func (i integer) print() {
	fmt.Println("i=",i)
}

func (i *integer) change() {
	*i = *i + 1
}

func main(){

	var i integer =10

	i.print()
	i.change()
	fmt.Println("i =",i)
}