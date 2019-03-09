package main

import "fmt"

type A struct{
	Name string
	age int
}

func (a *A)Sayok() {
	fmt.Println("A sayok",a.Name)
}
func (a *A)hello() {
	fmt.Println("A hello",a.Name)
}

type B struct {
	A
}

func main() {

	var b B
	b.A.Name = "Tom"
	b.A.age = 19
	b.A.Sayok()
	b.A.hello()
}