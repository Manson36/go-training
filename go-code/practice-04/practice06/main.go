package main

import "fmt"

//函数封装jinzita
	func printPyramid(totallevel int ){

		for i := 1; i <= totallevel; i++{

			for k:=1; k <= totallevel-i; k++{
				fmt.Print(" ")
			}
			for j:= 1; j<= 2*i -1;j++{
				fmt.Print("*")
			}
			fmt.Println()
		}
	}

func main(){

	var n int

	fmt.Println("请输入打印层数")
	fmt.Scanln(&n )
	printPyramid(n)

}