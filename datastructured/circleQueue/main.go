package main

import (
	"errors"
	"fmt"
	"os"
)

//使用一个结构体管理环形队列
type CircleQueue struct {
	maxSize int //4
	array [5]int
	head int
	tail int
}

//入队列
func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("queue full")
	}

	//分析出this.tail在队列的尾部，但是不包含最后的元素
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return
}

//出队列
func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("queue empty")
	}

	//取出。head是指向队首，并且包含队首元素
	val = this.array[this.head]
	this.head= (this.head + 1) % this.maxSize
	return
}

//显示队列
func (this *CircleQueue) ListQueue() {
	if this.Size() == 0 {
		fmt.Println("queue id empty")
	}

	//设计一个辅助变量，指向head
	tempHead := this.head
	for i := 0; i < this.Size(); i++ {
		fmt.Printf("array[%d]=%d\t", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()
}

//判断队列是否为满
func (this *CircleQueue) IsFull() bool {
	return (this.tail + 1) % this.maxSize == this.head
}

//判断队列是否为空
func (this *CircleQueue) IsEmpty() bool {
	return this.head == this.tail
}

//统计队列中数据的数量
func (this *CircleQueue) Size() int {
	return (this.tail - this.head + this.maxSize) % this.maxSize
}

//编写一个主函数，测试
func main() {
	//先创建一个队列
	queue := &CircleQueue{
		maxSize: 5,
		head: 0,
		tail: 0,
	}

	var key string
	var val int

	for {
		fmt.Println("1.输入 add 表示添加数据到队列")
		fmt.Println("2.输入 get 表示从队列中获取数据")
		fmt.Println("3.输入 show 表示显示队列数据")
		fmt.Println("4.输入 exit 表示退出队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要添加的数据")
			fmt.Scanln(&val)
			err := queue.Push(val)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("加入队列OK")
			}
		case "get":
			val, err := queue.Pop()

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("从队列中获取一个数：", val)
			}
		case "show" :
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
