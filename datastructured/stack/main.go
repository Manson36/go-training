package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int
	Top int
	arr [5]int
}

//入栈
func (this *Stack) Push(val int) (err error) {
	if this.Top == this.MaxTop - 1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}

	this.Top++
	this.arr[this.Top] = val
	return nil
}

//出栈
func (this *Stack) Pop() (val int, err error) {
	//先判断是否为空
	if this.Top == -1 {
		fmt.Println("Stack empty")
		return 0, errors.New("stack empty")
	}

	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

//展示
func (this *Stack) List() {
	if this.Top == -1 {
		fmt.Println("stack empty")
		return
	}

	fmt.Println("栈的情况如下")
	for i:= this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

func main() {
	var stack = &Stack{
		MaxTop: 5,
		Top:-1,
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)

	stack.List()
	val, _ := stack.Pop()
	fmt.Println("Pop val=", val)
	stack.List()
}
