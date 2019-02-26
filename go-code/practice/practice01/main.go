package main

import (
	"fmt"
	"unsafe"
)

func main(){
	var n =100
	//n是什么类型
	fmt.Printf("n的类型是 %T", n)
//查看某个变量在程序中的类型和占用字节大小
	var n2 int64 =10
	fmt.Printf("n2 的类型是 %T 占用的字节数是 %d", n2,unsafe.Sizeof(n2))
	var price float32 =89.22
	fmt.Println("price =", price)
	var num1 float32 =-0.0000043
	var num2 float64 =-7828080.22
	fmt.Println("num1 =", num1,"num2 =", num2)

}