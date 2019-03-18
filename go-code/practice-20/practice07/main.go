package maint

import (
	"fmt"
	"reflect"
)

//反射修改变量
func testInt(b interface{}) {

	val := reflect.ValueOf(b)
	fmt.Printf("val 的类型%T",val)
	val.Elem().SetInt(111)
	fmt.Println("val=",val)
}

func main() {

	var num int = 11
	testInt(&num)

	fmt.Println("num=",num)
}