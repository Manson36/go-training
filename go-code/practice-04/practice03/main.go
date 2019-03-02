package main

import "fmt"

//函数修改函数外的变量，使用指针
func test(n1 *int ){

	*n1 = *n1 + 10
	fmt.Println("test n1 =",*n1)

}

func main(){

	 num:= 20
	 test(&num)
	 fmt.Println("main() num=",num)
}