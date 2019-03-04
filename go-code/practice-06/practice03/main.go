package main

import "fmt"

func main(){

	var intArr [3]int
	//定义完数组，会有默认值

	intArr[0]= 10
	intArr[1]= 20
	intArr[2]= 30

	fmt.Println(intArr)
	fmt.Printf("intArr的地址%p [0]的地址%p [1]的地址%p [2]的地址%p \n" +
		"",
		&intArr,&intArr[0],&intArr[1],&intArr[2])

	//从终端输入5个成绩保存到数组中

	var score [5]float64

	for i := 0; i<5; i++{
		fmt.Println("请输入第%d个成绩",i+1)
		fmt.Scanln(&score[i])

	}
	for i:=0; i<5; i++{
		fmt.Printf("score[%d]=%v",i,score[i])
	}
}