package main

import "fmt"

func main(){

	//100 以内数求和 当和大于20 求当前数
	sum:=0

	for i:=0;i<=100;i++{
		sum+=i
		if sum>20 {
			fmt.Println("当sum大于20时，当前数是",i)
			break
		}
	}
}