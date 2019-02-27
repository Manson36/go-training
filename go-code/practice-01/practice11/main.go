package main

import "fmt"

func main(){

	//遍历字符串
	var str string ="Hello world!"

	for i:=0; i<len(str); i++{
		fmt.Printf("%c \n", str[i])
	}
	//字符串 遍历方式2
	str="abcdefg"

	for index,val:=range str {
		fmt.Printf("index=%v  val=c% \n",index,val)
	}
}