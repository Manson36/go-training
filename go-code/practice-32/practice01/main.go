package main

import (
	"errors"
	"fmt"
	"os"
)

//数组模拟环形队列

//使用一个结构体管理环形队列
type CircleQueue struct {
	maxSize int
	array [5]int
	head int
	tail int
}

//入队列
func (this *CircleQueue)Push(val int )(err error) {

	if this.IsFull() {
		return errors.New("queue full")
	}
	this.array[this.tail] = val
	this.tail = (this.tail + 1)% this.maxSize
	return
}

//出队列

func (this *CircleQueue)Pop() (val int, err error) {

	if this.IsEmpty() {
		return 0,errors.New("queue empty")
	}
	val = this.array[this.head]
	this.head = (this.head + 1)% this.maxSize
	return
}

//显示队列
func (this *CircleQueue)ListQueue() {

	fmt.Println("环形队列的情况如下")
	size := this.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}
	//设置一个辅助的变量，指向head
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t",tempHead,this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()
}

//判断环形队列是否已满
func (this *CircleQueue)IsFull()bool {
	return (this.tail + 1)%this.maxSize == this.head
}

//判断环形队列是否为空
func (this *CircleQueue)IsEmpty()bool {
	return this.head == this.tail
}

//计算环形队列有多少个元素
func (this *CircleQueue)Size() int {
	//这是一个关键的算法
	return (this.tail +  this.maxSize - this.head)% this.maxSize
}

func main() {

	//初始化一个环形队列
	queue := &CircleQueue{
		maxSize:5,
		head:0,
		tail:0,
	}
	var key string
	var val int
	var err error

	for {
		fmt.Println("1.输入add表示添加数据到队列")
		fmt.Println("2.输入get表示添从队列获取数据")
		fmt.Println("3.输入show表示显示队列")
		fmt.Println("4.输入exit表示退出队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要添加的数")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列成功")
			}
		case "get":
			val, err = queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出一个数", val)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("您的输入有误，请重新输入")
		}
	}
}