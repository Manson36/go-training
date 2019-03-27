package model

import "fmt"

type account struct {
	accountNo string
	pwd string
	balance float64

}

//工厂模式
func NewAccount(accountNo string,pwd string,balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号的长度不对")
		return nil
	}

	if len(pwd) != 6 {
		fmt.Println("密码的长度不对")
		return nil
	}

	if balance < 20 {
		fmt.Println("余额的数量不对")
		return nil
	}

	return &account{
		accountNo:accountNo,
		pwd:pwd,
		balance:balance,
	}
}

//方法
//存款
func (account *account)Deposite(money float64,pwd string) {
	if pwd != account.pwd {
		fmt.Println("您输入的密码有误")
		return
	}

	if money <= 0 {
		fmt.Println("您输入的金额有误")
		return
	}

	account.balance += money
	fmt.Println("存款成功")
}

//取款
func (account *account)WithDraw(money float64,pwd string) {
	if pwd != account.pwd {
		fmt.Println("您输入的密码有误")
		return
	}

	if money <= 0 || money > account.balance {
		fmt.Println("您输入的金额有误")
		return
	}

	account.balance -= money
	fmt.Println("取款成功")
}

//余额查询
func (account *account)Query(pwd string) {
	fmt.Printf("你的账号为%v，余额为%v\n",account.accountNo,account.balance)
}