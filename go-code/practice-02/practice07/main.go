package main

import "fmt"

func main(){

	//演示goto的使用

	var n1 =30

	if n1 >20{
		goto lable1
	}
	fmt.Println("ok1")
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
	lable1:
	fmt.Println("ok5")
	fmt.Println("ok6")
	fmt.Println("ok7")
	fmt.Println("ok8")

	//return 的使用

	for i:=1; i<10;i++{

		if i==3{
			return
		}
		fmt.Println("wawa",i)
	}
	fmt.Println("Hello World")
}