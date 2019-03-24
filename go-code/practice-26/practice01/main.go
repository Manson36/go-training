package main

import (
	"fmt"
	"github.com/go-training/go-code/practice-26/practice03"
)

type customerView struct {
	key string
	loop bool //是否循环显示主菜单
	//增加一个字段
	customerService *service.CustomerService
}

//显示所有的客户信息
func (this *customerView)list() {
	//首先，获取到当前所有的客户信息
	customers :=this.customerService.List()
	//显示
	fmt.Println("---------客户列表---------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n\n--------客户列表完成--------\n")
}

func (this *customerView)mainMenu() {
	for {

		fmt.Println("--------客户信息管理软件--------")
		fmt.Println("        1 添加客户")
		fmt.Println("        2 修改客户")
		fmt.Println("        3 删除客户")
		fmt.Println("        4 客户列表")
		fmt.Println("        5 退    出")
		fmt.Print("请选择1-5：")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			fmt.Println("添加客户")
		case "2":
			fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")
		case "4":
			fmt.Println("客户列表")
			this.list()
		case "5":
			fmt.Println("退    出")
			this.loop = false
		default:
			fmt.Println("您的输入有误，请重新输入")
		}
		if !this.loop {
			break
		}

	}
	fmt.Println("您退出了客户管理系统")
}

func main() {

	//创建一个customerView，并运行显示主菜单
	customerView := customerView{
		key : "",
		loop:true,

	}

	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}