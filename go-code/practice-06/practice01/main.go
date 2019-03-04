package main

import "fmt"

//为什么需要数组
	/*养鸡场有6只鸡，他们的体重分别是3,5，1,3.4,2,50（kg），
	请问六只鸡的总体中是多少，平均体重是多少
	 */
func main(){

	//传统方法：定义六个变量
	hen1 := 3.0
	hen2 := 5.0
	hen3 := 1.0
	hen4 := 3.4
	hen5 := 2.0
	hen6 := 50.0

	weight := hen1+hen2+hen3+hen4+hen5+hen6
	aveweight :=fmt.Sprintf("%.2f",weight/6)

	fmt.Println("总重量为",weight,"平均重量是",aveweight)
}