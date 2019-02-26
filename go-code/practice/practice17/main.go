package main

import "fmt"

func main(){

	var score int
	fmt.Println("请输入分数")
	fmt.Scanln(&score)

	if score==100 {
		fmt.Println("奖励一台bwv")
	}else if score>80&& score<=99{
				fmt.Println("奖励一台iPhone7 Plus")
			}else if score>=60&& score<=80{
				fmt.Println("奖励一台iPad")
		}else{
			fmt.Println("什么都没有")
		}
}