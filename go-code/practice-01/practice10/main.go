package main

import (
	"fmt"
)

func main(){

	//for 循环基础学习
	for i:= 1; i<5; i++{
		fmt.Printf("Hwllo %v \n",i)
	}

	//for循环的第二种写法
	j:=1
	for j<5 {
		fmt.Printf("Hello~ %v \n",j)
		j++
	}

	//for循环的第三种写法
	k:=1
	for{
		if k<5{
			fmt.Printf("Hello World %v \n",k)
		}else{
			break
		}
		k++
	}
}