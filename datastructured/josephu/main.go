package main

import (
	"fmt"
)

//小孩的结构体
type Boy struct {
	No int
	Next *Boy
}

//编写一个函数，构建环形单向链表，num表小孩个数，*boy表返回第一个小孩
func AddBoy(num int) *Boy {
	first := &Boy{}
	curBoy := &Boy{}

	if num < 1 {
		fmt.Println("num的值不对")
		return first
	}

	//循环构建环形列表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}

		//因为第一个小孩比较特殊
		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.Next = first
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first
		}
	}

	return first
}

//显示单向环形链表
func ShowBoy(first *Boy) {
	//处理列表为空
	if first.Next == nil {
		fmt.Println("列表为空")
		return
	}

	//创建一个指针，帮助遍历
	curBoy := first
	for {
		//do-while 的形式
		fmt.Printf("小孩id:%d->", curBoy.No)

		if curBoy.Next == first{
			break
		}

		curBoy = curBoy.Next
	}
}

//编号1-n个人围坐一圈，第k个人（1<=k<=n)从1开始报数，数到m的人出列，他的下一位开始从1报数，m的人出列，直到所有人出列
func PlayGame(first *Boy, startNo int, countNo int) {
	//1.空链表我们单独处理
	if first.Next == nil {
		fmt.Println("空的链表")
		return
	}

	//留一个， 判断startNo <= 小孩的总数
	//2.需要定义辅助指针，帮助我们删除，指向链表最后一个小孩
	tail := first
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}

	//3.将first移动到starNo位置，tail随其后【后面我们删除的小孩，就是first位置】
	for i := 1; i <= startNo - 1; i++ {
		first = first.Next
		tail = tail.Next
	}
	fmt.Println()

	//4.开始数countNo，删除first指向的小孩
	for {
		for i := 1; i <= countNo -1; i++ {
			first = first.Next
			tail = tail.Next
		}

		fmt.Printf("编号为%d的小孩出圈\n", first.No)
		//执行删除
		first = first.Next
		tail.Next = first
		//圈里只有一个小孩结束
		if tail == first {
			break
		}
	}
	fmt.Printf("最后编号为%d的小孩出圈\n", first.No)
}

func main() {
	first := AddBoy(50)
	ShowBoy(first)
	PlayGame(first, 10, 20)
}
