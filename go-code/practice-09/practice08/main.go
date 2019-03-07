package main

import "fmt"

type person struct {
	Name string
	Age int
	scores [5]float64
	ptr *int
	slice []int
	map1 map[string]string
}

func main() {

	var p1 person
	fmt.Println(p1)

	if p1.ptr == nil{
		fmt.Println("ok1")
	}
	if p1.slice == nil{
		fmt.Println("ok2")
	}
	if p1.map1 == nil{
		fmt.Println("ok3")
	}
	p1.slice = make([]int,10)
	p1.slice[0] = 100

	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "Tom"

	fmt.Println(p1)
}