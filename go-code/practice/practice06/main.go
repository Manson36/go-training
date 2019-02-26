package main

import (
	"fmt"
	"strconv"
)

func main(){

	//string 类型转基本数据类型
	var str string ="true"
	var b bool
	b,_=strconv.ParseBool(str)
	fmt.Printf("b type %T b =%v\n",b,b)

	var str2 string ="12345566"
	var n1 int64
	var n2 int
	n1,_=strconv.ParseInt(str2,10,64)
	n2 =int(n1)
	fmt.Printf("n1 type %T n1 =%v\n",n1,n1)
	fmt.Printf("n2 type %T n2 =%v\n",n2,n2)

	var str3 string ="123.456"
	var f1 float64
	f1,_=strconv.ParseFloat(str3,64)
	fmt.Printf("f1 type %T f1 =%v\n",f1,f1)

	//将Hello 转成整数失败，n3恢复默认值
	var str4 string ="Hello"
	var n3 int64 =11
	n3,_=strconv.ParseInt(str4,10,64)
	fmt.Printf("n3 type %T n3 =%v\n",n3,n3)
}