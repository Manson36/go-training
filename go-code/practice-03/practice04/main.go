package main

import "fmt"

func peach(n int) int {

	if n< 1|| n> 10{
		fmt.Println("输入天数不对")
		return 0
	}
	if n== 10{
		return 1
	}else {
		return (peach(n+1)+1)*2
	}
}

func main(){

	//有一堆桃子，猴子第一天吃了一半，外加一个，第二天再吃一半，外加一个，第十天在吃的时候
	//发现只剩下一个桃子，问这堆桃子有多少个？

	fmt.Println("第一天的桃子数是",peach(1))

}