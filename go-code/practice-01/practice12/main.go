package main

import "fmt"

func main(){

	//打印1到100之间所有9的倍数的个数和总和
	var count int =0
	var sum int=0

	for i:=1;i<=100;i++{
		if i%9==0{
			count++
			sum+=i
		}
	}
	fmt.Printf("count=%d sum=%d \n ",count,sum)

	var n int=6

	for i:=0; i<= n; i++{
		fmt.Printf("%v +%v = %v \n",i,n-i,n)

	}
}