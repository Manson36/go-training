package main

import (
	"errors"
	"fmt"
	"os"
)

//使用一个结构体管理队列
type Queue struct {
	maxSize int
	array [5]int //数组-》模拟队列
	front int //表示队列首
	rear int //表示队列尾
}

//添加数据到队列
func (this *Queue) AddQueue(val int) (err error) {
	//先判断队列是否已满
	if this.rear == this.maxSize - 1 {
		return errors.New("queue full")
	}

	this.rear++
	this.array[this.rear] = val

	return nil
}

//从队列中取出数据
func (this *Queue) GetQueue() (val int, err error) {
	//先判断队列是否为空
	if this.front == this.rear {
		return -1, errors.New("queue empty")
	}

	this.front++
	val = this.array[this.front]
	return val, nil
}

//显示队列：从队首遍历到队尾
func (this *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是：")

	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]= %d\t", i, this.array[i])
	}
	fmt.Println()
}

//编写一个主函数，测试
func main() {
	//先创建一个队列
	queue := &Queue{
		maxSize: 5,
		front: -1,
		rear: -1,
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
			err := queue.AddQueue(val)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("加入队列OK")
			}
		case "get":
			val, err := queue.GetQueue()

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("从队列中获取一个数：", val)
			}
		case "show" :
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
