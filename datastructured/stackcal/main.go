package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {
	MaxTop int
	Top int
	arr [20]int
}

//入栈
func (this *Stack) Push(val int) (err error) {
	if this.Top == this.MaxTop -1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}

	this.Top++
	this.arr[this.Top] = val
	return
}

//出栈
func (this *Stack) Pop() (val int, err error) {
	if this.Top == -1 {
		fmt.Println("stack empty")
		return -1, errors.New("stack empty")
	}

	val = this.arr[this.Top]
	this.Top--

	return val, nil
}

//遍历栈
func (this *Stack) List() {
	if this.Top == -1 {
		fmt.Println("stack empty")
		return
	}

	fmt.Println("栈的情况如下")
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]:%d", i, this.arr[i])
	}
}

//判断字符是不是运算符（+-*/)
func (this *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

//运算的方法
func (this *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
		case 42:
			res = num2 * num1
		case 43:
			res = num2 + num1
		case 45:
			res = num2 - num1
		case 47:
			res = num2 / num1
		default:
			fmt.Println("运算符错误")
	}

	return res
}

//编写一个方法，返回某个运算符的优先级【程序员第一
// * / -> 1; + - -> 0
func (this *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}

	return res
}

func main() {
	//数栈
	numStack := &Stack{
		MaxTop: 20,
		Top: -1,
	}

	//符号栈
	operStack := &Stack{
		MaxTop: 20,
		Top: -1,
	}

	exp := "30+30*6-4-6"
	//定义一个index，帮助扫描exp
	index := 0

	//为了配合运算，我们需要定义一些变量
	num1 := 0
	num2 := 0
	result := 0
	oper := 0
	keepNum := "" //处理多位数

	for {
		ch := exp[index:index+1]//字符串
		temp := int([]byte(ch)[0])//就是字符对应的ASCII的值

		if operStack.IsOper(temp) {//说明是符号
			if operStack.Top == -1 {//说明是空栈，直接放入
				operStack.Push(temp)
			}else {//如果发现栈顶的运算符优先级大于等于当前准备入栈的运算符的优先级，就从符号栈pop，数栈pop两个数，
				//进行运算，运算后的结果再重新入栈到数栈，运算符加入到符号栈
				if operStack.Priority(operStack.arr[operStack.Top]) >=
					operStack.Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					result = operStack.Cal(num1, num2, oper)

					numStack.Push(result)
					operStack.Push(temp)
				}else {
					operStack.Push(temp)
				}
			}
		} else {//说明是数
			//处理多位数思路：
			//1.使用字符串拼接
			keepNum += ch

			//2.每次向index后面一位测试是不是符号或者达到最后
			if index == len(exp) - 1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val))
			} else {
				if operStack.IsOper(int([]byte(exp[index+1: index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = ""//很关键，清空keepNum
				}
			}
			//val, _ := strconv.ParseInt(ch, 10, 64)
			//numStack.Push(int(val))//直接传入的temp是ASCII值，需转换为本来的数
		}

		//继续扫描，判断index是否到最后
		if index + 1 ==len(exp) {
			break
		}
		index++
	}

	//表达式扫描完毕，依次从符号栈和数栈中取出，计算
	for {
		if operStack.Top == -1 {
			break
		}

		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = operStack.Cal(num1, num2, oper)

		numStack.Push(result)
	}

	//如果我们的算法没有问题，表达式也正确，结果就是numStack最后的数
	res, _ := numStack.Pop()
	fmt.Printf("表达式%s = %v", exp, res)
}
