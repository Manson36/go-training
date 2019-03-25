package main

import "fmt"

func main() {

	//查找
	names := [4]string{"白眉", "青翼", "金毛", "紫衫"}
	var heroName= ""

	fmt.Println("请输入英雄的名字")
	fmt.Scanln(&heroName)

	//for i := 0; i < len(names); i++ {
	//	if heroName == names[i] {
	//		fmt.Printf("找到%v，下标为%v\n", heroName, i)
	//		break
	//	} else if i == (len(names) - 1) {
	//		fmt.Println("没有找到", heroName)
	//	}
	//}

	//方法二：
	var index= -1

	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			index = i
			break
		}
	}
		if index != -1 {
			fmt.Printf("找到%v，下标为%v\n", heroName, index)
		} else  {
			fmt.Println("没有找到", heroName)
		}
}