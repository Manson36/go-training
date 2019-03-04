package main

import (
	"fmt"
)

//编写一个函数，输出100-500之间的所有所有素数，每行显示5个；并求和

func main(){
	var sum=0
	var t=0

	for i:=100; i<500; i++{
		var j =2
		for j=2; j<i; j++ {

			if i%j == 0 {
				break
			}
		}
			if i==j{
				fmt.Printf("%d",i)
				t++
				sum += i
				if t%5==0{
					fmt.Println()
				}else{
					fmt.Printf(" ")
				}
			}

	}
	fmt.Println("总和为",sum)
}