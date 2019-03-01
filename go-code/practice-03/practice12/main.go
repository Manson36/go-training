package main

import "fmt"

var age = test()

type af func (string, int) bool

func test() int {
	fmt.Println("test()...")
	return 90
}
func init(){
	fmt.Println("init()...")
}
func main(){
	fmt.Println("main()...age =",age)
	c := test1("bbb", print2)
	c()
}

// 参数的类型，顺序，个数，返回值的个数，类型，顺序
func test1(result string, print af) func () {
	fmt.Println(print(result, 1))

	return func () {
		fmt.Println(result)
	}
}

func print1(ssss string, aa int) bool {
	fmt.Println(ssss)
	fmt.Println(aa)
	return true
}

func print2(sssssssss string, bb int) bool {
	fmt.Println(sssssssss)
	fmt.Println(bb)
	return true
}