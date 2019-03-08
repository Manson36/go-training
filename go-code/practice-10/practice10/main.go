package main

import "fmt"

type Visitor struct {
	Name string
	Age int
}

func (Visitor Visitor)showPriece(){
	if Visitor.Age >= 90||Visitor.Age < 8{
		fmt.Println("考虑安全就不要玩了")
		return
	}
	if Visitor.Age > 18{
		fmt.Println("游客年龄为%v 姓名为%v 门票20元",Visitor.Age,Visitor.Name)
	}else {
		fmt.Println("游客年龄为%v 姓名为%v 免费",Visitor.Age,Visitor.Name)
	}
}

func main() {

	var v Visitor
	for {
		fmt.Println("请输入你的名字")
		fmt.Scanln(&v.Name)
		if v.Name =="n"{
			fmt.Println("退出程序")
			break
		}
		fmt.Println("请输入你的年龄")
		fmt.Scanln(&v.Age)
		v.showPriece()
	}
}