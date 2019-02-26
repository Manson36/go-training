package main

import "fmt"

func main(){

	var days int =97
	var week int =days/7
	var day int =days%7
	fmt.Printf("%d 个星期零%d 天\n",week,day)

	var huashi float32 =134.2
	var sheshi float32 =5/9*(huashi-100)
	fmt.Printf("%v 对应的摄氏温度是 %v \n",huashi,sheshi)

	//关系运算符的使用
	var n1 int =8
	var n2 int =9
	fmt.Println(n1==n2)
	fmt.Println(n1!=n2)
	fmt.Println(n1>n2)
	fmt.Println(n1>=n2)
	fmt.Println(n1<n2)
	fmt.Println(n1<=n2)
	flag:=n1>n2
	fmt.Println("flag=",flag)

}