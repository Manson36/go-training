package main

import "fmt"

func Cal(n1 float64, n2 float64,operator byte)float64 {

	var res  float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("输入符号错误")

	}
	return res
}

func main(){

	var n1 float64=1.2
	var n2 float64=2.3
	var operator byte ='+'
	result := Cal(n1,n2,operator)
	fmt.Println("result=",result)
}