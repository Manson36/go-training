package main

import "fmt"

type Student struct {
	Name string
	Age int
}

func (stu *Student) String() string {

	str := fmt.Sprintf("Name=[%v] Age=[%v]",stu.Name,stu.Age)
	return str
}

func main() {

	stu := Student{
		Name:"tom",
		Age: 12,
	}
	fmt.Println(&stu)
}