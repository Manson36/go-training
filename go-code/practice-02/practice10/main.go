package main

import (
	"fmt"
	"github.com/go-training/go-code/practice-02/utils"
)

func main(){

	var n1 float64=1.2
	var n2 float64=2.3
	var operator byte ='/'
	result := utils.Cal(n1,n2,operator)
	fmt.Println("result=",result)
}