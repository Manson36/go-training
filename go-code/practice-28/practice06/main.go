package main

import "fmt"

type Calcuator struct {
	Num1 float64
	Num2 float64
}

func (calcuator *Calcuator)getSum()float64 {
	return calcuator.Num1 + calcuator.Num2
}

func (calcuator *Calcuator)getSub()float64 {
	return calcuator.Num1 - calcuator.Num2
}
func (calcuator *Calcuator)get3()float64 {
	return calcuator.Num1 * calcuator.Num2
}
func (calcuator *Calcuator)get4()float64 {
	return calcuator.Num1 / calcuator.Num2
}

//实现形式2
func (calcuator *Calcuator)getRes(operator byte)float64 {

	res := 0.0
	switch operator {
	case '+':
		res = calcuator.Num1 + calcuator.Num2
	case '-':
		res = calcuator.Num1 + calcuator.Num2
	case '*':
		res = calcuator.Num1 + calcuator.Num2
	case '/':
		res = calcuator.Num1 + calcuator.Num2
	default:
		fmt.Println("您的输入有误")
	}
	return res
}

func main() {


}