package main

import (
	"fmt"
	"time"
)

//使用defer和recover来捕获和处理异常

func test()  {

	defer func (){
		err:= recover()
		if err!= nil{
			fmt.Println("err=",err)
			//发送邮件给、、、
			fmt.Println("发送邮件给。。。")
		}
	}()
	num1 := 10
	num2 :=0
	res := num1/num2
	fmt.Println("res",res)
}

func main(){

	test()
	for {
		fmt.Println("main()下面的代码")
		time.Sleep(time.Second)

	}
}