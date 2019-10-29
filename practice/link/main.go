package main

import "fmt"

type heroNode struct {
	no int
	name string
	nickName string
	next *heroNode
}

//在最尾加【简单】
func (this *heroNode) insert(head *heroNode, newHero *heroNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	temp.next = newHero
}

//根据heronode的no插入
func (this *heroNode) insert2(head *heroNode, newHero *heroNode) {
	temp := head
	for {
		if temp.next == nil {
			temp.next = newHero
			break
		}
		if temp.next.no > newHero.no {
			newHero.next = temp.next
			temp.next = newHero
			break
		}
		if temp.no == newHero.no {
			fmt.Println("Hero 已经存在")
			return
		}

		temp = temp.next
	}
}

func main() {
	
}
