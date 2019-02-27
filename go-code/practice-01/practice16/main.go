package main

import "fmt"

func main(){

	//使用for循环打印金字塔

	//思路1.打印矩形
	//思路2.打印半个金字塔
	//思路3.打印整个金字塔
	//思路4.将层数变成一个变量，变活
	//思路5.打印空心金字塔

	//分析：在打印*号之前先考虑打印空格还是*

	// 第一层和最后一层全是*
	//中间第一个和最后一个是*

	var  totallevel int =3

	//i表示层数
	for i:=1;i<= totallevel;i++{
		//打印*前打印空格
		for k:=1;k<=totallevel-i;k++{
			fmt.Print("")
		}
		//j表示每层打印多少*
		for j:= 1;j<= 2*i -1;j++{
			if j==1||j== 2*i-1||i==totallevel{
				fmt.Print("*")
			}else{
				fmt.Print("")
			}
		}
		fmt.Println()
		
	}
}