package main

import (
	"fmt"
)

func main(){

	//求一个数组的最大值，并得到相应的下标
	var intArr [6]int = [...]int {1,33,55,1,-3,99}
	maxVal := intArr[0]
	maxIndex := 0

	for i:=1; i<len(intArr); i++{

		if maxVal< intArr[i]{
			maxVal=intArr[i]
			maxIndex=i
		}
	}
	fmt.Printf("maxVal=%v maxIndex=%v\n",maxVal,maxIndex)

	//求出数组的总和和平均值
	sum:=0
	//for i:=0; i<len(intArr); i++{
		//累计求和
		//sum+= intArr[i]
	//}
	for _, value := range intArr{
		//累计求和
		sum +=value
	}
	fmt.Printf("sum=%v, 平均值为%v\n",sum,float64(sum)/float64(len(intArr)))
}