package main

import (
	"fmt"
	"math/rand"
	"time"
)

//编写一个函数，随机猜数游戏；随机生成一个1-100的整数，有十次机会，
	//   如果第一次就猜中，提示“你真是一个天才”
	//  如果第2-3次猜中，提示“你很聪明，快赶上我了”
	//  如果第4-9次猜中，提示“一般般”
	//  如果最后一次猜中，提示“可算猜对了”
	//  以此都没猜中，提示“说你点啥好呢”

func main(){

	rand.Seed(time.Now().Unix())
	num := rand.Intn(100)+1
	var num2 int

	for i:=0; i<10; i++{
		fmt.Println("请输入你猜的数")
		fmt.Scanln(&num2)

		if num==num2{
			if i==0{
				fmt.Println("你真是一个天才")
			}else if i==1||i==2 {
				fmt.Println("你很聪明，快赶上我了")
			}else if i==9{
				fmt.Println("可算猜对了")
			}else{
				fmt.Println("一般般")
			}
		}
		}
	fmt.Println("说你点啥好呢？")
	}