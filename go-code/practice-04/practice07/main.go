package main

import "fmt"

//编写函数调用九九乘法表
	func printMulit(num int ){

		for i:= 1; i <= num; i++{

			for j:=1; j<= i; j++{
				fmt.Printf("%v * %v = %v ",j,i,j*i)
			}
			fmt.Println()
		}
	}

func main(){

	var n int

	fmt.Println("请输入打印层数")
	fmt.Scanln(&n )
	printMulit(n )

}