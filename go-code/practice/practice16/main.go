package main

import "fmt"

func main() {

	var age int
	fmt.Println("请输入你的年龄")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("你的年龄大于十八岁，你要对自己的行为负责！")
	} else {
		fmt.Println("你的年龄太小，这次放过你了。")
	}
	var n1 int32 = 10
	var n2 int32 = 50
	if n1+n2 >= 50 {
		fmt.Println("Hwllo World")
	}
	var n3 float64 =11.1
	var n4 float64 =17.2
	if n3>10&& n4<20{
		fmt.Println("和=",(n3+n4))
	}

}