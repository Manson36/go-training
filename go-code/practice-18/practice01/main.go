package main

import "fmt"

//单元测试 传统方法
func addUpper(n int)int {

	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func main() {

	res := addUpper(10)
	if res != 55 {
		fmt.Printf("addUpper 错误 返回值%v 期望值%v",res,55)
	}else {
		fmt.Printf("addUpper 正确 返回值%v 期望值%v",res,55)
	}

}