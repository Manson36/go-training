package main

import "fmt"

func main(){

	//continue 实现 打印1-100的奇数
	for i:=1; i<= 100; i++{
		if i%2 ==0{
			continue
		}
		fmt.Println("奇数是",i)
	}

	//从键盘输入不确定个数的值，统计正数和负数的个数，输入0结束

	var positiveCount int
	var negetiveCount int
	var num int

	for {
		fmt.Println("请输入数值")
		fmt.Scanln(&num)
		if num==0{
			break
		}
		if num>0 {
			positiveCount++
			continue
		}
		negetiveCount++
	}
	fmt.Printf("正数的个数是 %v 负数的个数是 %v",positiveCount,negetiveCount)

}