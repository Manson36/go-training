package utils

import (
	"fmt"
)

type FamilyAccount struct {
	//保存接受用户输入的选项
	key string
	//控制是否退出for
	loop bool
	//账户余额
	balance float64
	//每次收支金额
	money float64
	//收支说明
	note string
	//是否有收支行为
	flag bool
	//收支记录
	details string

}
//编写一个工厂模式，返回实例
func NewFamilyAccount () *FamilyAccount{
	return &FamilyAccount{
		key:"",
		loop:true,
		balance:10000.0,
		money :0.0,
		note:"",
		flag:false,
		details:"收支\t账户金额\t收支金额\t说明",

	}
}
//将显示明细写成一个方法
func (this *FamilyAccount)showDetails(){
	if this.flag {
		fmt.Println(this.details)
	}else {
			fmt.Println("当前没有收支，来一笔把")
	}
}
//将登记收入写成一个方法
func (this *FamilyAccount)income() {

	fmt.Println("本次收入金额为")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v",this.balance,this.money,this.note)
	this.flag = true
}
//将登记支出写成一个方法
func (this *FamilyAccount)pay () {
	fmt.Println("本次支出金额为")
	fmt.Scanln(&this.money)
	if this.money > this.balance{
		fmt.Println("您的余额不足")
	}else {
		this.balance -= this.money
		fmt.Println("本次支出说明")
		fmt.Scanln(&this.note)
		this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
		this.flag = true
	}
}

//将退出系统写成一个方法
func (this *FamilyAccount)exit() {

	fmt.Println("你确定要退出吗？")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice=="y" || choice == "n" {
			break
		}
		fmt.Println("您的输入有误，请重新输入")
		if choice == "y" {
			this.loop = true
		}
	}
}
func (this *FamilyAccount)MainMenu() {

	for {
		fmt.Println("----------家庭收支记账软件----------")
		fmt.Println("            1.收支明细")
		fmt.Println("            2.登记收入")
		fmt.Println("            3.登记支出")
		fmt.Println("            4.退出软件")
		fmt.Print("请选择（1-4）")
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()

		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项")
		}
		if !this.loop{
			break
		}
	}
}
