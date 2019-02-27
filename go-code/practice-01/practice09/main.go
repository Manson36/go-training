package main

import "fmt"

func main(){

	//打印月份所属季节
	var month byte

	fmt.Println("请输入月份")
	fmt.Scanln(&month)

	switch month{
	case 3,4,5:
		fmt.Println("Spring")
	case 6,7,8:
		fmt.Println("summer")
	case 9,10,11:
		fmt.Println("Autumn")
	case 12,1,2:
		fmt.Println("Winter")
	default:
		fmt.Println("输入有误")
	}
}