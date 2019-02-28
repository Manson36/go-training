package main

import "fmt"

func main(){

	//指定标签的使用
	lable:
		for i:=0;i<4; i++{
			for j:=0; j<10; j++{
				if j== 2{
					break lable
				}
				fmt.Println("j=",j)
			}
		}
}