package utils

import (
	"fmt"
)

type FamilyAccount struct {
	key string
	loop bool
	balance float64
	money float64
	note string
	details string
	flag bool
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:"",
		loop:true,
		balance:0.0,
		money:0.0,
		note:"",
		details:"收支\t,账户金额\t收支金额\t说明",
		flag:false,

	}
}
func(this *FamilyAccount)showDetails() {
	if this.flag{
		fmt.Println(this.details)
	}else {
		fmt.Println("当前没有记录，来一笔把")
	}
}

func(this *FamilyAccount)income() {
	fmt.Println("本次收入金额")
	fmt.Scanln(&this.money)
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	this.balance += this.money
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v",this.balance,this.money,this.note)
	this.flag = true
}
func(this *FamilyAccount)pays() {
	fmt.Println("本次支出金额")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足")
	}else {
		fmt.Println("本次支出原因")
		fmt.Scanln(&this.note)
		this.balance -=this.money
		this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v",this.balance,this.money,this.note)
		this.flag =true
	}
}
func(this *FamilyAccount)exit() {
	fmt.Println("你确定要退出吗？y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice  == "y" || choice  == "n" {
			break
		}else {
			fmt.Println("你的输入有误，请重新输入")
		}
	}
	if choice == "y"{
		this.loop = false
	}
}

func (this *FamilyAccount)MainMenu() {
	for {
		fmt.Println("-----------家庭收支记账软件-----------")
		fmt.Println("            1 收支明细")
		fmt.Println("            2 登记收入")
		fmt.Println("            3 登记支出")
		fmt.Println("            4 退出软件")
		fmt.Print("请选择（1-4）：")
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pays()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项")
		}
		if !this.loop {
			break
		}
	}
}