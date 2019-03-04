package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	//随机生成5个数，并将其反转打印
	var intArr [5]int
	len:=len(intArr)

	rand.Seed(time.Now().UnixNano())
	for i:=0; i< len; i++{
		intArr[i]= rand.Intn(100)
	}
	fmt.Println("交换前intArr=",intArr)

	//临时变量
	temp:=0
	for i:=0; i<len/2; i++{
		temp= intArr[i]
		intArr[i] = intArr[len-1-i]
		intArr[len-1-i] = temp
	}
	fmt.Println("交换后intArr=",intArr)
}