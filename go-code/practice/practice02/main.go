package main

import (
	"fmt"
	"unsafe"
)

func main(){
	var b = false
	fmt.Println("b =" , b)
	//b的占用存储空间是多少
	fmt.Printf("bool的存储空间为 %T\n", unsafe.Sizeof(b))
	var address string ="北京长城 110 Hello World"
	fmt.Println(address)

	var a int
	var d float32
	var c float64
	var isMarraied  bool
	var name string

	fmt.Printf("a =%d, d =%v, c =%v, isMarraied =%v,name = %v",a,d,c,isMarraied,name)




}