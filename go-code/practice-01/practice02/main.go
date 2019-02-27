package main

import "fmt"

func main(){

	var month byte
	var age byte
	var price int = 60

	fmt.Println("请输入游玩月份")
	fmt.Scanln(&month)
	fmt.Println("请输入游客年龄")
	fmt.Scanln(&age)

	if month >4&& month<10{
		if age > 60{
			fmt.Printf("%v 月 票价%v 年龄%v", month,price/3,age)
		}else if age>=18{
			fmt.Printf("%v 月 票价%v 年龄%v", month,price,age)
		}else{
			fmt.Printf("%v 月 票价%v 年龄%v", month,price/2,age)
		}
	}else {
		if age>=18 && age<=60 {
			fmt.Println("淡季成人票价 40")
		}else {
			fmt.Println("淡季儿童和老人的票价 20")
		}
	}
}