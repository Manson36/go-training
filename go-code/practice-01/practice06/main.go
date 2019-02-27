package main

import "fmt"

func main(){

	//awitch 后也可以不带表达式
	var age int =10

	switch {
	case age== 10:
		fmt.Println("age == 10")
	case age == 20:
		fmt.Println("age == 20")
	default:
		fmt.Println("没有匹配到")
	}

	var score int =90

	switch {
	case score >=90:
		fmt.Println("成绩优秀")
	case score>=70&&score<90:
		fmt.Println("成绩良好")
	case score >=60&& score <70:
		fmt.Println("成绩及格")
	default:
		fmt.Println("不及格")
	}
}