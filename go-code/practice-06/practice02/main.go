package main

import "fmt"

func main(){

	//使用数组
	var hens [7]float64

	hens[0] = 150.0
	hens[1] = 3.0
	hens[2] = 5.0
	hens[3] = 1.0
	hens[4] = 3.4
	hens[5] = 2.0
	hens[6] = 50.0

	var weight = 0.0

	for i:=0; i<7; i++{
		weight += hens[i]
	}

	aveweight :=fmt.Sprintf("%.2f",weight/6)

	fmt.Println("总重量为",weight,"平均重量是",aveweight)
}