package main

import "fmt"

func main(){

	//二维数组
	var arr [4][6]int
	//赋初值
	arr[1][3] =1
	arr[2][1] =2
	arr[2][3] =3

	for i:=0; i<4; i++{
		for j:=0; j < 6; j++{
			fmt.Print(arr[i][j]," ")
		}
		fmt.Println()
	}
}