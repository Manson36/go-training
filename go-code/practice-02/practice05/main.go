package main

import "fmt"

func main(){

	//实现登录有三次验证，如果用户名为张无忌，密码为888，显示通过
	//不通过，显示还有几次机会
	var name string
	var pwd string
	var loginChance =3

	for i:=1;i<=3;i++{
		fmt.Println("请输入用户名")
		fmt.Scanln(&name)
		fmt.Println("请输入密码")
		fmt.Scanln(&pwd)

		if name=="张无忌"&&pwd=="888"{
			fmt.Println("恭喜你登录成功")
			break
		}else{
			loginChance--
			fmt.Printf("密码错误，你还有%v 次机会，请珍惜",loginChance)
		}
		if loginChance==0{
			fmt.Println("机会用完，登录失败！")
		}
	}
}