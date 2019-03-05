package main

import "fmt"

func main(){

	//slice与string
	str := "Hello lilly"
	slice := str[6:]

	fmt.Println("slice=",slice)

	//字符串的修改
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("arr1",arr1)

	str2 := "南京欢迎"
	arr2 := []rune(str2)
	arr2[0] = '北'
	str = string(arr1)
	fmt.Println("arr2",arr2)


}