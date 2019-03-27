package main

import "fmt"

type Account struct {
	AccountNo string
	Pwd string
	Banlance float64
}

//存款
func (account *Account)Deposite(money float64,pwd string)  {

	if pwd != account.Pwd {
		fmt.Println("你输入的密码错误")
		return
	}
	if money < 0 {
		fmt.Println("你输入的金额不正确")
		return
	}
	account.Banlance += money

	fmt.Println("存款成功")
}
//取款
func (account *Account)WithDraw(money float64,pwd string)  {

	if pwd != account.Pwd {
		fmt.Println("你输入的密码错误")
		return
	}
	if money < 0 {
		fmt.Println("你输入的金额不正确")
		return
	}
	account.Banlance -= money

	fmt.Println("取款成功")
}

//查询金额
func (account *Account)Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("你输入的密码错误")
		return
	}
	fmt.Printf("你的账号为%v，余额为%v\n",account.AccountNo,account.Banlance)
}

func main() {

	account := Account{
		AccountNo:"gs111111",
		Pwd:"222",
		Banlance:100.0,
	}
	account.Query("222")
	account.Deposite(111.0,"222")
	account.Query("222")
	account.WithDraw(11,"222")
	account.Query("222")
}