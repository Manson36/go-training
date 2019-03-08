package main

import "fmt"

type Person struct{

	Name string
	Age int
}
func main() {

	var p1 Person
	p1.Age = 10
	p1.Name = "小明"
	var p2 *Person = &p1

	fmt.Println((*p2).Age)
	fmt.Println(p2.Age)

	p2.Name = "Tom"
	fmt.Printf("p2.name=%v p1.name=%v",p2.Name,p2.Name)
	fmt.Printf("p2.name=%v p1.name=%v",(*p2).Name,p2.Age)

	fmt.Printf("p1地址%v\n",&p1)
	fmt.Printf("p2地址%v，p2的值%v",&p2,p2)

}