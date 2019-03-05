package main

import "fmt"

func main(){

//切片的遍历

//方式1 传统方式
	var arr [5]int = [5]int {12,3,44,55,66}
	slice := arr[1:4]

	for i := 0; i < len(slice); i++{
		fmt.Printf("slice[%d]=%v",i,slice[i])
	}

	fmt.Println()

	for i, v := range slice {
		fmt.Printf("i=%v, v=%v",i,v)
	}

	//切片可以继续切片
	slice2 := slice[1:2]
	slice2[0] = 100

	fmt.Println("slice2=",slice2)
	fmt.Println("slice=",slice)
	fmt.Println("arr=",arr)

	//利用append内置函数，对切片进行追加
	var slice3 []int = []int {11,22,33}

	slice3 = append(slice3,55,55,66)
	fmt.Println("slice3=",slice3)

	slice3 = append(slice3,slice3...)
	fmt.Println("slice3=",slice3)

	//切片的拷贝
	var slice4 []int = []int {1,2,3,4,5}
	var slice5 []int = make([]int,10)
	copy(slice5,slice4)

	fmt.Println("slice4=",slice4)
	fmt.Println("slice5=",slice5)
}