package main

import (
	"fmt"
)

func main () {

	key := ""
	loop := true

	balance := 1000.0
	money :=0.0
	note := ""
	details := "收支\t账户余额\t收支金额\t说    明"

	for {
		fmt.Println("-----------家庭收支记账软件-----------")
		fmt.Println("            1 收支明细")
		fmt.Println("            2 登记收入")
		fmt.Println("            3 登记支出")
		fmt.Println("            4 退出软件")
		fmt.Print("请选择（1-4）：")
		fmt.Scanln(&key)

		switch key {
		case"1":
			fmt.Println("-----------当前收支记账明细-----------")
			fmt.Println(details)
		case "2":
			fmt.Println("本次收入金额")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v",balance,money,note)
		case "3":
			fmt.Println("登记支出")
		case "4":
			loop = false
		default:
			fmt.Println("请输入正确的选项。。。")
		}
		if !loop {
			break
		}
	}
	fmt.Println("你退出家庭记账软件")
}