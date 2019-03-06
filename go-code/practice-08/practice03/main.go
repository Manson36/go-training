package main

import "fmt"

func main(){

	var scores [3][5]float64

	for i := 0; i < len(scores); i++{
		for j := 0; j < len(scores[i]); j++{
			fmt.Printf("请输入第%d个班的第%d个学生的成绩\n")
			fmt.Scanln(&scores[i][j])
		}
	}
	totalSum := 0.0
	for i := 0; i < len(scores); i++{
		sum := 0.0
		for j := 0; j < len(scores[i]); j++{
			sum += scores[i][j]
		}
		totalSum += sum
		fmt.Printf("第%d个班的总成绩为%v 平均成绩%v",
			i+1,sum,sum/float64(len(scores[i])))
		fmt.Printf("所有班级的总成绩%d，平均成绩%d",
			totalSum,totalSum/15)
	}

}