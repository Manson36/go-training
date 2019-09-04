package main

import "fmt"

type HeroNode struct {
	no int
	name string
	nickName string
	next *HeroNode
}

//给链表插入一个节点

//方法一：在链表的最后加入【简单】
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	//思路：1.先找到链表最后那个节点，2.创建一个辅助结点，跑龙套
	temp := head

	for {
		if temp.next == nil {//表示找到最后
			break
		}

		temp = temp.next //让temp不断指向下一个节点
	}

	//3.将newHeroNode添加到链表最后
	temp.next = newHeroNode
}

//第二种方法，根据no大小进行排序【实用】
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	//思路：1.先找到链表最后那个节点，2.创建一个辅助结点，跑龙套
	temp := head
	flag := true

	for {
		if temp.next == nil {//说明到链表的最后
			break
		} else if temp.next.no > newHeroNode.no {
			//说明newHeroNode就应该插入到temp只后
			break
		} else if temp.next.no == newHeroNode.no {
			//说明我们链表中已经有这个no，就不能插入
			flag = false
			break
		}

		temp = temp.next //让temp不断指向下一个节点
	}

	if !flag {
		fmt.Println("对不起，已经存在 no=", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}

	//3.将newHeroNode添加到链表最后
	temp.next = newHeroNode
}

func ListHeroNode(head *HeroNode) {
	//1.创建一个辅助结点
	temp := head

	//2.先判断这个链表是不是一个空链表
	if temp.next == nil {
		fmt.Println("link empty")
		return
	}

	for {
		temp = temp.next
		fmt.Printf("[%d, %s, %s]->", temp.no, temp.name, temp.nickName)
		//判断是否到最后
		if temp.next == nil {
			break
		}
	}
}

func DelHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false

	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next
	}

	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("sorry id不存在")
	}
}

func main() {
	//1.创建一个头结点
	head := &HeroNode{}

	//2.创建新的链表结点
	hero1 := &HeroNode{
		no : 1,
		name : "宋江",
		nickName: "及时雨",
	}
	hero2 := &HeroNode{
		no : 2,
		name : "卢俊义",
		nickName: "玉麒麟",
	}
	hero3 := &HeroNode{
		no : 3,
		name : "林冲",
		nickName: "豹子头",
	}
	hero4 := &HeroNode{
		no : 4,
		name : "吴用",
		nickName: "智多星",
	}

	//InsertHeroNode(head, hero1)
	//InsertHeroNode(head, hero2)
	//InsertHeroNode(head, hero3)
	//InsertHeroNode(head, hero4)

	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero2)
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero4)

	ListHeroNode(head)
}
