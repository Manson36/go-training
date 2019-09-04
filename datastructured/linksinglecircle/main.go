package main

import "fmt"

type CatNode struct {
	no int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newNode *CatNode) {
	//首先，判断是不是添加的第一只猫
	if head.next == nil {
		head.no = newNode.no
		head.name = newNode.name
		head.next = head //自身构成一个环

		fmt.Println(newNode, "加入到环形列表")
		return
	}

	//在最后加入节点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}

	temp.next = newNode
	newNode.next = head
	fmt.Println(newNode, "加入链表")
}

//输出环形列表
func ListCatNode(head *CatNode) {
	fmt.Println("环形列表的信息如下：")
	temp := head
	if temp.next == nil {
		fmt.Println("empty")
		return
	}

	for {
		fmt.Printf("猫的信息[no: %d, name: %s]\n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

//删除一只猫
func DelCatNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head

	//将helper定位到链表的最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	//空链表
	if temp.next == nil {
		fmt.Println("链表为空")
		return head
	}

	//只有一个节点
	if temp.next == head {
		if temp.no == id {
			temp.next = nil
		}

		return head
	}

	//如果有两个及以上的节点
	flag := true
	for {
		if temp.next == head { //说明比较到最后一个了，最后一个没有比较
			break
		}

		if temp.no == id { //找到了
			if temp == head {//要删除的是头节点
				head = head.next
			}

			//删除temp节点
			helper.next = temp.next
			fmt.Printf("删除的猫猫：%d", id)
			flag = false
			break
		}

		temp = temp.next     //移动，用于id的比较
		helper = helper.next //移动，辅助删除temp节点
	}

	if flag {//flag为true，说明前面没有删除，在这里判断最后一个节点
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("删除猫猫：%d", id)
		}else {
			fmt.Printf("没有找到id为%d的猫猫\n", id)
		}
	}

	return head
}

func main() {
	//初始化一个环形节点
	head := &CatNode{}

	//创建一只猫
	cat1 := &CatNode{
		no : 1,
		name: "tom",
	}
	cat2 := &CatNode{
		no : 2,
		name: "jom",
	}
	cat3 := &CatNode{
		no : 3,
		name: "john",
	}

	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	ListCatNode(head)

	head = DelCatNode(head, 11)

	ListCatNode(head)
}
