package main

import "fmt"

func main(){

	//switch 把小写类型的char类型转换为大写
	var char byte
	fmt.Println("请输入一个字符")
	fmt.Scanf("%c",&char)

	switch char{
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	case 'c':
		fmt.Println("C")
	case 'd':
		fmt.Println("D")
	default:
		fmt.Println("other")

	}

	//输入成绩 大于60为合格，小于60不合格，不能大于100
	var score int

	fmt.Println("请输入成绩")
	fmt.Scanln(&score)

	switch int(score/60){
	case 1:
		fmt.Println("及格")
	case 0:
		fmt.Println("不及格")
	default:
		fmt.Println("输入有误")
	}
}