package main

import "fmt"

//支持对函数的返回值命名

func getSumAndSub(n1 int, n2 int )(sum int ,sub int){

	sum = n1+n2
	sub = n1-n2
	return
}
func main(){

	a1,b1:=getSumAndSub(1,3)
	fmt.Println("a1 =",a1,"b1 =",b1)

	//使用_标识符忽略返回值
	res,_:=getSumAndSub(3,2)
	fmt.Println("res=",res )
}