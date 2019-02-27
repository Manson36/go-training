package main

import "fmt"

func main(){

	//三个条件：身高180cm，财富1000w，帅
	var height float32
	var money float64
	var handsome bool

	fmt.Println("请输入身高（厘米）")
	fmt.Scanln(&height)
	fmt.Println("请输入财富（千万）")
	fmt.Scanln(&money)
	fmt.Println("请输入是否帅")
	fmt.Scanln(&handsome)

	if height>180&& money> 1.0&& handsome{
		fmt.Println("我一定嫁给他！")
	}else if height>180||money>1.0|| handsome{
		fmt.Println("嫁吧，比上不足，比下有余。")
	}else{
		fmt.Println("不嫁、、、")
	}

	//百米比赛，成绩13秒之内进入决赛，根据性别分配男子组和女子组
	var second float64

	fmt.Println("请输入百米用时")
	fmt.Scanln(&second)

	if second <= 13 {
		var genter string

		fmt.Println("请输入性别")
		fmt.Scanln(&genter)
		if genter =="男"{
			fmt.Println("进入决赛男子组")
		}else if genter == "女"{
			fmt.Println("请进入女子组")
		}
	}else {
		fmt.Println("你被淘汰了")
	}
}