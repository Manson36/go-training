package main

import "fmt"

type firstS struct {
	in1 int
	in2 int
}

type secondS struct {
	b int
	c float32
	int //匿名字段
	firstS //匿名字段
}

func main() {

	sec := new(secondS)
	sec.b = 6
	sec.c = 7.5
	sec.int = 60
	sec.in1 = 5
	sec.in2 = 10

	fmt.Println("sec.b is:", sec.b)
	fmt.Println("sec.c is:", sec.c)
	fmt.Println("sec.int is:", sec.int)
	fmt.Println("sec.in1 is:", sec.in1)
	fmt.Println("sec.in2 is:", sec.in2)

	sec2 := secondS{6,7.5, 60, firstS{5, 10}}
	fmt.Println("sec2 is:", sec2)
}
