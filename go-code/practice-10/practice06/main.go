package main

import "fmt"

type Calcuator struct{
	n1 float64
	n2 float64
}

func (cal *Calcuator)getSum() float64{
	return cal.n1+cal.n2
}

func (cal *Calcuator)getSub() float64{
	return cal.n1-cal.n2
}

func main() {

//实现小小计算器
	var cal Calcuator
	cal.n1 = 12
	cal.n2 = 33
	fmt.Println(cal.getSum())
	fmt.Println(cal.getSub())

}