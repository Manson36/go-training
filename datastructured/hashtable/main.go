package main

import (
	"fmt"
	"os"
)

type Emp struct {
	Id int
	Name string
	Next *Emp
}

//我们这里的EmpLink不带表头，第一个节点就放雇员
type EmpLink struct {
	Head *Emp
}

//显示emp内容
func (this *Emp) ShowMe() {
	fmt.Printf("链表%d 找到该雇员 %d\n", this.Id % 7, this.Id)
}

//1.添加员工的方法，保证添加时是从小到大
func (this *EmpLink) Insert(emp *Emp) {
	cur := this.Head //这是一个辅助指针
	var pre *Emp = nil //这是一个辅助指针在cur之前

	//如果当前链表为空
	if cur == nil {
		this.Head = emp
		return
	}

	//如果不是空，找到位置插入
	for {
		if cur != nil {
			//比较
			if cur.Id >= emp.Id {
				break //找到位置
			}

			pre = cur
			cur = cur.Next
		}else {
			break
		}
	}

	//插入，分析位置3个：头，中间，尾部。
	pre.Next = emp
	emp.Next = cur
}

//显示链表信息
func (this *EmpLink) ShowLink(no int) {
	if this.Head == nil {
		fmt.Printf("链表%d为空\n", no)
		return
	}

	cur := this.Head
	for {
		if cur != nil {
			fmt.Printf("链表：%d，雇员id=%d，雇员姓名=%s", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}

//根据id查找雇员
func (this *EmpLink) FindById(id int) *Emp {
	cur := this.Head

	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}

		cur = cur.Next
	}

	return nil
}

//定义hashtable，含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

//给hashTable增加一个添加的方法：
func (this *HashTable) Insert(emp *Emp) {
	//使用散列函数，确定将雇员添加到哪个链表
	linkNo := this.HashFun(emp.Id)
	//使用对应的链表添加
	this.LinkArr[linkNo].Insert(emp)
}

//编写一个散列方法
func (this *HashTable) HashFun(id int) int {
	return id % 7//得到的值就是链表的下标
}

//编写方法，显示hashtable所有雇员
func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}

//编写方法，完成查找
func (this *HashTable) FindById(id int) *Emp {
	linkNo := this.HashFun(id)
	return this.LinkArr[linkNo].FindById(id)
}

func main() {
	key := ""
	id := 0
	name := ""
	var hashTable HashTable

	for {
		fmt.Println("=====雇员系统菜单======")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show 表示显示雇员")
		fmt.Println("find 表示查找雇员")
		fmt.Println("exit 表示退出")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("输入雇员name")
			fmt.Scanln(&name)
			emp := Emp{Id: id, Name: name}

			hashTable.Insert(&emp)
		case "show":
			hashTable.ShowAll()
		case "find":
			fmt.Println("请输入id号")
			fmt.Scanln(&id)
			emp := hashTable.FindById(id)

			if emp == nil {
				fmt.Println("id = %d 的雇员不存在\n", id)
			}else {
				emp.ShowMe()
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入错误")
		}

	}
}
