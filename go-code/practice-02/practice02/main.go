package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	//打印九九乘法表
	var num int = 9

	for i:=1;i<= num ;i++{
		for j := 1; j<= i; j++{
			fmt.Printf("%v * %v = %v",j,i,j*i)
		}
		fmt.Println()
	}

	//break 的快速入门
	//随机生成1-100 的数，直到输出99 停止，计算用了多少次

	var count int = 0

	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100)+ 1
		fmt.Println("n=",n)
		count++
		if n==99{
		break
		}
	}
	fmt.Println("输出99 一共使用",count)
}