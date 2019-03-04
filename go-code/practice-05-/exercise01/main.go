package main

import "fmt"

//循环打印输入的月份的天数。【使用continue实现】
//要有判断输入的月份是否错误的语句
func main(){

	var year int
	var Month int

	for {
		fmt.Println("请输入月份")
		fmt.Scanln(&Month)

		if Month>0&&Month<=12{
			if Month ==4 ||Month ==6 ||Month ==9 ||Month ==11{
				fmt.Println("该月份有30天")
			}else if Month==2{
				fmt.Println("请输入年份")
				fmt.Scanln(&year)
				if (year%4==0&&year%100!=0)||year%400==0{
					fmt.Println("该月份有29天")
				}else {
					fmt.Println("该月份有28天")
				}
			}else {
				fmt.Println("该月份有31天")
			}
		}else {
			fmt.Println("输入的月份错误")
			continue
		}
	}
}