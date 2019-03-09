package main

import "fmt"

type Account struct {
	Accountno string
	Pwd       string
	Balance   float64
}
//方法
//1.存款
func (account *Account) Deposite(money float64, pwd string) {
	//检查密码是否正确
	if pwd != account.Pwd{
		fmt.Println("你输入的密码不正确")
		return
	}
	//看看存款金额是否正确
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}
	account.Balance += money
	fmt.Println("存款成功")
}
//取款
func (account *Account)WithDrow(money float64,pwd string )  {
	//看下输入的密码是否正确
	if pwd != account.Pwd{
		fmt.Println("你输入的密码不正确")
		return
	}
	//看看存款金额是否正确
	if money <= 0 ||money> account.Balance{
		fmt.Println("你输入的金额不正确")
		return
	}
	account.Balance -= money
	fmt.Println("取款成功")
}
//查询余额
func (account *Account)Query(pwd string ) {
	//看下输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	fmt.Printf("你的账号为 %v 余额为 %v",account.Accountno,account.Balance)
}


func main() {

	account := Account{
		Accountno:"gs111111",
		Pwd:"666666",
		Balance:100.0,
	}
	account.Query("666666")
	account.Deposite(200.0,"666666")
	account.Query("666666")
	account.WithDrow(150.0,"666666")
	account.Query("666666")
}