package main

import "fmt"

func main(){
	//golang中基本数据类型的转换
	var i int32 =100
	var n1 float32 =float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)
	fmt.Printf("i=%v, n1=%v, n2=%v, n3=%v \n",i,n1,n2,n3)
	fmt.Printf("i的数据类型是 %T",i)

}