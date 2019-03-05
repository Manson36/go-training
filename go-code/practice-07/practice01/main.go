package main

import "fmt"

func main(){

	//切片的入门
	var intArr [5]int = [...]int {1,22,33,45,66}

	slice:= intArr[1:3]

	fmt.Println("intArr=",intArr)
	fmt.Println("slice 的元素",slice)
	fmt.Println("slice的元素个数",len(slice))
	fmt.Println("slice的容量",cap(slice))

	//切片的使用2 make

	var slice2 []float64 = make([]float64,5,10)
	slice2[2]=19
	slice2[4] = 33

	fmt.Println(slice2 )
	fmt.Println("slice2的元素个数",len(slice2))
	fmt.Println("slice2的容量",cap(slice2))

	//切片的使用方式3
	var slice3 []string = []string {"Tom","Mary","Jack"}

	fmt.Println(slice3 )
	fmt.Println("slice2的元素个数",len(slice3))
	fmt.Println("slice2的容量",cap(slice3))

}