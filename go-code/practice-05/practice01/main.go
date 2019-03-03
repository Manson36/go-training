package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){

	//十进制转二进制
	str := strconv.FormatInt(123,2)
	fmt.Printf("123对应的二进制是%v\n",str)

	//十进制转十六进制
	str = strconv.FormatInt(123,16)
	fmt.Printf("123对应的十六进制是%v\n",str)

	//查找子串是否在指定的字符串中
	b:= strings.Contains("seefood","food")
	fmt.Printf("b =%v\n",b)
	b= strings.Contains("seefood","fd")
	fmt.Printf("b =%v\n",b)

	//统计一个字符串有几个指定的子串
	num:= strings.Count("cheese","e")
	fmt.Printf("num=%v\n",num)

	//不区分大小写的比较
	b = strings.EqualFold("abc","ABc")
	fmt.Printf("b=%v\n",b)

	fmt.Print("结果","abc"=="ABc")
}
