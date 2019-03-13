package main

import (
	"fmt"
	"github.com/go-training/go-code/practice-15/practice02/model"
	"github.com/go-training/go-code/practice-15/practice02/service"
)

type customerView struct {
	key string
	loop bool
	customerService *service.CustomerService
}

func (this *customerView)list() {
	customers := this.customerService.List()
	//显示
	fmt.Println("-----------客户列表-----------")
	fmt.Println("编号\t姓名\t年龄\t电话\t邮箱")
	for i:= 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("\n----------客户列表完成----------\n\n")
}

func (this *customerView)mainMenu(){

	for {
		fmt.Println("----------客户信息管理软件----------")
		fmt.Println("           1.添 加 客 户")
		fmt.Println("           2.修 改 客 户")
		fmt.Println("           3.删 除 客 户")
		fmt.Println("           4.客 户 列 表")
		fmt.Println("           5.退      出")
		fmt.Print("请选择（1-5）")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.add()
		case "2":
			fmt.Println("修 改 客 户")
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
		if !this.loop {
			break
		}

	}
	fmt.Println("你退出了客户管理系统")
}

func main()  {

	customerView := customerView{
		loop:true,
		key:"",
	}
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}
func (this *customerView)add() {
	fmt.Println("----------添加客户----------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age )
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮：")
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCustomer2(name,gender,age,phone,email)
	//调用
	if this.customerService.Add(customer){
	fmt.Println("----------添加完成----------")
	}else {
		fmt.Println("----------添加失败----------")
	}
}

func (this *customerView)delete() {
	fmt.Println("----------删除客户----------")
	fmt.Println("请选择删除的客户编号（—1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("是否确认删除（Y/N）")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		if this.customerService.Delete(id) {
			fmt.Println("----------删除完成----------")
		}else {
			fmt.Println("----------删除失败，输入的ID不存在-----")
		}
	}
}

func (this *customerView)exit() {
	fmt.Println("请确认是否退出（Y/N）")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" ||this.key == "y" ||this.key == "n" ||this.key == "N" {
			break
		}
		fmt.Println("你的输入有误,确认是否退出（Y/N)")
	}
	if this.key == "Y" ||this.key == "y" {
		this.loop = false
	}
}