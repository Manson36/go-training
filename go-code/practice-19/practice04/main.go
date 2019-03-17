package main

import "fmt"

func main() {

	//演示一下管道的使用
	var intChan chan int
	intChan = make(chan int ,3)

	fmt.Printf("intChan 的值为%v，本身的地址为%v\n",intChan,&intChan )

	intChan<-10
	num:= 211
	intChan <- num

	fmt.Printf("管道的长度为%v，容量为%v\n",len(intChan),cap(intChan))

	var num2 int
	num2 = <- intChan
	fmt.Println("num2=",num2)
	fmt.Printf("channel len = %v, channel cap = %v",len(intChan),cap(intChan))
}