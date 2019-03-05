package main

import "fmt"

func main(){

	//顺序查找
	names := [4]string{"白眉","金毛","紫衫","青衣"}
	var heroName = ""
	fmt.Println("请输入名字")
	fmt.Scanln(&heroName)

	for i:=0; i< len(names); i++{
		if names[i]==heroName{
			fmt.Printf("找到%v 下标%v\n",heroName,i)
			break
		}else if i == len(names)-1{
			fmt.Printf("没有找到%v",heroName)
		}
	}

	//方式2 推荐
	index:= -1

	for i:=0; i< len(names); i++{
		if names[i]==heroName{
			index = i
			break
		}
		if index!= -1{
			fmt.Printf("找到%v 下标%v\n",heroName,i)
		}else {
			fmt.Printf("没有找到%v",heroName)
		}
	}
}