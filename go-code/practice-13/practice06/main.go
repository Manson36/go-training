package main

import "fmt"

//打印空心金字塔
func main() {

	var totalLevel int = 20

	for i := 1; i <= totalLevel; i++ {
		for k := 1; k <= totalLevel - i; k++{
			fmt.Print(" ")
		}
		for j:= 1; j <= 2*i-1; j++ {
			if j ==1 || j == 2*i-1 || i == totalLevel {
				fmt.Print("*")
			}else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	//打印九九乘法表
	var num  int = 9

	for i:= 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v",j,i,j*i )
		}
		fmt.Println()
	}
}