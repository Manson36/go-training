package main

import (
	"errors"
	"fmt"
	"os"
)

type singleQueue struct {
	maxNum int
	array  [5]int
	front  int
	rear   int
}

func (this *singleQueue) AddQueue(val int) (err error) {
	if this.rear == this.maxNum - 1 {
		return errors.New("queue is full")
	}
	this.rear++
	this.array[this.rear] = val
	return nil
}

//从队列中取出
func (this *singleQueue) GetQueue() (val int, err error) {
	if this.front == this.rear {
		fmt.Println("queue is empty")
		return -1, errors.New("queue is empty")
	}

	this.front++
	val = this.array[this.front]
	return val, nil
}

func (this *singleQueue) ShowQueue() {
	if this.front == this.rear {
		fmt.Println("queue is empty")
	}
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("%d\t", this.array[i])
		fmt.Println()
	}
}

func main() {
	singleQueue := &singleQueue{
		maxNum: 5,
		front: -1,
		rear: -1,
	}

	var key string
	var val int
	for {
		fmt.Println("1.输入add表示往队列中添加数据")
		fmt.Println("2.输入get表示往队列中获取数据")
		fmt.Println("3.输入show表示往队列中显示数据")
		fmt.Println("4.输入exit表示退出")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要添加的数据")
			fmt.Scanln(&val)
			err := singleQueue.AddQueue(val)
			if err != nil {
				fmt.Println(err)
			}else {
				fmt.Println("add ok ")
			}
		case "get":
			val, err := singleQueue.GetQueue()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("从队列中获取到数据:%d\n", val)
			}
		case "show":
			singleQueue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}

}
