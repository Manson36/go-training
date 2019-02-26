package main

import "fmt"

func test() bool{
	fmt.Println("Test...")
	return true
}
func main(){
//说明 短路与 短路或
var i int =10
if i<9 && test(){
	fmt.Println("ok")
}
if i>9 || test(){
	fmt.Println("hello")
}

//赋值运算符的使用
	a:=9
	b:=2
	fmt.Printf("交换前的情况是 a=%v b=%v \n",a,b)
	//定义一个临时变量
	t:=a
	a=b
	b=t
	fmt.Printf("交换后的情况是 a=%v b=%v \n",a,b)

	a+=17
	fmt.Println("a=",a)
}