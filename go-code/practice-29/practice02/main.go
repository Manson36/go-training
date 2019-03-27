package main

import "fmt"

//结构体 景区门票案例

type Visitor struct {
	name string
	age int
}
func (visitor *Visitor)showPrice() {
	if visitor.age < 8 || visitor.age > 90{
		fmt.Println("考虑到年龄，就不要玩了")
		return
	}else if visitor.age > 18 {
		fmt.Printf("游客的姓名%v，年龄%v，门票价格为20\n",visitor.name,visitor.age)
	}else {
		fmt.Printf("游客的姓名%v，年龄%v，门票免费\n",visitor.name,visitor.age)
	}
}

func main() {

	var v Visitor
	for {

		fmt.Println("请输入你的姓名")
		fmt.Scanln(&v.name)
		if v.name == "n" {
			fmt.Println("退出程序")
			break
		}

		fmt.Println("请输入你的年龄")
		fmt.Scanln(&v.age)

		v.showPrice()
}
}