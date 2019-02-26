package main

import "fmt"

func main(){

	//从控制台接受用户信息 ：姓名，年龄，薪水，是否通过考试。Scanln。
	var name string
	var age int
	var sal float32
	var ispass bool
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过考试")
	fmt.Scanln(&ispass)

	fmt.Printf("姓名是 %v 年龄是 %v 薪水是 %v 是否通过考试 %v \n",name,age,sal,ispass)

	fmt.Println("请输入你的姓名，年龄，薪水，是否通过考试")
	fmt.Scanf("%d %d %f %t",&name,&age,&sal,&ispass)
	fmt.Printf("姓名是 %v 年龄是 %v 薪水是 %v 是否通过考试 %v",name,age,sal,ispass)
}