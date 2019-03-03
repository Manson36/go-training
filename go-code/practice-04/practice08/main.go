package main

import (
	"fmt"
	"strconv"
)

func main(){

	//统计字符串的字节长度
	str:="Hello bejing"
	fmt.Println("str len =",len(str))

	str2 := "Hello 北京"
	r:=[]rune(str2)
	for i:=0; i< len(r); i++{
		fmt.Printf("字符为%c \n",r[i])
	}

	//字符串转整数
	n, err := strconv.Atoi("hello")
	if err != nil{
		fmt.Println("转换错误",err)
	}else {
		fmt.Println("转成的结果是",n)
	}

	//整数转字符串
	str = strconv.Itoa(123455)
	fmt.Printf("str = %v, str= %T",str,str)

	var bytes =[]byte("hello go")
	fmt.Printf("bytes=%v \n",bytes)

	str =string([]byte{87,55,44})
	fmt.Printf("str=%v",str )
}