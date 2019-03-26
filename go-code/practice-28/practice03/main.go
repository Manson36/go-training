package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {

	var p1 Person
	p1.Age = 10
	p1.Name = "tom"
	var p2 = &p1

	fmt.Println((*p2).Age)
	fmt.Println(p2.Age)

	p2.Name = "Mary"
	fmt.Printf("p2.name=%v ,p1.name=%v",p2.Name,p1.Name)
	fmt.Printf("p2.name=%v ,p1.name=%v",(*p2).Name,p1.Name)

	fmt.Printf("p1的地址%p\n",&p1)
	fmt.Printf("p2的地址%p，p2的值%v\n，",&p2,p2)
}