package main

import (
	"errors"
	"fmt"
	"os"
)

type CircleQueue struct {
	maxSize int
	array [5]int
	head int
	tail int
}

func (this *CircleQueue) addQueue(val int) (err error) {
	if (this.tail + 1) % this.maxSize == this.head {
		fmt.Println("Queue is full")
		return errors.New("Full")
	}

	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return nil
}

func (this *CircleQueue) Pop() (val int, err error) {
	if this.head == this.tail {
		fmt.Println("queue is empty")
		return -1, errors.New("empty")
	}

	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return
}

func (this *CircleQueue) showQueue() {
	if this.head == this.tail {
		fmt.Println("Queue is empty")
		return
	}

	//创建辅助变量，指向head
	tempHead := this.head
	for i := 0; i < this.size(); i++ {
		fmt.Printf("array[%d]:%d\n", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
}

//计算队列的长度
func (this *CircleQueue) size() int {
	return (this.tail - this.head + this.maxSize) % this.maxSize
}

func main() {
	circleQueue := CircleQueue{
		maxSize: 5,
		head: 0,
		tail: 0,
	}

	var key string
	var val int
	for {
		fmt.Println("输入add表示添加数据")
		fmt.Println("输入pop表示取出数据")
		fmt.Println("输入show表示显示数据")
		fmt.Println("输入exit表示退出")
		fmt.Scanln(&key)

		switch key {
		case "add" :
			fmt.Println("输入要添加的数据")
			fmt.Scanln(&val)

			err := circleQueue.addQueue(val)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("添加完成")
			}
		case "pop":
			val, err := circleQueue.Pop()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("取到数据：", val)
			}
		case "show" :
			circleQueue.showQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
