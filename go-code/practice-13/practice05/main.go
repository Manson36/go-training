package main

import "fmt"

func TypeJudge (items... interface{}){
	for index, x := range items{
		switch x.(type) {
		case bool:
			fmt.Printf ("第%v个 参数是bool类型，值是%v",index,x)
		case float32:
			fmt.Printf ("第%v个 参数是float32类型，值是%v",index,x)
		case float64:
			fmt.Printf ("第%v个 参数是float64类型，值是%v",index,x)
		case int,int32,int64:
			fmt.Printf ("第%v个 参数是整数 类型，值是%v",index,x)
		case string:
			fmt.Printf ("第%v个 参数是string类型，值是%v",index,x)
		default:
			fmt.Printf ("第%v个 参数是 不确定 类型，值是%v",index,x)
		}
	}
}

func main() {

	var n1 float32 =1.2
	var n2 float64 = 3.4
	var n3 int32 = 4
	var n4 string = "Tom"
	address := "北京"
	n5 := 300

	TypeJudge(n1,n2,n3,n4,n5,address)
}