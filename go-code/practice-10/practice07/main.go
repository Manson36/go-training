package main

import "fmt"

type Calcuator struct{
	n1 float64
	n2 float64
}

func (cal Calcuator)getRes(operator byte)float64{
	res := 0.0

	switch operator {
	case '+':
		res = cal.n1 + cal.n2
	case '-':
		res = cal.n1 - cal.n2
	case '*':
		res = cal.n1 * cal.n2
	case '/':
		res = cal.n1 / cal.n2
	default:
		fmt.Println("运算符输入有误")

	}
	return res
}

func main() {

	var cal Calcuator
	cal.n1 = 12
	cal.n2 = 33
	fmt.Println(cal.getRes('*'))
}