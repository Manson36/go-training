package main

import "fmt"

func main(){

	//函数前 请输入两个数 在输入运算符 求结果

	var n1 float64 =1.2
	var n2 float64 =2.3
	var operator byte ='-'
	var res float64

	switch operator {
	case '+':
		res= n1+n2
	case '-':
		res = n1-n2
	case '*':
		res = n1*n2
	case '/':
		res = n1/n2
	default:
		fmt.Println("输入符号错误")

	}
	fmt.Println("res=",res)
}