package main

import "fmt"

func main(){

	//统计三个班的成绩，每个班有五名学生
	//求每个班的成绩和所有班级的总成绩

	//分析实现思路1
	//1、统计一个班5名同学的平均分
	//2.学生数就是5个
	//3.声明一个sum 计算班级总分

	//分析实现思路2
	//1.统计三个班级的情况
	//2.j表示第几个班级
	//3.定义一个变量定义总成绩

	//分析实现思路3
	//1.把代码激活
	//2.定义两个变量，表示班级数和每个班的人数

	var classNum int = 3
	var stuNum int =5
	var totalSum float64 =0.0

	for j:=1;j<= classNum;j++ {
		var sum float64 = 0.0
		for i := 1; i <= stuNum; i++ {
			var score float64
			fmt.Printf("请输入第%d 班的第 %d 学生成绩\n", j, i)
			fmt.Scanln(&score)
			sum += score
		}
		fmt.Printf("第 %d 班的平均成绩为%v \n", j, sum/float64(stuNum))
		totalSum += sum
	}
	fmt.Println("总的平均成绩为", totalSum/float64(stuNum))
}