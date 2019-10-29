package main

import "fmt"

type Hero struct {
	No int
	Name string
	Left *Hero
	Right *Hero
}

//前序遍历：先输出根节点，左节点，右节点
func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("no=%d, name=%s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

//中序遍历：左节点，根节点，右节点
func InfixOrder(node *Hero) {
	if node != nil {
		PreOrder(node.Left)
		fmt.Printf("no=%d, name=%s\n", node.No, node.Name)
		PreOrder(node.Right)
	}
}

//后序遍历
func PostOrder(node *Hero) {
	if node != nil {
		PreOrder(node.Left)
		PreOrder(node.Right)
		fmt.Printf("no=%d, name=%s\n", node.No, node.Name)
	}
}

func main() {
	//创建一个二叉树
	root := &Hero{
		No: 1,
		Name: "宋江",
	}
	Left1 := &Hero{
		No: 2,
		Name: "吴用",
	}
	right1 := &Hero{
		No: 3,
		Name: "卢俊义",
	}
	right2 := &Hero{
		No: 4,
		Name: "林冲",
	}

	root.Left = Left1
	root.Right = right1
	right1.Right = right2

	PreOrder(root)
	InfixOrder(root)
	PostOrder(root)
}
